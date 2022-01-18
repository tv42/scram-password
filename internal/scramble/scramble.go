package scramble

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/xdg-go/scram"
)

func makeSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}

func hashWithKF(username string, password string, kf scram.KeyFactors) (string, error) {
	// We could expose this as a command-line flag, but first need a use case we can test against.
	const authID = ""
	// We could make the algorithm a command-line flag.
	client, err := scram.SHA256.NewClient(username, password, authID)
	if err != nil {
		return "", err
	}

	credentials := client.GetStoredCredentials(kf)

	// SCRAM-SHA-256$<iter>:<salt>$<StoredKey>:<ServerKey>
	hashed := fmt.Sprintf("SCRAM-SHA-256$%d:%s$%s:%s",
		credentials.Iters,
		base64.StdEncoding.EncodeToString([]byte(credentials.Salt)),
		base64.StdEncoding.EncodeToString([]byte(credentials.StoredKey)),
		base64.StdEncoding.EncodeToString([]byte(credentials.ServerKey)),
	)
	return hashed, nil
}

func Hash(username string, password string) (string, error) {
	// We could take a known salt (as base64) as a command-line flag.

	// We could take salt size as a command-line flag.
	//
	// Postgres 14 uses salt size 16.
	// We'd rather be ahead of the curve than behind.
	const saltSize = 24
	salt, err := makeSalt(saltSize)
	if err != nil {
		return "", err
	}
	kf := scram.KeyFactors{
		Salt: string(salt),
		// We could take iterations as a command-line flag.
		Iters: 4096,
	}
	return hashWithKF(username, password, kf)
}
