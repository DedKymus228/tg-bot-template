package main

import (
	"context"
	"log"
	"os"
	"tg-bot-template/internal/config"
	"tg-bot-template/internal/handlers"
	"tg-bot-template/internal/jobs"
	"tg-bot-template/pkg/fsm"
	"tg-bot-template/pkg/middlewares"
	"tg-bot-template/pkg/mylogger"
	"tg-bot-template/pkg/scheduler"
	storage2 "tg-bot-template/pkg/storage"
	"tg-bot-template/pkg/telegram"

	"github.com/ilyakaznacheev/cleanenv"
	"go.uber.org/zap"
)

func main() {
	var cfg config.Config
	if _, err := os.Stat(".env"); err == nil {
		err = cleanenv.ReadConfig(".env", &cfg)
		if err != nil {
			log.Fatal("config not found:", zap.Error(err))
		}
	} else {
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Fatal("error read config:", zap.Error(err))
		}
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
	bot.UseMW(middlewares.Recovery(logger))

	cron := scheduler.New(logger)
	cron.Run()
	defer cron.Stop()

	jobs.RegisterJobs(cron, bot, store, logger)
	bot.Run()
}
