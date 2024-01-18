package user

import (
	"context"
	"log"
	"time"

	"event/internal/pkg/event"
)

const Deleted event.Name = "user.deleted"

type DeletedEvent struct {
	Time time.Time
	ID   string
	Who  string
}

func (e DeletedEvent) Handle(ctx context.Context) {
	log.Printf("deleting: %+v\n", e)
}
