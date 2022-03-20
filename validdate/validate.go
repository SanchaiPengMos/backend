package validdate

import (
	"fmt"
	"server/key"
	"server/models/aut"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(token string) (aut.JwtClaims, error) {

	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %s", token.Header["alg"])

		}
		return []byte(key.TOKEN_SECRET), nil
	})

	var result aut.JwtClaims
	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		result.Id = int(claims["id"].(float64))
		return result, nil

	} else {
		return result, err
	}

}
