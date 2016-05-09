package rolesMiddleware

import (	
	"github.com/kataras/iris"
 	//"fmt"	 
    	"startup/models/Users"
	"encoding/json"
	//"strings"
)




func IsAdmin(c *iris.Context) {	
	data, _ := json.Marshal(c.Get("user"))
	s := string(data)
	bytes := []byte(s)
    	var user userModel.Users
    	json.Unmarshal(bytes, &user)

	if user.Role != "" && user.Role == "admin" {		        
	   c.Next() 
	   return	    	        		
	}
	
	c.JSON(401, map[string]string{"error": "Unauthorized"})
	return	

}
