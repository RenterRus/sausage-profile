package entity

import "time"

type User struct {
	UserLogin string
	Uuid      string
	OtpHash   string
	OtpLink   string
	CreatedAt time.Time
	Confirmed *bool
}
