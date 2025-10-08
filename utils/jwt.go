package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"oot.me/todo-list-api/config"
)

var key = []byte(config.JWT_KEY)

func SignJWT(name,email string) (string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
  		jwt.MapClaims{ 
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

func ValidateJWT(token string) ([]string,error) {
	parsedToken,err := jwt.Parse(token,func(t *jwt.Token) (any, error) {
		return key,nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return []string{},err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok {
		var claimList []string
		claimList = append(claimList, claims["name"].(string), claims["email"].(string),claims["exp"].(string))
		return claimList,nil
	} else {
		return []string{},err
	}
}