package telegram

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"tg-bot-template/pkg/fsm"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	log         *zap.Logger
	api         *tgbotapi.BotAPI
	fsm         fsm.StateManager
	routes      map[string]HandlerFunc
	stateRoutes map[string]HandlerFunc
}

type HandlerFunc func(api *tgbotapi.BotAPI, update tgbotapi.Update)

func New(logger *zap.Logger, stateManager fsm.StateManager, token string) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("error with api token")
	}
	return &Bot{
		log:         logger,
		api:         api,
		fsm:         stateManager,
		routes:      make(map[string]HandlerFunc),
		stateRoutes: make(map[string]HandlerFunc),
	}

}

func (b *Bot) Handle(command string, handler HandlerFunc) {
	b.routes[command] = handler
}

func (b *Bot) HandleState(state string, handler HandlerFunc) {
	b.stateRoutes[state] = handler
}

func (b *Bot) Run() {
	go b.Stop()
	u := tgbotapi.NewUpdate(0)
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		userID := update.Message.From.ID
		text := update.Message.Text

		currentState := b.fsm.GetState(userID)

		if currentState != "" {
			if handler, exists := b.stateRoutes[currentState]; exists {
				handler(b.api, update)
				continue
			}
		}

		if handler, exists := b.routes[text]; exists {
			handler(b.api, update)
		} else {
			continue
		}
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
