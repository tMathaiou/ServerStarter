package rolesMiddleware

import (	
	"github.com/kataras/iris"
 	//"fmt"	 
    	"startup/models/Users"
	"encoding/json"
	//"strings"
)




func IsAdmin(c *iris.Context) {	
	user := c.Get("user").(userModel.Users)

	if user.Role != "" && user.Role == "admin" {		        
	   c.Next() 
	   return	    	        		
	}
	
	c.JSON(401, map[string]string{"error": "Unauthorized"})
	return	

}
