package logic

import (
	"context"

	"telegram-bot/bot/telegramBot/internal/svc"
	"telegram-bot/bot/telegramBot/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TelegramBotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTelegramBotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TelegramBotLogic {
	return &TelegramBotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TelegramBotLogic) TelegramBot(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
