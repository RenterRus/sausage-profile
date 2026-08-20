package entity

import "time"

const (
	STATUS_OK     = "ok"
	STATUS_FAILED = "failed"
)

type User struct {
	UserLogin string
	Uuid      string
	OtpHash   string
	OtpLink   string
	CreatedAt time.Time
	Confirmed *bool
}
