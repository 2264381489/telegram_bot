package telegram

import (
	"log"
	"os"
	"telegram-bot/bot/telegramBot/internal/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func InitBot(c config.Config) *tgbotapi.BotAPI {
	err := os.Setenv("TELEGRAM_APITOKEN", c.TelegramBotToken)
	if err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if err != nil {
		panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}
