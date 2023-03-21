package main

import (
	"erteadmin_backend/auth"
	"erteadmin_backend/handler"
	"erteadmin_backend/lingkungan"
	"erteadmin_backend/user"
	"erteadmin_backend/wilayah"
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// Repository
	userRepository := user.NewRepository(db)
	lingkunganRepository := lingkungan.NewRepository(db)
	wilayahRepository := wilayah.NewRepository(db)

	// Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	lingkunganService := lingkungan.NewService(lingkunganRepository)
	wilayahService := wilayah.NewService(wilayahRepository)

	// Handler
	userHandler := handler.NewUserHandler(userService, authService)
	lingkunganHandler := handler.NewLingkunganHandler(lingkunganService)
	wilayahHandler := handler.NewWilayahHandler(wilayahService)

	// gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("erteadmin", cookieStore))

	api := router.Group("/api/v1")

	// Lingkungan
	api.POST("/lingkungan", lingkunganHandler.CreateLingkungan)

	// Wilayah
	api.GET("/provinsi", wilayahHandler.GetAllProvinsi)
	api.POST("/kabkota", wilayahHandler.GetKabKotaByProvinsiID)
	api.POST("/kecamatan", wilayahHandler.GetKecamatanByKabKotaID)
	api.POST("/kelurahan", wilayahHandler.GetKelurahanByKecamatanID)

	// User
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run()
}
