package user

import (
	"context"
	"log"
	"time"

	"event/internal/pkg/event"
)

const Updated event.Name = "user.updated"

type UpdatedEvent struct {
	Time     time.Time
	ID       string
	Key      string
	OldValue string
	NewValue string
}

func (e UpdatedEvent) Handle(ctx context.Context) {
	log.Printf("updating: %+v\n", e)
}
