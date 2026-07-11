package main

import (
	"tg-bot-template/internal/config"
	"tg-bot-template/internal/telegram"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}

	bot := telegram.New(cfg.BotToken)

	bot.Run()
}
