package kerbalwzygo

import (
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestJWTToken(t *testing.T) {
	claims := &CustomJWTClaims{
		CustomData: "hello world",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 5,
		},
	}
	res, err := CreateJWTToken(*claims, []byte("Testing"))
	if nil != err {
		t.Fatal(err)
	} else {
		t.Log(res)
	}

	claims, err = ParseJWTToken(res, []byte("Testing"))
	if nil != err {
		t.Fatal(err)
	} else {
		t.Log(*claims)
	}

	res, err = RefreshJWTToken(res, []byte("Testing"), 100*time.Second)
	if nil != err {
		t.Fatal(err)
	} else {
		t.Log(res)
		claims, _ = ParseJWTToken(res, []byte("Testing"))
		t.Log(*claims)
	}
}
