package usecase

import (
	"context"
	"fmt"

	"github.com/RenterRus/sausage-profile/internal/entity"
	"github.com/RenterRus/sausage-profile/internal/repo/otp"
	"github.com/RenterRus/sausage-profile/internal/repo/psql/db"
)

type register struct {
	otpRepo   otp.OTP
	usersRepo db.Querier
}

func NewRegisterManager(otpRepo otp.OTP, usersRepo db.Querier) Register {
	return &register{
		otpRepo:   otpRepo,
		usersRepo: usersRepo,
	}
}

func (r *register) Registration(ctx context.Context, login string) (string, error) {
	hash, url, err := r.otpRepo.GenerateHash(login)
	if err != nil {
		return "", nil
	}

	if err := r.usersRepo.Register(ctx, db.RegisterParams{
		Login: &login,
		Hash:  &hash,
		Link:  &url,
	}); err != nil {
		return "", fmt.Errorf("Registration.Register: %w", err)
	}

	return url, nil
}

func (r *register) Confirmed(ctx context.Context, login, code string) error {
	hash, err := r.usersRepo.Hash(ctx, &login)
	if err != nil {
		return fmt.Errorf("Confirmed.Hash: %w", err)
	}

	isValid, err := r.otpRepo.ValidateCode(code, hash)
	if err != nil {
		return fmt.Errorf("Confirmed.ValidateCode: %w", err)
	}

	if isValid {
		return fmt.Errorf("Confirmed.ValidateCode(invalid): %w", entity.ErrCodeInvalid)
	}

	if err = r.usersRepo.Confirmed(ctx, &login); err != nil {
		return fmt.Errorf("Confirmed.Confirmed: %w", err)
	}

	return nil

}
