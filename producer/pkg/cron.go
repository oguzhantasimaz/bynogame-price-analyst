package pkg

import (
	"github.com/go-co-op/gocron"
	log "github.com/sirupsen/logrus"
	"time"
)

type CronClient interface {
	Start()
	Schedule(interval string, fn func())
	Stop()
}

type cronClient struct {
	sch *gocron.Scheduler
}

func NewCronClient() CronClient {
	return &cronClient{
		sch: gocron.NewScheduler(time.UTC),
	}
}

func (c *cronClient) Start() {
	c.sch.StartAsync()
}

func (c *cronClient) Schedule(interval string, fn func()) {
	_, err := c.sch.Every(interval).Do(func() {
		fn()
	})
	if err != nil {
		log.Error("Error scheduling cron " + err.Error())
	}
}

func (c *cronClient) Stop() {
	c.sch.Stop()
}
