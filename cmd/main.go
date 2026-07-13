package main

import (
	"context"
	"log"
	"tg-bot-template/internal/config"
	"tg-bot-template/internal/fsm"
	"tg-bot-template/internal/handlers"
	"tg-bot-template/internal/storage"
	"tg-bot-template/internal/telegram"
	"tg-bot-template/pkg/mylogger"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config:", zap.Error(err))
	}

	logger := mylogger.New(cfg.Env)

	db, err := storage.New(context.Background(), cfg.DB)
	if err != nil {
		logger.Fatal("error connecting to db:", zap.Error(err))
	}
	defer db.ClosePool()

	fsmManager := fsm.NewMemoryFSM()

	bot := telegram.New(logger, fsmManager, cfg.BotToken)
	handlers.RegisterRoute(bot)
	bot.Run()
}
