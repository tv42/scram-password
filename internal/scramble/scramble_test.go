package scramble

import (
	"testing"

	"github.com/xdg-go/scram"
)

func TestSimple(t *testing.T) {
	salt := [32]byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	}
	hash, err := hashWithKF("jdoe", "s3kr1t", scram.KeyFactors{
		Iters: 4096,
		Salt:  string(salt[:]),
	})
	if err != nil {
		t.Fatal(err)
	}
	const want = `SCRAM-SHA-256$4096:AAECAwQFBgcAAQIDBAUGBwABAgMEBQYHAAECAwQFBgc=$3OKulhqxk9w6FbPtpUHCuIkEsW+2F2cjX0/ABNgYsbI=:BZ55glbzmkm4V5VjvpHHENWSEZE/IVxZWuAqeLUsikQ=`
	if g, e := hash, want; g != e {
		t.Errorf("wrong hash:\n\tgot\t%s\n\twant\t%s\n", g, e)
	}
}
