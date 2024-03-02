package middleware

import (
	"BelajarAPI/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(hp string) (string, error) {
	var data = jwt.MapClaims{}
	data["hp"] = hp
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(config.JWTSECRET))

	if err != nil {
		return "", err
	}

	return result, nil
}
