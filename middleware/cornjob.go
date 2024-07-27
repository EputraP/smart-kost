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
	userRepo        repository.UserRepository
	userCurrentRepo repository.UserCurrentLocationRepo
}

type CornJobConfig struct {
	UserRepo        repository.UserRepository
	UserCurrentRepo repository.UserCurrentLocationRepo
}

func NewCorn(config CornJobConfig) CornJob {
	return &cornJob{
		userRepo:        config.UserRepo,
		userCurrentRepo: config.UserCurrentRepo,
	}
}

func (c cornJob) UpdateUserToOffline() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(5).Minutes().Do(func() {
		c.userRepo.UpdateUserOffline()
	})
	s.Every(15).Seconds().Do(func() {
		c.userCurrentRepo.UpdateLocationStatus()
	})
	s.StartAsync()
}
