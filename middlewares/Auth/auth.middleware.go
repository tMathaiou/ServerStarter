package authMiddleware

import (
	"startup/models/Users"
	"github.com/kataras/iris"
 	"fmt"	
 	"github.com/dgrijalva/jwt-go"
)


var mySigningKey = []byte("dfgdwwer439738hsduihf3897r9wfh")

func JwtStrategy(c *iris.Context) {
	token := c.RequestHeader("Authorization")
	valid,err := Decode(token)
	if err != nil {
		c.JSON(401, map[string]string{"error": "Invalid Token"})
		return
	}
	
	exist := userModel.FindById(fmt.Sprintf("%.2f", valid.Claims["id"]))
	

	if exist.Email == "" {
		c.JSON(401, map[string]string{"error": "Invalid Token"})
		return
	}

	c.Set("user",exist)

	c.Next()
}

  

func Decode(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        	return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }        
        return []byte(mySigningKey), nil
    })

    if err == nil && token.Valid {
		return token,nil
	} else {
		return nil,err
	}
}