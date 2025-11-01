package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 异常追加
func AppendError(exitError, newError error) error {
	if newError == nil {
		return exitError
	}

	if exitError == nil {
		return newError
	} else {
		return fmt.Errorf("%v,%w", exitError, newError)
	}
}

func Encrypt(value string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), err
}

func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
