package tempest

import (
	"context"
	"time"
)

type MessageRepo interface {
	LoadMessages(ctx context.Context, start, end time.Time, msg Type)
	SaveMessage(ctx context.Context, message WeatherMessage)
	DeleteMessages(ctx context.Context, start, end time.Time, msg Type)
}
