package scheduler

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler struct {
	cron *cron.Cron
	log  *zap.Logger
}

func New(logger *zap.Logger) *Scheduler {
	c := cron.New(cron.WithSeconds())
	return &Scheduler{
		cron: c,
		log:  logger,
	}
}

func (s *Scheduler) AddTask(schedule string, task func()) error {
	_, err := s.cron.AddFunc(schedule, task)
	return err
}

func (s *Scheduler) Run() {
	s.cron.Start()
	s.log.Info("Scheduler has started")
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	s.log.Info("Scheduler has stopped")
}
