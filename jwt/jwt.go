package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(m map[string]interface{}, signKey []byte, accessExpireTime, refreshExpireTime time.Duration) (string, string, error) {
	// Access token yaratish
	accessToken := jwt.New(jwt.SigningMethodHS256)
	aClaims := accessToken.Claims.(jwt.MapClaims)

	for key, value := range m {
		aClaims[key] = value
	}
	aClaims["exp"] = time.Now().Add(accessExpireTime).Unix()

	// Refresh token yaratish
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rClaims := refreshToken.Claims.(jwt.MapClaims)

	for key, value := range m {
		rClaims[key] = value
	}
	rClaims["exp"] = time.Now().Add(refreshExpireTime).Unix()

	// Tokenlarni imzolash
	accessTokenStr, err := accessToken.SignedString(signKey)
	if err != nil {
		return "", "", err
	}

	refreshTokenStr, err := refreshToken.SignedString(signKey)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func ExtractClaims(tokenString string, signKey []byte) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Agar token yaroqli bo'lsa, claimslarni olish
	if ok := token.Valid; ok {
		return token.Claims.(jwt.MapClaims), nil
	}

	return nil, fmt.Errorf("Invalid token")
}
