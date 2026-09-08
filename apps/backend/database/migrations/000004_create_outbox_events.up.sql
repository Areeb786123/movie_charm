CREATE TABLE outbox_events (
    id BIGSERIAL PRIMARY KEY,
    event_type VARCHAR(255) NOT NULL,
    aggregate_id BIGINT NOT NULL,
    payload JSONB NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    published_at TIMESTAMP NULL
);

CREATE INDEX idx_outbox_events_unpublished
ON outbox_events (created_at)
WHERE published_at IS NULL;