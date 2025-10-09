package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"oot.me/todo-list-api/config"
)

var key = []byte(config.JWT_KEY)

func SignJWT(uid int64,name,email string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
  		jwt.MapClaims{ 
			"uid" : uid,
    		"name": name, 
    		"email": email,
			"exp": time.Now().Add(time.Hour).Unix(),
  		}) 
	signedToken,err := token.SignedString(key) 
	if err != nil {
		return "",err
	}
	return signedToken,nil
}

func ValidateJWT(token string) (map[string]string,error) {
	parsedToken,err := jwt.Parse(token,func(t *jwt.Token) (any, error) {
		return key,nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return map[string]string{},err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok {
		claimList := make(map[string]string)
		claimList["uid"] = fmt.Sprintf("%f", claims["uid"])
		claimList["name"] = claims["name"].(string)
		claimList["email"] = claims["email"].(string)
		claimList["exp"] = fmt.Sprintf("%f", claims["exp"])
		return claimList,nil
	} else {
		return map[string]string{},err
	}
}