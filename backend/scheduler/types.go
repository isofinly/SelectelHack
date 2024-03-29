package scheduler

import (
	"time"
)

type RedisConfig struct {
	Host         string
	DbID         int
	PollInterval uint
}

type Task struct {
	Name          string    `json:"name"`
	BeginTime     time.Time `json:"begin_time"`
	Payload       any       `json:"payload"`
	Type          TaskType  `json:"type"`
	Repeating     bool      `json:"repeating"`
	Interval      uint64    `json:"interval"`
	NextExecution time.Time `json:"next_execution"`
}

type TaskType string

const (
	NotificationSchedule TaskType = "report:schedule"
)

type NotificationTaskPayload struct {
	ChatID  uint   `json:"chat_id"`
	Message string `json:"message"`
	PlanId  uint   `json:"plan_id"`
}
