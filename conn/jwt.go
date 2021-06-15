package conn

import (
	"Asm/moled"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtket = []byte("a_secret_crect")

type Claims struct {
	Userid uint
	jwt.StandardClaims
}

func tokenst(user moled.User) (string, error) {
	timeouttoken := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Userid: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeouttoken.Unix(),
			Issuer:    "oe",
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenstring, err := token.SignedString(jwtket)
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}
func Tokenjs(tokenstring string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtket, nil
	})
	return token, claims, err
}
