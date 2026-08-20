package psql

import (
	"context"

	"github.com/RenterRus/sausage-profile/internal/repo/psql/db"
)

type UsersRepo interface {
	Confirmed(ctx context.Context, login *string) error
	Hash(ctx context.Context, login *string) (string, error)
	IsExist(ctx context.Context, login *string) (bool, error)
	Register(ctx context.Context, arg db.RegisterParams) error
}
