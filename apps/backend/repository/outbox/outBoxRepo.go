package outbox

import (
	"backend/models"
	"context"

	"gorm.io/gorm"
)

type KafkaOutBoxRepo interface {
	Create(ctx context.Context, tx *gorm.DB, event *models.OutboxEvent) error
	GetPendingEvents(ctx context.Context, limit int) ([]models.OutboxEvent, error)
	MarkPublished(ctx context.Context, id uint) error
}
