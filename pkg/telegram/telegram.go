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
	log            *zap.Logger
	api            *tgbotapi.BotAPI
	fsm            fsm.StateManager
	routes         map[string]HandlerFunc
	stateRoutes    map[string]HandlerFunc
	callbackRoutes map[string]HandlerFunc
	middlewares    []MiddlewareFunc
}

type HandlerFunc func(api *tgbotapi.BotAPI, update tgbotapi.Update)
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

func GetKeyboardButtonData(text, data string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(text, data),
		),
	)
}

func New(logger *zap.Logger, stateManager fsm.StateManager, token string) *Bot {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("error with api token")
	}
	return &Bot{
		log:            logger,
		api:            api,
		fsm:            stateManager,
		routes:         make(map[string]HandlerFunc),
		stateRoutes:    make(map[string]HandlerFunc),
		callbackRoutes: make(map[string]HandlerFunc),
		middlewares:    make([]MiddlewareFunc, 0),
	}

}

func (b *Bot) Handle(command string, handler HandlerFunc) {
	b.routes[command] = handler
}

func (b *Bot) HandleState(state string, handler HandlerFunc) {
	b.stateRoutes[state] = handler
}

func (b *Bot) HandleCallback(callbackData string, handler HandlerFunc) {
	b.callbackRoutes[callbackData] = handler
}

func (b *Bot) UseMW(mw MiddlewareFunc) {
	b.middlewares = append(b.middlewares, mw)
}

func (b *Bot) applyMiddlewares(handler HandlerFunc) HandlerFunc {
	for i := len(b.middlewares) - 1; i >= 0; i-- {
		handler = b.middlewares[i](handler)
	}
	return handler
}

func (b *Bot) Run() {
	go b.Stop()
	u := tgbotapi.NewUpdate(0)
	updates := b.api.GetUpdatesChan(u)
	for update := range updates {

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			b.api.Request(callback)

			data := update.CallbackQuery.Data
			if handler, exists := b.callbackRoutes[data]; exists {
				finalHandler := b.applyMiddlewares(handler)
				go finalHandler(b.api, update)
			} else {
				b.log.Info("Unknown callback:", zap.String("data:", data))
			}
			continue
		}
		if update.Message == nil {
			continue
		}
		userID := update.Message.From.ID
		text := update.Message.Text

		currentState := b.fsm.GetState(userID)

		if currentState != "" {
			if handler, exists := b.stateRoutes[currentState]; exists {
				finalHandler := b.applyMiddlewares(handler)
				go finalHandler(b.api, update)
				continue
			}
		}

		if handler, exists := b.routes[text]; exists {
			finalHandler := b.applyMiddlewares(handler)
			go finalHandler(b.api, update)
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

func (b *Bot) API() *tgbotapi.BotAPI {
	return b.api
}
