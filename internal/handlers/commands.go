package handlers

import (
	"tg-bot-template/pkg/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	startCmd = "/start"
)

const (
	testBtn = "test_btn"
)

//here you will find commands that you can write specifically for your own bot

func StartCommand(api *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	msg := tgbotapi.NewMessage(chatID, "You use 'start' command")

	keyboard := telegram.GetKeyboardButtonData("BTN", testBtn)
	msg.ReplyMarkup = keyboard
	api.Send(msg)
}

func TestButtonCallback(api *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	msg := tgbotapi.NewMessage(chatID, "Status OK 200")
	api.Send(msg)
}
