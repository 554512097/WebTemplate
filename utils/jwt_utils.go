package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signKey = []byte("kevinxie")

type tokenClaim struct {
	UserId uint
	jwt.RegisteredClaims
}

func GenerateJWT(userId uint) (string, error) {
	tc := tokenClaim{userId, jwt.RegisteredClaims{
		Issuer:    "web_template",                                        // 签发者
		Subject:   "all",                                                 // 签发对象
		Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP", "Web_APP"}, //签发受众
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),         //过期时间
		NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)),       //最早使用时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                        //签发时间
		ID:        time.Now().String(),                                   // wt ID, 类似于盐值
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println(err)
		log.Printf("Something went wrong : %s", err.Error())
		return "", err

	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		log.Println(err)
		log.Printf("Something went wrong : %s", err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
