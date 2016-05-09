package authController

import (
	//"startup/models/Users" 	
 	"github.com/kataras/iris" 	
	"github.com/dgrijalva/jwt-go"
	"time"
	//"encoding/json"
)

var mySigningKey = []byte("dfgdwwer439738hsduihf3897r9wfh")

func Auth(c *iris.Context) {		
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["id"] = c.Get("id")
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		c.JSON(500, map[string]string{"error": "Something Went Wrong"})
		return		
	}

	mapToken := map[string]string{"token": tokenString}  
   
	c.JSON(200, mapToken)
	return
	
}