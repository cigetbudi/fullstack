package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("nickname sudah terdaftar")
	}

	if strings.Contains(err, "email") {
		return errors.New("email sudah terdaftar")
	}

	if strings.Contains(err, "title") {
		return errors.New("title sudah terdaftar")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("password salah")
	}
	return errors.New("detail salah")
}
