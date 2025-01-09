package util_jwt

import "github.com/golang-jwt/jwt/v5"

type UtilsJwt struct{}

func UseJwt() UtilsJwt {
	return UtilsJwt{}
}

func (u UtilsJwt) SignToken(data map[string]interface{}, secret string) (*string, error) {
	claims := jwt.MapClaims(data)
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &token, err
}
