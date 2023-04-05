package password_test

import (
	"testing"

	"github.com/nandanurseptama/go-password"
)

func TestHashing(t *testing.T) {
	plainString := "wkwk"
	expectedM5HashedString := "0bef1939b3c02eea4b89f1a8247419cf"

	md5HashedString, err := password.HashMethodMD5.Hashing(plainString)
	if expectedM5HashedString != *md5HashedString || err != nil {
		t.Fatalf(`Hashing("wkwk") = %v, %v, want match for %#q, nil`, *md5HashedString, err, expectedM5HashedString)
	}
	expectedSha256String := "4499c41eec361a4d8c208b5da66870e1f0ee57ef2cc6fd80d0df5fc9d81b7682"
	sha256HashedString, err := password.HashMethodSHA256.Hashing(plainString)

	if expectedSha256String != *sha256HashedString || err != nil {
		t.Fatalf(`Hashing("wkwk") = %v, %v, want match for %#q, nil`, *sha256HashedString, err, expectedSha256String)
	}

	expectedSha1String := "0d0de34f2b52f8f2c28f9014e8a848e9dda8b3a6"
	sha1HashedString, err := password.HashMethodSHA1.Hashing(plainString)

	if expectedSha1String != *sha1HashedString || err != nil {
		t.Fatalf(`Hashing("wkwk") = %v, %v, want match for %#q, nil`, *sha1HashedString, err, expectedSha1String)
	}

	expectedSha512String := "ca46d60713dbe128ae701fba550300d943dcaeffe1d83248cc041ad65c35230cc79146a77b48edaaa9c47478fd3a7bb32ba7bc14597068f55d236fa54a194efb"
	sha512HashedString, err := password.HashMethodSHA512.Hashing(plainString)

	if expectedSha512String != *sha512HashedString || err != nil {
		t.Fatalf(`Hashing("wkwk") = %v, %v, want match for %#q, nil`, *sha512HashedString, err, expectedSha512String)
	}

}
