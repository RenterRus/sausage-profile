package usecase

import "context"

type Register interface {
	Registration(ctx context.Context, login string) (string, error)
	Confirmed(ctx context.Context, login, code string) error
}
