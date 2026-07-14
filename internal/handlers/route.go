package handlers

import (
	"tg-bot-template/pkg/storage"
	"tg-bot-template/pkg/telegram"
)

func RegisterRoute(bot *telegram.Bot, store storage.Storage) {
	bot.Handle("/start", StartCommand)
}
