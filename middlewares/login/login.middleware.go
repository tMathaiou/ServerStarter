package loginMiddleware

import (
	"ServerStarter/models/Users"
	"github.com/kataras/iris"
 	//"fmt"	
)



func LocalStrategy(c *iris.Context) {
	user := &userModel.Users{}
	err := c.ReadJSON(user)

	if err != nil { 
		c.JSON(500, map[string]string{"error": "Something Went Wrong"}) 
		return 
	}

 	if user.Email == "" || user.Password == "" { 
 		c.JSON(200, "Every Field Is Required" )
 		return
 	}

	exist := userModel.FindByEmail(user.Email)

	if exist.Email == "" { 
		c.JSON(422, map[string]string{"error": "The Password And Email Dont Match"} ) 
		return
	}

	err = userModel.ComparePass(user.Password, exist.Password)

	if err != nil { 
		c.JSON(400, map[string]string{"error": "The Password And Email Dont Match"} ) 
		return
	}

	c.Set("id",exist.Id)

	c.Next()
}
