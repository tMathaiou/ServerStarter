package userController

import(
	"startup/models/Users"
 	"github.com/kataras/iris"
	//"fmt"
	"strconv"
)


func Post(c *iris.Context) {
	user := &userModel.Users{}
	err := c.ReadJSON(user)

	if err != nil { 
		c.JSON(500, map[string]string{"error": "Something Went Wrong" })  
		return
	}

 	if user.Email == "" || user.Password == "" { 
 		c.JSON(422, map[string]string{"error": "Every Field Is Required"} )
 		return
 	}

	exist := userModel.FindByEmail(user.Email)

	if exist.Email != "" { 
		c.JSON(422, map[string]string{"error": "User Already Exist"} ) 
		return
	}

	userModel.Save(user)

	if err != nil { 
		c.JSON(500, err ) 
		return
	}



	c.JSON(200, map[string]string{"data": "User Saved"})
	return
	
}

func GetOne(c *iris.Context) {
	id := c.Param("id")
	//c.SetHeader("Access-Control-Allow-Origin", []string{"*"})
	
	user := userModel.FindById(id)

	c.JSON(200, map[string]userModel.Users{"data": user})
	return		
}

func Get(c *iris.Context) {	
	var users userModel.UsersArray
	users = userModel.FindAll()	

	c.JSON(200, map[string]userModel.UsersArray{"data": users})
	return		
}

func Put(c *iris.Context) {
	id := c.Param("id")
	user := &userModel.Users{}
	err := c.ReadJSON(user)

	if err != nil { 
		c.JSON(500, map[string]string{"error": "Something Went Wrong" })  
		return
	}

 	if user.Email == "" { 
 		c.JSON(422, map[string]string{"error": "Email Is Required"} )
 		return
 	}

	exist := userModel.FindById(id)
	
	if exist.Email == "" { 
		c.JSON(422, map[string]string{"error": "User Doesnt Exist"} ) 
		return
	}

	exist = userModel.FindByEmail(user.Email)

	uint_id, _ := strconv.ParseUint(id, 10, 32)
	if exist.Email != "" && exist.Id != uint(uint_id) { 
		c.JSON(422, map[string]string{"error": "The provided email already exist"} ) 
		return
	}
	
	userModel.Update(id, user)

	c.JSON(200, map[string]string{"message": "User Updated"})
	return		
}

func Delete(c *iris.Context) {
	id := c.Param("id")
	
	user := userModel.FindById(id)

	if user.Email == "" {
		c.JSON(422, map[string]string{"error": "User Doesnt Exist"} ) 
		return
	}

	userModel.Delete(id)

	c.JSON(200, map[string]string{"message": "User Deleted"})
	return		
}