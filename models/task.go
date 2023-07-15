package models

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type TaskStatus int

const (
	Done TaskStatus = iota
	InProgress
	NotNedeed
	ReadyToStart
	Backlog
)

func (status TaskStatus) String() string {
	options := [...]string{
		"Done",
		"In Progress",
		"Not Nedeed",
		"Ready To Start",
		"Backlog",
	}

	if status < Done || status > Backlog {
		return "Unknown"
	}

	return options[status]

}

func getRandomTaskStatus() TaskStatus {
	rand.Seed(time.Now().UnixNano())
	min := int(Done)
	max := int(Backlog)

	randomStatus := rand.Intn(max-min+1) + min

	return TaskStatus(randomStatus)
}

func getRandomTaskStatusString() string {
	randomStatus := getRandomTaskStatus().String()
	return randomStatus
}

type Task struct {
	id      uuid.UUID
	name    string
	parents []uuid.UUID
}

func NewTask(id uuid.UUID, name string, parents []uuid.UUID) Task {
	return Task{
		id:      id,
		name:    name,
		parents: parents,
	}
}
