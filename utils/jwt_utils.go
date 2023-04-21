package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var signKey = []byte("kevinxie")

// func GenerateJWT(userId uint) (string, error) {
// 	rc := jwt.RegisteredClaims{
// 		Issuer:    "web_template",                                  // 签发者
// 		Subject:   strconv.FormatUint(uint64(userId), 10),          // 签发对象
// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),   //过期时间
// 		NotBefore: jwt.NewNumericDate(time.Now().Add(time.Second)), //最早使用时间
// 		IssuedAt:  jwt.NewNumericDate(time.Now()),                  //签发时间
// 		ID:        time.Now().String(),                             // wt ID, 类似于盐值
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, rc)
// 	tokenString, err := token.SignedString(signKey)
// 	if err != nil {
// 		log.Println(err)
// 		log.Printf("Something went wrong : %s", err.Error())
// 		return "", err

// 	}
// 	return tokenString, nil
// }

//	func VerifyJWT(tokenString string) (*jwt.Claims, error) {
//		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//			return signKey, nil
//		})
//		if err != nil {
//			log.Println(err)
//			log.Printf("Something went wrong : %s", err.Error())
//			return nil, err
//		}
//		if claims := token.Claims; token.Valid {
//			return &claims, nil
//		} else {
//			return nil, errors.New("invalid token")
//		}
//	}

func GenerateJWT(userId uint) (string, error) {
	mc := jwt.MapClaims{
		"uid": userId,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Minute)),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	s, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}
	return s, nil
}

func VerifyJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		log.Println(err)
		log.Printf("Something went wrong : %s", err.Error())
		return 0, err
	}
	if mc := token.Claims.(jwt.MapClaims); token.Valid {
		return uint(mc["uid"].(float64)), nil
	} else {
		return 0, errors.New("invalid token")
	}
}
