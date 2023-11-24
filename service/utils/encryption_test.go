package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const MySecret string = "abc&1*~#^2^#s0^=)^^7%b34"
const text string = "password"

func TestEncryptDecrypt(t *testing.T) {
	encrypted, err := Encrypt(text, MySecret)
	assert.Nil(t, err)
	assert.NotEmpty(t, encrypted)
	decrypted, err := Decrypt(encrypted, MySecret)
	assert.Nil(t, err)
	assert.NotEmpty(t, decrypted)
	assert.Equal(t, decrypted, text)
}
