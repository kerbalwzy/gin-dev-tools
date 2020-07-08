package kerbalwzygo

import (
	"github.com/dgrijalva/jwt-go"

	"errors"
	"time"
)

// Error notes
var (
	TokenExpired     = errors.New(" Token is expired")
	TokenNotValidYet = errors.New(" Token not active yet")
	TokenMalformed   = errors.New(" That's not even a token")
	TokenInvalid     = errors.New(" Couldn't handle this token:")
)

// 自定义载体, CustomData用于保存自定义的数据; jwt.StandardClaims用于存储载体附属信息, 特别是过期时间
type CustomJWTClaims struct {
	CustomData interface{} `json:"custom_data"`
	jwt.StandardClaims
}

// 生成JWT TOKEN: claims 载体数据, salt加密盐值
func CreateJWTToken(claims CustomJWTClaims, salt []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(salt)
}

// 解析JWT TOKEN: tokenStr TOKEN字符串, salt加密盐值
func ParseJWTToken(tokenStr string, salt []byte) (*CustomJWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return salt, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			switch {
			case ve.Errors&jwt.ValidationErrorMalformed != 0:
				return nil, TokenMalformed
			case ve.Errors&jwt.ValidationErrorExpired != 0:
				return nil, TokenExpired
			case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
				return nil, TokenNotValidYet
			default:
				return nil, TokenInvalid
			}
		}
	} else if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 刷新JWT TOKEN tokenStr TOKEN字符串, salt加密盐值, survivalTime存活时间
func RefreshJWTToken(tokenStr string, salt []byte, survivalTime time.Duration) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &CustomJWTClaims{},
		func(token *jwt.Token) (interface{}, error) { return salt, nil })

	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomJWTClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(survivalTime).Unix()
		return CreateJWTToken(*claims, salt)
	}
	return "", TokenInvalid
}
