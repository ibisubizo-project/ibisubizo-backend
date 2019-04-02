package users

import (
	"crypto/rand"
	"crypto/subtle"
	"fmt"
	"strings"

	"github.com/ofonimefrancis/problemsApp/config/must"
	"golang.org/x/crypto/scrypt"
)

const (
	SaltLen        = 32
	HashLen        = 64
	MinPasswordLen = 8
)

type WrongPasswordError string

func (self WrongPasswordError) Error() string {
	return string(self)
}

type PasswordHash struct {
	Hash []byte `json:"hash"`
	Salt []byte `json:"salt"`
}

func NewPasswordHash(password string) (*PasswordHash, error) {
	salt := generateSalt()
	hash, err := createPasswordHash(password, salt)
	if err != nil {
		return nil, err
	}
	return &PasswordHash{Hash: hash, Salt: salt}, nil
}

func generateSalt() []byte {
	salt := make([]byte, SaltLen)

	must.DoF(func() error {
		_, err := rand.Read(salt)
		return err
	})

	return salt
}

func createPasswordHash(password string, salt []byte) ([]byte, error) {
	password = strings.TrimSpace(password)

	if len(password) < MinPasswordLen {
		return nil, fmt.Errorf("The password is too short")
	}

	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, HashLen)

	if err != nil {
		return nil, err
	}
	return hash, nil
}

func VerifyPassword(password string, hash []byte, salt []byte) bool {
	newhash, err := createPasswordHash(password, salt)
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(newhash, hash) == 1
}
