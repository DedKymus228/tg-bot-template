package middlewares

import (
	"sync"
	"tg-bot-template/pkg/telegram"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

func Recovery(log *zap.Logger) telegram.MiddlewareFunc {

	return func(next telegram.HandlerFunc) telegram.HandlerFunc {
		return func(api *tgbotapi.BotAPI, update tgbotapi.Update) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("Panic ", zap.Any("error", err))

					var chatID int64
					if update.Message != nil {
						chatID = update.Message.Chat.ID
					} else if update.CallbackQuery != nil {
						chatID = update.CallbackQuery.Message.Chat.ID
					}

					if chatID != 0 {
						msg := tgbotapi.NewMessage(chatID, "Sorry,error")
						api.Send(msg)
					}
				}

			}()
			next(api, update)
		}
	}
}

// RateLimit is not usable
func RateLimit(log *zap.Logger, interval time.Duration) telegram.MiddlewareFunc {
	var mu sync.RWMutex
	lastMessageTimes := make(map[int64]time.Time)

	return func(next telegram.HandlerFunc) telegram.HandlerFunc {
		return func(api *tgbotapi.BotAPI, update tgbotapi.Update) {

			var chatID int64
			if update.Message != nil {
				chatID = update.Message.Chat.ID
			} else if update.CallbackQuery != nil {
				chatID = update.CallbackQuery.Message.Chat.ID
			}

			if chatID != 0 {
				mu.Lock()

				lastTime, exists := lastMessageTimes[chatID]
				now := time.Now()

				if exists && time.Since(lastTime) < interval {
					mu.Unlock()
					//logging
					return
				}

				lastMessageTimes[chatID] = now

				mu.Unlock()

			}
			next(api, update)
		}
	}
}
