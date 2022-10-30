package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	TargetChannel    int64  `json:"targetChannel"`
	TelegramBotToken string `json:"telegramBotToken"`
}
