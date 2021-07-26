package usecase

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	_helpers "github.com/siends/siends-api/helpers"
	_models "github.com/siends/siends-api/models"
)

const envKey = "SECRET"

type encryptPasswordInput struct {
	ID       uint
	Password string
}

type comparePasswordInput struct {
	ID                uint
	UserPassword      string
	PasswordToCompare string
}

func encryptPassword(input encryptPasswordInput) (string, error) {
	password := []byte(input.Password + fmt.Sprint(input.ID))
	encryptedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(encryptedPassword), nil
}

func comparePassword(input comparePasswordInput) bool {
	passwordToCompare := []byte(input.PasswordToCompare + fmt.Sprint(input.ID))
	err := bcrypt.CompareHashAndPassword([]byte(input.UserPassword), passwordToCompare)

	return err == nil
}

func createToken(claims _models.CustomClaims) (string, error) {
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte(_helpers.GetEnvironmentVariable(envKey))
	return token.SignedString(secret)
}
