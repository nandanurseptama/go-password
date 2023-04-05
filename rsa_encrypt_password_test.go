package password_test

import (
	"testing"

	password "github.com/nandanurseptama/go-password"
)

var privk = "-----BEGIN RSA PRIVATE KEY-----\n" +
	"MIICXQIBAAKBgQCq1hzaz+09WetZR1E0W38YjPYNeCa9BplRjmZu9J1MnxFo/G3B" +
	"zvlG47ZiRT8C628ZId+PX9SDT8PYLwHo8cqqhKKIwIGmosgnDRHo19rMrV41roFq" +
	"GbSVv6PdUdVLYPpl4EZ1OZ7XUPGuXm60c7fSohAHjq+j0JdL+scpbFo6XwIDAQAB" +
	"AoGABWgFPb96yIhB9G9XWIrUuNgtKTv2LeE5lOUmxLglCjm2eVYTuyBrrxyhtvZu" +
	"Cg+dUnRJ+OWVehpaHktOiiqsuS94pOw42sYAuymAyDXckppOL+O54vEuOGV54DK3" +
	"buAyl6X69CB/RZ4xGtttfqKhDXEK0B/yrnPoQPc35yfzWMECQQDgMv61pL8u53Ut" +
	"YqVIKXei6DuBjjWOjWGDSBhbQUhYAU1gL4N4fmuSKDaXqe/shDdLnyhuijW6+oNw" +
	"SkZlFSvvAkEAwxFvRiNiiG3CI1UxdewHaqQd+1uMDJ1CGQXmdr9DgzzBiGikzM20" +
	"4j8X7Tx9altX7nqaJ/JHgCvbnWOpCHEokQJBAL0J4oCQP++xk7jH82sMI+cFf582" +
	"pGvlQ/Jo6LiRLAmgV3iht23y0orzQ0zQKg+4T9OAiJvkB3f34LnetspmdpsCQAdA" +
	"KqaswgjYK5Mso34Cm/P2zbZ+HKKLZd2QPpTDXzsNkxQC7709GDAgsD79NJwzcP1I" +
	"SVqawtWcy7j571qE+kECQQCWYR3yUwAZmXoSKfh6F+w/t20r4VL1iZH5sl/Xn0F9" +
	"PBOFPj/L4pZnpZgRw44m2OofjOSzXTBFX8VvbA+IQlcL" +
	"\n-----END RSA PRIVATE KEY-----"

func TestEncrypt(t *testing.T) {
	rsaEncryptPassword := password.RSAEncryptPassword{
		Pk: privk,
	}
	res, err := rsaEncryptPassword.Encrypt("wkwk")
	if res == nil || err != nil {
		t.Fatalf("err %v", err)
	}
}

func TestDecrypt(t *testing.T) {
	rsaEncryptPassword := password.RSAEncryptPassword{
		Pk: privk,
	}
	res, err := rsaEncryptPassword.Decrypt("28a33fa9c5dcfb0f12e602424034cf62850c23394be9208b9fec5ead9d939052305fbef91c12001470bc09ce47b4fa047391c14bfd26bd20cc9713dbece7d247776ce7696d6429bda0f6e1a61fefb2dde98e8e4a73f3fb92021c7b647eb230af626c1367ea40dd8356278d1b648911cfb675585c946fce2364cb3c445b962d7a")

	if res == nil || err != nil {
		t.Fatalf("err %v", err)
	}
}
