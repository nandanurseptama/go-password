package password

// HashPassword interface of Hash Algorithm
type HashPassword interface {
	// abstract function of hashing
	// will be implemented on each hashing algorithm
	Hashing(plainString string) (*string, error)
}

var (
	HashMethodMD5    *MD5HashPassword
	HashMethodSHA1   *SHA1HashPassword
	HashMethodSHA256 *SHA256HashPassword
	HashMethodSHA512 *SHA512HashPassword
)

// initiate hash method
func init() {
	HashMethodMD5 = &MD5HashPassword{}
	HashMethodSHA256 = &SHA256HashPassword{}
	HashMethodSHA1 = &SHA1HashPassword{}
	HashMethodSHA512 = &SHA512HashPassword{}
}
