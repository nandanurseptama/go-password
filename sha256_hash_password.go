package password

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

// SHA256HashPassword implements hashing password with SHA256 algorithm
type SHA256HashPassword struct {
}

// Hashing plain string to hashed string with SHA256 algorithm
//
// will return nil string and error if hashing fail
//
// will return hashed string and nil error is hashing success
func (h SHA256HashPassword) Hashing(plainString string) (*string, error) {
	m := sha256.New()
	_, err := io.WriteString(m, plainString)
	if err != nil {
		return nil, err
	}
	bs := m.Sum(nil)
	r := hex.EncodeToString(bs)

	return &r, nil
}
