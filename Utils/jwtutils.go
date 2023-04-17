package Utils

import (
	"face/Global"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
)

func Decode_jwt_token(encoded_jwt string, node string) string {
	JWT_SECRET_KEY := Global.JWTKey
	token, err := jwt.Parse(encoded_jwt, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET_KEY), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "107"
		}
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return "106"
			}
			return "108"
		}
		return "108"
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//fmt.Println("验证通过")
		if node == "username" {
			return get_username_from_jwt(claims)
		} else if node == "userid" {

			return get_userid_from_jwt(claims)
		}

	}
	return "108"
}
func get_username_from_jwt(claims jwt.MapClaims) string {
	if data, ok := claims["data"].(map[string]interface{}); ok {
		if username, ok := data["username"].(string); ok {
			return username
		}
	}
	return ""
}
func get_userid_from_jwt(claims jwt.MapClaims) string {

	if data, ok := claims["data"].(map[string]interface{}); ok {

		if userid, ok := data["userid"].(float64); ok {
			//fmt.Println(userid)
			return strconv.Itoa(int(userid))
		}
	}
	return ""
}
