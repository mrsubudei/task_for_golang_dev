// Package hasher implements hashing password
package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
)

// interface -.
type PasswordHasher interface {
	Hash(salt, entered string) string
	CheckPassword(existHashed, salt, entered string) bool
}

// Md5Hasher -.
type Md5Hasher struct {
}

// NewMd5Hasher -.
func NewMd5Hasher() *Md5Hasher {
	return &Md5Hasher{}
}

// Hash -.
func (m *Md5Hasher) Hash(salt, entered string) string {
	h := md5.New()
	io.WriteString(h, salt)
	io.WriteString(h, entered)
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	return hashed
}

// CheckPassword -.
func (m *Md5Hasher) CheckPassword(existHashed, salt, entered string) bool {
	hashed := m.Hash(salt, entered)

	if existHashed != hashed {
		return false
	}

	return true
}
