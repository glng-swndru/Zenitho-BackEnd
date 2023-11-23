// Package main adalah entry point utama dari aplikasi Campaignku.
package main

import (
	"campaignku/handler"
	"campaignku/user"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Konfigurasi database dari variabel lingkungan (.env)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Menginisialisasi koneksi database menggunakan GORM
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Membuat repository dan service untuk pengguna
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// Mencari pengguna berdasarkan email
	userByEmail, err := userRepository.FindByEmail("gilangswandaru27@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Menampilkan hasil pencarian pengguna
	if userByEmail.ID == 0 {
		fmt.Println("User tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}

	// Membuat handler untuk pengguna
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
