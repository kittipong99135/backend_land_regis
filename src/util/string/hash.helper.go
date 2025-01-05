package util_string

import (
	"golang.org/x/crypto/bcrypt"
)

type UtilsHash struct{}

func UseHashedString() UtilsHash {
	return UtilsHash{}
}

func (u UtilsHash) HashedString(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func (u UtilsHash) CompareString(hashStr, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashStr), []byte(str))
}
