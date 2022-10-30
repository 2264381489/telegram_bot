package main

import (
	"encoding/json"
	"flag"
	"log"
	"telegram-bot/bot/telegramBot/internal/config"
	"telegram-bot/bot/telegramBot/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "/Users/yanzezhong/Desktop/code/telegram-bot/telegramBot/etc/telegrambot-api.yaml", "the config file")

var ForwardMap map[int64]bool

func main() {
	ForwardMap = map[int64]bool{}
	bot := telegram.InitBot()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	c := config.Config{}
	conf.MustLoad(*configFile, &c)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}
		//if !update.Message.IsCommand() { // ignore any non-command Messages
		//	continue
		//}

		if update.FromChat() != nil {
			//todo 如果是转发来的,直接转发给对应channel
			if update.Message.ForwardFromMessageID != 0 {
				forChatId, err := telegram.ForwardMsg(update, c.TargetChannel, bot)
				if err != nil {
					panic(err)
				}
				// 转发信息触发翻转标记
				ForwardMap[forChatId] = true
				continue
			}

			//todo 想办法增加 forward 加评论的功能
			forChatId := update.FromChat().ID
			f := ForwardMap[forChatId]
			if f {
				message := tgbotapi.NewCopyMessage(c.TargetChannel, forChatId, update.Message.MessageID) // 成功
				_, err := bot.CopyMessage(message)
				if err != nil {
					panic(err)
				}
				delete(ForwardMap, forChatId)
			}

		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		// 让技术
		commands := telegram.CommandList()
		marshal, err := json.Marshal(update)
		if err != nil {
			panic(err)
		}
		logx.Info("测试数据 : %s", string(marshal))

		// Extract the command from the Message.
		switch update.Message.Command() {

		case "start":
			request, err := bot.Request(commands)
			if err != nil {
				panic(err)
			}
			logx.Info(request.Result)
			msg.Text = "已初始化命令面板"
		case "post":

		case "help":
			msg.Text = "I understand /sayhi and /status."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."

		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}

	}
}
