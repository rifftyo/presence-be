package utils

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hashed string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
    return err == nil
}