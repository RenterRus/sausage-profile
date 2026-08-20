package otp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/pquerna/otp/totp"
)

type otpManager struct {
	// Ключ должен быть строго 16, 24 или 32 байта (для AES-128, 192 или 256)
	key    []byte
	issuer string
}

func NewOTPManager(key []byte, issuer string) OTP {
	return &otpManager{
		key:    key,
		issuer: issuer,
	}
}

func (o *otpManager) GenerateHash(username string) (string, string, error) {
	// Генерируем ключ
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      o.issuer,
		AccountName: username,
	})
	if err != nil {
		return "", "", err
	}

	// Secret() — строка для сохранения в БД (в зашифрованном виде!)
	secretKey, err := o.encrypt(key.Secret())
	if err != nil {
		return "", "", fmt.Errorf("GenerateHash.encrypt: %w", err)
	}

	// URL для построения QR-кода, который сканирует пользователь в приложении
	return secretKey, key.URL(), nil
}

func (o *otpManager) ValidateCode(passcode, secretKey string) (bool, error) {
	secret, err := o.decrypt(secretKey)
	if err != nil {
		return false, fmt.Errorf("ValidateCode.decrypt: %w", err)
	}

	// Валидирует код по текущему времени сервера
	return totp.Validate(passcode, secret), nil
}

// Encrypt шифрует строку и возвращает Base64-код
func (o *otpManager) encrypt(plainText string) (string, error) {
	block, err := aes.NewCipher(o.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherBytes := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// Decrypt расшифровывает Base64-строку обратно в текст
func (o *otpManager) decrypt(secureText string) (string, error) {
	cipherBytes, err := base64.StdEncoding.DecodeString(secureText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(o.key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(cipherBytes) < nonceSize {
		return "", fmt.Errorf("ошибка: текст слишком короткий")
	}

	nonce, cipherText := cipherBytes[:nonceSize], cipherBytes[nonceSize:]
	plainBytes, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		return "", err
	}

	return string(plainBytes), nil
}
