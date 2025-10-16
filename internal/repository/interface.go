package repository

import (
	"context"
	"time"
)

type UserProvider interface {
	CountUserCreated(ctx context.Context, time time.Time) (int64, error)
	CountUserUpdated(ctx context.Context, time time.Time) (int64, error)
	CountUserDeleted(ctx context.Context, time time.Time) (int64, error)
}
