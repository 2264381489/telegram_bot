package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type COMMAND tgbotapi.BotCommand

func CommandList() tgbotapi.SetMyCommandsConfig {

	test := tgbotapi.BotCommand{
		Command:     "test",
		Description: "a test command",
	}
	twitter := tgbotapi.BotCommand{
		Command:     "twitter",
		Description: "splite as twitter",
	}
	rss := tgbotapi.BotCommand{
		Command:     "rss",
		Description: "descripet Rss",
	}
	restart := tgbotapi.BotCommand{
		Command:     "restart",
		Description: "restart server",
	}
	international := tgbotapi.BotCommand{
		Command:     "international",
		Description: "no use",
	}
	return tgbotapi.NewSetMyCommands(test, twitter, rss, restart, international)
}
