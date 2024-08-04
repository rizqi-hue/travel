package utils

import (
	"crypto/md5"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 5)

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func HashPasswordMD5(password string) string {
	hash := md5.New()
	_, _ = hash.Write([]byte(password))

	md5 := hash.Sum(nil)

	str := fmt.Sprintf("%x", md5)

	sub1 := str[16:len(str)]
	sub2 := str[0:16]

	return sub1 + sub2
}

func CheckHashPasswordMD5(password string, has string) bool {

	log.Println(HashPasswordMD5(password))
	log.Println(has)

	return HashPasswordMD5(password) == has
}
