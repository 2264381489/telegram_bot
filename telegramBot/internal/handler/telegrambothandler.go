package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"telegram-bot/bot/telegramBot/internal/logic"
	"telegram-bot/bot/telegramBot/internal/svc"
	"telegram-bot/bot/telegramBot/internal/types"
)

func TelegramBotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewTelegramBotLogic(r.Context(), svcCtx)
		resp, err := l.TelegramBot(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
