package handlers

import (
	"tg-bot-template/internal/storage"
	"tg-bot-template/internal/telegram"
)

func RegisterRoute(bot *telegram.Bot, store storage.Storage) {
	bot.Handle("/start", StartCommand)
}
