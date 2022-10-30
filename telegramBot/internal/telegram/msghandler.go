package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// ForwardMsg 处理转发消息
func ForwardMsg(update tgbotapi.Update, targetChannel int64, bot *tgbotapi.BotAPI) (int64, error) {
	forChatId := update.FromChat().ID
	message := tgbotapi.NewCopyMessage(targetChannel, forChatId, update.Message.MessageID) // 成功
	_, err := bot.CopyMessage(message)
	if err != nil {
		return 0, err
	}

	newMessage := tgbotapi.NewMessage(update.FromChat().ID, "请发送备注信息") // 失败
	_, err = bot.Send(newMessage)
	if err != nil {
		return 0, err
	}
	return forChatId, nil
}
