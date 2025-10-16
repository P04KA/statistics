package usecase

import (
	"context"
	"time"

	"github.com/P04KA/statistics/internal/repository"
	"github.com/P04KA/statistics/pkg/stats"
)

type StatsUseCase interface {
	GetUserStats(ctx context.Context, period string) (*stats.UserStatsResponse, error)
}

type statsUseCase struct {
	repo repository.UserProvider
}

func NewStatsUseCase(repo repository.UserProvider) StatsUseCase {
	return &statsUseCase{repo: repo}
}

func (uc *statsUseCase) GetUserStats(ctx context.Context, period string) (*stats.UserStatsResponse, error) {
	since := uc.calculateSinceTime(period)

	created, err := uc.repo.CountUserCreated(ctx, since)
	if err != nil {
		return nil, err
	}

	updated, err := uc.repo.CountUserUpdated(ctx, since)
	if err != nil {
		return nil, err
	}

	deleted, err := uc.repo.CountUserDeleted(ctx, since)
	if err != nil {
		return nil, err
	}

	// Используем напрямую gRPC структуру
	return &stats.UserStatsResponse{
		UsrCreated: created,
		UsrUpdated: updated,
		UsrDeleted: deleted,
		Period:     period,
	}, nil
}

func (uc *statsUseCase) calculateSinceTime(period string) time.Time {
	now := time.Now()
	switch period {
	case "hour":
		return now.Add(-1 * time.Hour)
	case "day":
		return now.Add(-24 * time.Hour)
	case "week":
		return now.Add(-7 * 24 * time.Hour)
	default:
		return now.Add(-24 * time.Hour)
	}
}
