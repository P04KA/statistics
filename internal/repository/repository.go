package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *UserRepo {
	return &UserRepo{conn: conn}
}

func (r *UserRepo) CountUserCreated(ctx context.Context, time time.Time) (int64, error) {
	var count int64
	err := r.conn.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE created_at > $1 AND deleted_at IS NULL`, time).Scan(&count)
	return count, err
}

func (r *UserRepo) CountUserUpdated(ctx context.Context, time time.Time) (int64, error) {
	var count int64
	err := r.conn.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE updated_at > $1 AND created_at != updated_at AND deleted_at IS NULL`, time).Scan(&count)
	return count, err
}

func (r *UserRepo) CountUserDeleted(ctx context.Context, time time.Time) (int64, error) {
	var count int64
	err := r.conn.QueryRow(ctx, `SELECT COUNT(*) FROM users WHERE deleted_at > $1`, time).Scan(&count)
	return count, err
	//что будет возвращать queryRow, если нет записей. мб pgx вернется. ВОЗВРАЩАЕТ 0
}

//запилить новую репу statistics, там запускаю сервер, с соблюдение архитектуры user stats server - usecase, его надо запихнуть в internal/usecase
// в statistics gateway, клиент(убрать его в Handle и добавить зависимость по типу internal/usecase/stat.go)
// stats.prorto запихнуть в каждый репозиторий
//запилить новый репозиторий all-in-one, там пропаисать docker compose, который будет поднимать весь сервис/проект(там же конфиги для графаны)
