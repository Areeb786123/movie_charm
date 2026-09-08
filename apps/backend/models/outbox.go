package models

import "time"

type OutboxEvent struct {
	ID          uint   `gorm:"primaryKey"`
	EventType   string `json:"event_type"`
	Payload     []byte `json:"payload"`
	AggregateID int64  `json:"aggregate_id"`
	CreatedAt   time.Time
	PublishedAt *time.Time
}
