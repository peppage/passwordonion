package passwordonion

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Encrypt a password with sha512, bcrypt then aes256 with a random iv
func Encrypt(pepper, password string) ([]byte, error) {

	// First sha512 to take care of any issues with bcrypt.
	hash := sha512.New()
	passwordSha := hash.Sum([]byte(password))

	passwordBcrypt, err := bcrypt.GenerateFromPassword(passwordSha, 11)
	if err != nil {
		return []byte{}, err
	}

	block, err := aes.NewCipher([]byte(pepper))
	if err != nil {
		return []byte{}, err
	}

	cipherText := make([]byte, aes.BlockSize+len(passwordBcrypt))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], passwordBcrypt)

	return cipherText, nil
}

// Compare decrypts the storedPassword and returns if it matches password
func Compare(pepper, storedPassword, password string) error {
	p := []byte(storedPassword)
	block, err := aes.NewCipher([]byte(pepper))
	if err != nil {
		return err
	}

	if len(p) < aes.BlockSize {
		return errors.New("Ciphertext too short")
	}

	iv := p[:aes.BlockSize]
	text := p[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)

	hash := sha512.New()
	passwordSha := hash.Sum([]byte(password))

	return bcrypt.CompareHashAndPassword(text, passwordSha)
}
