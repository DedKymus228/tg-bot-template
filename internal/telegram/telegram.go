package telegram

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	log *zap.Logger
	api *tgbotapi.BotAPI
}

func New(logger *zap.Logger, token string) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("error with api token")
	}
	return &Bot{
		log: logger,
		api: api,
	}

}

func (b *Bot) Run() {
	go b.Stop()
	u := tgbotapi.NewUpdate(0)
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		b.log.Info("Get msg: ", zap.String("text", update.Message.Text))
	}
	return
}

func (b *Bot) Stop() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	b.log.Info("TG-Bot stopping...")
	b.api.StopReceivingUpdates()

}
