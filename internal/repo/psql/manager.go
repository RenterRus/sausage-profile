package psql

import (
	"context"
	"fmt"

	"github.com/RenterRus/sausage-profile/internal/repo/psql/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	Queries *db.Queries
}

//connStr := "postgres://username:password@localhost:5432/database_name"

func NewDBManager(conn string) (db.Querier, error) {
	pgx, err := pgxpool.New(context.Background(), conn)
	if err != nil {
		return nil, fmt.Errorf("NewDBManager.New: %w", err)
	}

	return &UserRepo{
		Queries: db.New(pgx),
	}, nil
}
