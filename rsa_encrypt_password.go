package password

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"strings"
)

// RSAEncryptPassword is implementation of RSA Encryption
type RSAEncryptPassword struct {
	Pk string
}

// Get Private Key
func (r RSAEncryptPassword) privateKey() (*rsa.PrivateKey, error) {
	x := strings.NewReader(r.Pk)
	pemBytes, err := ioutil.ReadAll(x)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}
	enc := x509.IsEncryptedPEMBlock(block)
	b := block.Bytes
	if enc {
		b, err = x509.DecryptPEMBlock(block, nil)
		if err != nil {
			return nil, err
		}
	}
	key, err := x509.ParsePKCS1PrivateKey(b)
	if err != nil {
		return nil, err
	}
	return key, nil
}

// Encrypt
//
// Will encrypt plain string to encryted string in (hex format)
//
// if success will return encrypted string (hex format) and nil error
//
// if fail will return encrypted string nil and error
func (r RSAEncryptPassword) Encrypt(plainString string) (*string, error) {
	pk, err := r.privateKey()
	if err != nil {
		return nil, err
	}
	pb := pk.PublicKey
	hash := sha256.New()
	rand.Reader.Read([]byte("A"))
	result, err := rsa.EncryptOAEP(hash, rand.Reader, &pb, []byte(plainString), nil)
	if err != nil {
		return nil, err
	}
	encryptedText := hex.EncodeToString(result)
	return &encryptedText, nil
}

// Decrypt
//
// Will Decrypt encrypted string (hex format) to plain string
//
// if success will return plain string and nil error
//
// if fail will return nil string and error
func (r RSAEncryptPassword) Decrypt(encryptedString string) (*string, error) {
	encryptedStringBytes, err := hex.DecodeString(encryptedString)
	if err != nil {
		return nil, err
	}

	pk, err := r.privateKey()
	if err != nil {
		return nil, err
	}
	hash := sha256.New()
	result, err := rsa.DecryptOAEP(hash, rand.Reader, pk, encryptedStringBytes, nil)
	if err != nil {
		return nil, err
	}
	x := string(result)
	return &x, nil
}
