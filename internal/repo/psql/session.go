package psql

import (
	"context"
	"fmt"

	"github.com/RenterRus/sausage-profile/internal/entity"
	"github.com/RenterRus/sausage-profile/internal/repo/psql/db"
)

func (u *UserRepo) GetRefreshToken(ctx context.Context, login *string) (entity.GetRefreshTokenRow, error) {
	if login == nil || *login == "" {
		return entity.GetRefreshTokenRow{}, fmt.Errorf("GetRefreshToken: %w", entity.ErrParametrNoFound)
	}

	resp, err := u.Queries.GetRefreshToken(ctx, login)
	if err != nil {
		return entity.GetRefreshTokenRow{}, fmt.Errorf("GetRefreshToken.GetRefreshToken: %w", err)
	}

	return entity.GetRefreshTokenRow{
		RefreshHash: resp.RefreshHash,
		IsExpired:   resp.IsExpired,
		Block:       resp.Block,
	}, nil
}

func (u *UserRepo) SetBlockRefresh(ctx context.Context, refreshHash *string) error {
	if refreshHash == nil || *refreshHash == "" {
		return fmt.Errorf("SetBlockRefresh: %w", entity.ErrParametrNoFound)
	}

	if err := u.Queries.SetBlockRefresh(ctx, refreshHash); err != nil {
		return fmt.Errorf("SetBlockRefresh.SetBlockRefresh: %w", err)
	}

	return nil
}

func (u *UserRepo) SetRefreshHash(ctx context.Context, arg entity.SetRefreshHashParams) error {
	if err := u.Queries.SetRefreshHash(ctx, db.SetRefreshHashParams{
		RefreshHash: arg.RefreshHash,
		Login:       arg.Login,
		ExpiredAt:   arg.ExpiredAt,
	}); err != nil {
		return fmt.Errorf("SetRefreshHash.SetRefreshHash: %w", err)
	}

	return nil
}
