// Package hasher implements hashing password
package hasher

import (
	"crypto/md5"
	"fmt"
	"io"
)

// interface -.
type PasswordHasher interface {
	Hash(salt, entered string) (string, error)
	CheckPassword(existHashed, salt, entered string) (bool, error)
}

// Md5Hasher -.
type Md5Hasher struct {
}

// NewMd5Hasher -.
func NewMd5Hasher() *Md5Hasher {
	return &Md5Hasher{}
}

// Hash -.
func (m *Md5Hasher) Hash(salt, entered string) (string, error) {
	h := md5.New()
	_, err := io.WriteString(h, salt)
	if err != nil {
		return "", fmt.Errorf("hasher - Hash - WriteString #1: %w", err)
	}
	_, err = io.WriteString(h, entered)
	if err != nil {
		return "", fmt.Errorf("hasher - Hash - WriteString #2: %w", err)
	}
	hashed := fmt.Sprintf("%x", h.Sum(nil))

	return hashed, nil
}

// CheckPassword -.
func (m *Md5Hasher) CheckPassword(existHashed, salt, entered string) (bool, error) {
	hashed, err := m.Hash(salt, entered)
	if err != nil {
		return false, fmt.Errorf("CheckPassword - %w", err)
	}

	return existHashed == hashed, nil
}
