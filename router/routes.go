package router	

import(
	"github.com/kataras/iris"
	"startup/controllers/Users"
	"startup/controllers/Auth"
	"startup/middlewares/login" 
	"startup/middlewares/Auth"
	"startup/middlewares/Roles"
)

func Routes(){   

	iris.Post("/login", loginMiddleware.LocalStrategy, authController.Auth)	

	iris.Get("/users", authMiddleware.JwtStrategy, userController.Get)

	iris.Get("/users/:id", authMiddleware.JwtStrategy, userController.GetOne)

	iris.Post("/users", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Post)	

	iris.Put("/users/:id", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Put)

	iris.Delete("/users/:id", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Delete)

	
}