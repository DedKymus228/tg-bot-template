package main

import (
	"context"
	"log"
	"tg-bot-template/internal/config"
	"tg-bot-template/internal/handlers"
	"tg-bot-template/pkg/fsm"
	"tg-bot-template/pkg/mylogger"
	storage2 "tg-bot-template/pkg/storage"
	"tg-bot-template/pkg/telegram"

	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config:", zap.Error(err))
	}

	logger := mylogger.New(cfg.Env)

	var store storage2.Storage
	if cfg.DBEnabled {
		db, err := storage2.New(context.Background(), logger, cfg.DB)
		if err != nil {
			logger.Fatal("error connecting to db:", zap.Error(err))
		}
		store = db
		defer db.ClosePool()
	} else {
		store = &storage2.NoopStorage{}
	}

	fsmManager := fsm.NewMemoryFSM()

	bot := telegram.New(logger, fsmManager, cfg.BotToken)
	handlers.RegisterRoute(bot, store)
	bot.Run()
}
