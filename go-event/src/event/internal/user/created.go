package user

import (
	"context"
	"log"
	"time"

	"event/internal/pkg/event"
)

const Created event.Name = "user.created"

type CreatedEvent struct {
	Time time.Time
	ID   string
}

func (e CreatedEvent) Handle(ctx context.Context) {
	log.Printf("creating: %+v\n", e)
}
