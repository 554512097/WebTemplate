package utils

import (
	"errors"
	"log"
	"math/rand"
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
		Issuer:    "Auth_Server",                                   // 签发者
		Subject:   "Tom",                                           // 签发对象
		Audience:  jwt.ClaimStrings{"Android_APP", "IOS_APP"},      //签发受众
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
		NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
		IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
		ID:        randStr(10),                                     // wt ID, 类似于盐值
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tc)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["userId"] = userId
	claims["exp"] = time.Now().Add(time.Minute + 30).Unix()
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println(err)
		log.Printf("Something went wrong : %s", err.Error())
		return "", err

	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return signKey, nil
	})
	if err != nil {
		log.Println(err)
		log.Printf("Something went wrong : %s", err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func randStr(str_len int) string {
	rand_bytes := make([]rune, str_len)
	for i := range rand_bytes {
		rand_bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(rand_bytes)
}
