// Package main adalah titik masuk utama dari aplikasi Campaignku.
package main

import (
	"campaignku/auth"
	"campaignku/campaign"
	"campaignku/handler"
	"campaignku/helper"
	"campaignku/user"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Load konfigurasi database dari variabel lingkungan (.env)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inisialisasi koneksi database menggunakan GORM
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Buat repository dan service untuk pengguna
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	campaigns, err := campaignRepository.FindAll()

	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println("debug")
	fmt.Println(len(campaigns))

	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// Buat handler untuk pengguna
	userHandler := handler.NewUserHandler(userService, authService)

	// Buat router menggunakan framework Gin
	router := gin.Default()
	api := router.Group("/api/v1")

	// Tetapkan endpoint untuk pendaftaran pengguna, login, pengecekan ketersediaan email, dan unggah avatar
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)

	// Jalankan server pada port default (8080)
	router.Run()
}

// authMiddleware adalah fungsi middleware untuk otentikasi.
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapatkan header Authorization dari permintaan
		authHeader := c.GetHeader("Authorization")

		// Periksa apakah header mengandung token Bearer
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Ekstrak token dari header
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		// Validasi token
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Ekstrak claim dari token
		claim, ok := token.Claims.(jwt.MapClaims)

		// Periksa apakah claim valid
		if !ok || !token.Valid {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Ekstrak ID pengguna dari claim
		userID := int(claim["user_id"].(float64))

		// Dapatkan informasi pengguna dari layanan
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Tetapkan kunci "currentUser" di konteks dengan informasi pengguna
		c.Set("currentUser", user)
	}
}
