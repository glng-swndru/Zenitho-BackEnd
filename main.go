package main

import (
	"campaignku/handler"
	"campaignku/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Memasukkan detail database ke dalam env agar lebih aman
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Mengonfigurasi koneksi database menggunakan GORM
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Membuat instance repository dan service untuk entitas pengguna
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// Membuat instance handler untuk pengguna
	userHandler := handler.NewUserHandler(userService)

	// Membuat router menggunakan framework Gin
	router := gin.Default()
	api := router.Group("/api/v1")

	// Menetapkan endpoint untuk mendaftarkan pengguna
	api.POST("/users", userHandler.RegisterUser)

	// Menjalankan server pada port default (8080)
	router.Run()

	// Catatan:
	// - input dari user
	// - handler, mapping input dari user -> struct input
	// - service: melakukan mapping dari struct input ke struct user
	// - repository
	// - db
}
