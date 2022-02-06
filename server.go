package main

import (
	"log"
	"os"

	conf "new-proj/config"
	cont "new-proj/controller"
	"new-proj/middleware"
	"new-proj/repositories"
	"new-proj/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = conf.SetupDBConnection()
	userRepo repositories.UserRepository = repositories.NewUserRepository(db)
	bookRepo repositories.BookRepository = repositories.NewBookRepository(db)
	
	jwtService service.JWTService = service.NewJWTService()
	userService service.UserService = service.NewUserService(userRepo)
	authService service.AuthService = service.NewAuthService(userRepo)
	bookService service.BookService = service.NewBookService(bookRepo)

	authController cont.AuthController = cont.NewAuthController(authService, jwtService)
	userController cont.UserController = cont.NewUserController(userService, jwtService)
	bookController cont.BookController = cont.NewBookController(bookService, jwtService)
)

func Server () {
	defer conf.CloseDBConnection(db)
	
	router := gin.Default()

	router.Use(middleware.CORS())

	auth := router.Group("/api/auth")

	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
	}

	userRoute := router.Group("/api/user")

	{
		userRoute.PUT("/update", userController.Update)
		userRoute.GET("/profile", userController.Profile)
	}

	bookRoute := router.Group("/api/book")

	{
		bookRoute.PUT("/update", bookController.Update)
		bookRoute.POST("/create", bookController.Insert)
		bookRoute.GET("/get-all", bookController.All)
		bookRoute.DELETE("/remove/:id", bookController.Delete)
		bookRoute.GET("/by-id/:id", bookController.FindById)
	}

	app_port := os.Getenv("PORT")

	if (app_port == "") {
		app_port = "8080"
	}

	log.Fatal(router.Run(":" + app_port))
}