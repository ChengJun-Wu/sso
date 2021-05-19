package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func PasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func PasswordVerify(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomString(len int) string {
	var (
		source rand.Source
		r *rand.Rand
	)
	source = rand.NewSource(time.Now().UnixNano())
	r = rand.New(source)
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}