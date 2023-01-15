package auth

import (
	"log"

	"github.com/alexedwards/argon2id"
)

type Authorization struct {
}

func NewAuthorization() (auth *Authorization) {
	auth = new(Authorization)

	return
}

func (*Authorization) CreateHash(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func (*Authorization) CheckHash(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
