package utils

import (
    "time"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
    "webdev-intern-assignment/models"
)

var jwtKey = []byte("taikey")

func GenerateToken(user models.User) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &models.Claims{
        Username: user.Username,
        Role:     user.Role,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}