package handlers

import "tg-bot-template/internal/telegram"

func RegisterRoute(bot *telegram.Bot) {
	bot.Handle("/start", StartCommand)
}
