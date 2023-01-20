package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rnwxyz/wishlist-sewa/config"
	"github.com/rnwxyz/wishlist-sewa/constans"
	authCont "github.com/rnwxyz/wishlist-sewa/internal/controllers/auth"
	productCont "github.com/rnwxyz/wishlist-sewa/internal/controllers/products"
	userCont "github.com/rnwxyz/wishlist-sewa/internal/controllers/users"
	productRepo "github.com/rnwxyz/wishlist-sewa/internal/repositories/products"
	userRepo "github.com/rnwxyz/wishlist-sewa/internal/repositories/users"
	authServ "github.com/rnwxyz/wishlist-sewa/internal/services/auth"
	productServ "github.com/rnwxyz/wishlist-sewa/internal/services/products"
	userServ "github.com/rnwxyz/wishlist-sewa/internal/services/users"
	customMiddleware "github.com/rnwxyz/wishlist-sewa/middlewares"
	"github.com/rnwxyz/wishlist-sewa/utils/jwt"
	"github.com/rnwxyz/wishlist-sewa/utils/password"
	customValidator "github.com/rnwxyz/wishlist-sewa/utils/validator"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB) {
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(
		middleware.CORSConfig{
			AllowOrigins: []string{
				"*",
			},
		},
	))
	customValidator.NewCustomValidator(e)
	md := customMiddleware.NewCustomMiddleware(config.Env.JWT_SECRET_ACCESS)
	allAccess := ""
	// owner := constans.RoleOwner
	admin := constans.RoleAdmin
	// user := constans.RoleUser

	api := e.Group("/api")

	// version
	v1 := api.Group("/v1")

	// init repository
	userRepository := userRepo.NewUserRepository(db)
	productRepository := productRepo.NewProductRepository(db)

	// init service
	jwtService := jwt.NewJWTService(config.Env.JWT_SECRET_ACCESS)
	passwordService := password.NewPasswordService()
	authService := authServ.NewAuthService(userRepository, passwordService, jwtService)
	userService := userServ.NewUserService(userRepository)
	productService := productServ.NewProductService(productRepository)

	// init controller
	authController := authCont.NewAuthController(authService)
	userController := userCont.NewUserController(userService, jwtService)
	productController := productCont.NewProductController(productService)

	// routes
	// auth routes
	auth := v1.Group("/auth")
	auth.POST("/login", authController.Login).Name = "login"
	auth.POST("/user-register", authController.UserRegister).Name = "user-register"

	// user routes
	user := v1.Group("/users")
	user.GET("/profile", userController.GetProfile, md.JWTMiddleware(allAccess)).Name = "profile"

	// product routes
	product := v1.Group("/products")
	product.GET("", productController.GetProducts, md.JWTMiddleware(allAccess)).Name = "products"
	product.GET("/:id", productController.GetProductByID, md.JWTMiddleware(allAccess)).Name = "product"
	product.POST("", productController.CreateProduct, md.JWTMiddleware(admin)).Name = "create-product"
	product.PUT("/:id", productController.UpdateProduct, md.JWTMiddleware(admin)).Name = "update-product"
	product.DELETE("/:id", productController.DeleteProduct, md.JWTMiddleware(admin)).Name = "delete-product"

	// fist time create default account
	authService.CreateDefaultAccount()
}
