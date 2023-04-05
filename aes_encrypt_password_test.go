package password_test

import (
	"fmt"
	"testing"

	password "github.com/nandanurseptama/go-password"
)

var secretKey = "1234567812345678"
var plainText = "wkwk"
var encrypted = "5aebdc5fdc2182cfba9cbbda60c5831e462e208f"

func TestAesEncrypt(t *testing.T) {
	encryptor := password.AESEncryptPassword{
		SecretKey: secretKey,
	}
	result, err := encryptor.Encrypt(plainText)
	if err != nil || result == nil || *result != encrypted {
		t.Fail()
	}

	t.Cleanup(func() {

	})
}
func TestAesDecrypt(t *testing.T) {
	encryptor := password.AESEncryptPassword{
		SecretKey: secretKey,
	}
	result, err := encryptor.Decrypt(encrypted)
	if err != nil || result == nil || *result != plainText {
		fmt.Println(err)
		t.Fail()
	}

	t.Cleanup(func() {

	})
}
