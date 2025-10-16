package app

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/P04KA/statistics/config"
	"github.com/P04KA/statistics/internal/handler"
	"github.com/P04KA/statistics/internal/repository"
	"github.com/P04KA/statistics/internal/storage"
	"github.com/P04KA/statistics/internal/usecase"
)

type App struct {
	StatsHandler *handler.UserStatsServer
	conn         *pgxpool.Pool
}

func New() *App {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Printf("Failed to load config: %v", err) // Добавьте лог
		return nil
	}

	conn, err := storage.GetConnect(cfg.DB.DBURL)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err) // Добавьте лог
		return nil
	}

	repo := repository.New(conn)
	statsUC := usecase.NewStatsUseCase(repo)
	statsHandler := handler.NewUserStatsServer(statsUC)

	return &App{
		StatsHandler: statsHandler,
		conn:         conn,
	}
}

func (a *App) GetStatsHandler() *handler.UserStatsServer {
	if a == nil || a.StatsHandler == nil { // Добавьте проверку на StatsHandler
		log.Fatal("app dont have handler")
		return nil
	}
	return a.StatsHandler
}

func (a *App) Close() {
	a.conn.Close()
}
