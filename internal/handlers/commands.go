package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

//here you will find commands that you can write specifically for your own bot

func StartCommand(api *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	msg := tgbotapi.NewMessage(chatID, "You use 'start' command")
	api.Send(msg)
}
