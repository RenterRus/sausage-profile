package entity

import "github.com/jackc/pgx/v5/pgtype"

type RegisterParams struct {
	Login *string
	Hash  *string
	Link  *string
}

type GetRefreshTokenRow struct {
	RefreshHash *string
	IsExpired   bool
	Block       bool
}

type SetRefreshHashParams struct {
	RefreshHash *string
	ExpiredAt   pgtype.Timestamp
	Login       *string
}
