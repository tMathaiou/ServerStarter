package router	

import(
	"github.com/kataras/iris"
	"ServerStarter/controllers/Users"
	"ServerStarter/controllers/Auth"
	"ServerStarter/middlewares/login" 
	"ServerStarter/middlewares/Auth"
	"ServerStarter/middlewares/Roles"
)

func Routes(){   

	iris.Post("/login", loginMiddleware.LocalStrategy, authController.Auth)	

	iris.Get("/users", authMiddleware.JwtStrategy, userController.Get)

	iris.Get("/temp", authMiddleware.JwtStrategy, func(c *iris.Context) {
        user := c.Get("user")
        c.JSON(200, user)
    })

	iris.Get("/users/:id", authMiddleware.JwtStrategy, userController.GetOne)

	iris.Post("/users", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Post)	

	iris.Put("/users/:id", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Put)

	iris.Delete("/users/:id", authMiddleware.JwtStrategy, rolesMiddleware.IsAdmin, userController.Delete)

	
}