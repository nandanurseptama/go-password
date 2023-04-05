package password

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
)

// SHA1HashPassword implements hashing password with SHA1 algorithm
type SHA1HashPassword struct {
}

// Hashing plain string to hashed string with SHA1 algorithm
//
// will return nil string and error if hashing fail
//
// will return hashed string and nil error is hashing success
func (h SHA1HashPassword) Hashing(plainString string) (*string, error) {
	m := sha1.New()
	_, err := io.WriteString(m, plainString)
	if err != nil {
		return nil, err
	}
	bs := m.Sum(nil)
	r := hex.EncodeToString(bs)

	return &r, nil
}
