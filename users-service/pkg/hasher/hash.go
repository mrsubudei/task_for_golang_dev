package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
)

type PasswordHasher interface {
	Hash(salt, entered string) string
	CheckPassword(existHashed, salt, entered string) bool
}

type Md5Hasher struct {
}

func NewMd5Hasher() *Md5Hasher {
	return &Md5Hasher{}
}

func (m *Md5Hasher) Hash(salt, entered string) string {
	h := md5.New()
	io.WriteString(h, salt)
	io.WriteString(h, entered)
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	return hashed
}

func (m *Md5Hasher) CheckPassword(existHashed, salt, entered string) bool {
	hashed := m.Hash(salt, entered)

	if existHashed != hashed {
		return false
	}

	return true
}
