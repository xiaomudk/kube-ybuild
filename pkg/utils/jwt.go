// 主要是配合jwt来生成用户登录token

package utils

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

// Sign signs the payload with the specified secret.
// The token content.
// iss: （Issuer）签发者
// iat: （Issued At）签发时间，用Unix时间戳表示
// exp: （Expiration Time）过期时间，用Unix时间戳表示
// aud: （Audience）接收该JWT的一方
// sub: （Subject）该JWT的主题
// nbf: （Not Before）不要早于这个时间
// jti: （JWT ID）用于标识JWT的唯一ID
func Sign(payload map[string]interface{}, secret string, timeout int64) (tokenString string, err error) {
	now := time.Now().Unix()
	claims := make(jwt.MapClaims)
	claims["nbf"] = now
	claims["iat"] = now
	claims["exp"] = now + timeout

	for k, v := range payload {
		claims[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}

func Secret() jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	}
}

type User struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func ParseToken(tokens string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokens, &User{}, Secret())
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("that's not even a token")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("token not active yet")
			} else {
				return nil, errors.New("couldn't handle this token")
			}
		}
	}
	if user, ok := token.Claims.(*User); ok && token.Valid {
		return user, nil
	}
	return nil, errors.New("couldn't handle this token")
}
