package login

import (
	"backend/key"
	"backend/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func createJwtToken(id int) (string, error) {

	claims := models.JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := rawToken.SignedString([]byte(key.TOKEN_SECRET))

	if err != nil {

		return "", err

	}

	return token, nil
}
