package jobs

import (
	"tg-bot-template/pkg/scheduler"
	"tg-bot-template/pkg/storage"
	"tg-bot-template/pkg/telegram"

	"go.uber.org/zap"
)

func RegisterJobs(s *scheduler.Scheduler, bot *telegram.Bot, store storage.Storage, log *zap.Logger) {
	err := s.AddTask("*/10 * * * * *", func() {})
	if err != nil {
		log.Error("error adding task:", zap.Error(err))
	}

}
