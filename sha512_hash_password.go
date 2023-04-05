package password

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
)

// SHA512HashPassword implements hashing password with SHA256 algorithm
type SHA512HashPassword struct {
}

// Hashing plain string to hashed string with SHA512 algorithm
//
// will return nil string and error if hashing fail
//
// will return hashed string and nil error is hashing success
func (h SHA512HashPassword) Hashing(plainString string) (*string, error) {
	m := sha512.New()
	_, err := io.WriteString(m, plainString)
	if err != nil {
		return nil, err
	}
	bs := m.Sum(nil)
	r := hex.EncodeToString(bs)

	return &r, nil
}
