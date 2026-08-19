package otp

type OTP interface {
	GenerateHash(username string) (string, string, error)
	ValidateCode(passcode, secretKey string) bool
	encrypt(plainText string) (string, error)
	decrypt(secureText string) (string, error)
}
