package password

import (
	"crypto/md5"
	"fmt"
	"io"
)

// MD5HashPassword implements hashing password with MD5 algorithm
type MD5HashPassword struct {
}

// Hashing plain string to hashed string with MD5 algorithm
//
// will return nil string and error if hashing fail
//
// will return hashed string and nil error is hashing success
func (h MD5HashPassword) Hashing(plainString string) (*string, error) {
	m := md5.New()
	_, err := io.WriteString(m, plainString)
	if err != nil {
		return nil, err
	}
	bs := m.Sum(nil)
	r := fmt.Sprintf("%x", bs)
	return &r, nil
}
