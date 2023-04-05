package password

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

type AESEncryptPassword struct {
	SecretKey string
}

// Encrypt string with AES-GCM
//
// Will encrypt plain string to encryted string in (hex format)
//
// if success will return encrypted string (hex format) and nil error
//
// if fail will return encrypted string nil and error
func (r AESEncryptPassword) Encrypt(plainString string) (*string, error) {
	key := []byte(r.SecretKey)

	c, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	plainText := []byte(plainString)

	aesGCM, err := cipher.NewGCM(c)

	nonce := make([]byte, aesGCM.NonceSize())

	cipherTextByte := aesGCM.Seal(nil, nonce, plainText, nil)

	cipherText := hex.EncodeToString(cipherTextByte)
	return &cipherText, nil

}

// Decrypt string with AES-GCM
//
// Will encrypt plain string to encryted string in (hex format)
//
// if success will return encrypted string (hex format) and nil error
//
// if fail will return encrypted string nil and error
func (r AESEncryptPassword) Decrypt(encryptedString string) (*string, error) {
	encryptedStringBytes, err := hex.DecodeString(encryptedString)

	if err != nil {
		return nil, err
	}

	c, err := aes.NewCipher([]byte(r.SecretKey))
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(c)

	nonceSize := aesGCM.NonceSize()
	nonce := make([]byte, nonceSize)

	plainTextByte, err := aesGCM.Open(nil, nonce, encryptedStringBytes, nil)
	if err != nil {
		return nil, err
	}

	plainText := fmt.Sprintf("%s", plainTextByte)
	return &plainText, nil

}
