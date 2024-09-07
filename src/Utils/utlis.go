package Utils

import (
	"book-fiber/Model"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var secretKey = []byte("SECRET:)")

func GeneratorToken(user Model.Auth) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = user.Id
	claims["user_name"] = user.UserName
	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token Error")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &claims, nil
	}

	return nil, errors.New("invalid token")
}

func GeneratorPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func ValidatePassword(userPassword, hashPassword string) error {
	vPass := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword))
	if vPass != nil {
		return errors.New("invalid password")
	}
	return nil
}
