package password

type EncryptPassword interface {
	Encrypt(plainString string)
	Decrypt(encryptedString string)
}
