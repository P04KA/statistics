package handler

import (
	"context"

	"github.com/P04KA/statistics/internal/usecase"
	"github.com/P04KA/statistics/pkg/stats"
)

type UserStatsServer struct {
	stats.UnimplementedUserStatsServiceServer
	statUC usecase.StatsUseCase
}

func NewUserStatsServer(statUC usecase.StatsUseCase) *UserStatsServer {
	return &UserStatsServer{
		statUC: statUC,
	}
}

func (u *UserStatsServer) GetUserStats(ctx context.Context, req *stats.UserStatsRequest) (*stats.UserStatsResponse, error) {
	return u.statUC.GetUserStats(ctx, req.GetPeriod())
}
