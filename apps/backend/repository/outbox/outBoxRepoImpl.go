package outbox

import (
	"backend/models"
	"context"
	"time"

	"gorm.io/gorm"
)

type outboxRepository struct {
	db *gorm.DB
}

func CreateOutboxRepository(db *gorm.DB) KafkaOutBoxRepo {
	return &outboxRepository{
		db: db,
	}
}

func (r *outboxRepository) Create(ctx context.Context, tx *gorm.DB, event *models.OutboxEvent) error {
	return tx.WithContext(ctx).Create(event).Error
}

func (r *outboxRepository) GetPendingEvents(ctx context.Context, limit int) ([]models.OutboxEvent, error) {
	var data []models.OutboxEvent
	err := r.db.WithContext(ctx).
		Where("published_at IS NULL").
		Order("created_at ASC").
		Limit(limit).
		Find(&data).Error
	return data, err
}

func (r *outboxRepository) MarkPublished(ctx context.Context, id uint) error {
	now := time.Now()
	return r.db.WithContext(ctx).
		Model(&models.OutboxEvent{}).
		Where("id = ?", id).
		Update("published_at", &now).
		Error
}
