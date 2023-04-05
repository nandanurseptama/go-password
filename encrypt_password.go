package password

type EncryptPassword interface {
	Encrypt(plainString string, key string) (*string, error)
	Decrypt(encryptedString string, key string) (*string, error)
}
