package main 



import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/cors"
	"ServerStarter/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"ServerStarter/database"
	//"startup/models/Users"
)


func main() {	
	var err error	
	db.Db,err = gorm.Open("mysql", "user:pass@/gostarter?charset=utf8&parseTime=True&loc=Local")	

	if err != nil{
   		panic("failed to connect database")   		
  	}	

	iris.Use(cors.New(cors.Options{AllowedOrigins: []string{"*"},}))
	//userModel.Migrate()
	router.Routes()
	
	println("Server is running at :3000")
	iris.Listen(":3000")
	
}

