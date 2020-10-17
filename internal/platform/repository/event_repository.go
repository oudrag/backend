package repository

import (
	"github.com/oudrag/server/internal/domain/events"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventRepository struct {
	db *mongo.Database
}

func NewEventRepository(db *mongo.Database) *EventRepository {
	return &EventRepository{db: db}
}
