package middleware

import (
	"smart-kost-backend/repository"
	"time"

	"github.com/go-co-op/gocron"
)

type CornJob interface {
	UpdateUserToOffline()
}

type cornJob struct {
	userRepo repository.UserRepository
}

type CornJobConfig struct {
	UserRepo repository.UserRepository
}

func NewCorn(config CornJobConfig) CornJob {
	return &cornJob{
		userRepo: config.UserRepo,
	}
}

func (c cornJob) UpdateUserToOffline() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Minutes().Do(func() {
		c.userRepo.UpdateUserOffline()
		println("test1")
	})
	s.StartAsync()
}
