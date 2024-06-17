package password

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "mysecretpassword"

	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Ensure the hash is valid by comparing it with the original password
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	assert.NoError(t, err)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Test that the correct password returns true
	isValid := CheckPasswordHash(password, hash)
	assert.True(t, isValid)

	// Test that an incorrect password returns false
	isValid = CheckPasswordHash("wrongpassword", hash)
	assert.False(t, isValid)
}
