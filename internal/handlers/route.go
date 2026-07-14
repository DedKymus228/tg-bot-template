package handlers

import (
	"tg-bot-template/pkg/storage"
	"tg-bot-template/pkg/telegram"
)

func RegisterRoute(bot *telegram.Bot, store storage.Storage) {
	//commands
	bot.Handle(startCmd, StartCommand)

	// buttons
	bot.HandleCallback(testBtn, TestButtonCallback)
}
