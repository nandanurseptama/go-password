
# Go Password

A module that implement hashing and (encrypt/decrypt) algorithms

Current supported algorithms :


- Hash
    - MD5
    - SHA1
    - SHA256
    - SHA512
## Authors

- [@nandanurseptama](https://www.github.com/nandanurseptama)


## Installation Guidelines

Go get

```bash
  go get -u github.com/nandanurseptama/go-password
```
Import

```go
  import "github.com/nandanurseptama/go-password"
```

## Examples

```go
import "github.com/nandanurseptama/go-password"

plainString := "wkwk"

# return `0bef1939b3c02eea4b89f1a8247419cf`, nil
md5HashedString, err := HashMethodMD5.Hashing(plainString)

```

## License

This code is licensed under the MIT license.
