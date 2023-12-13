// Package main adalah titik masuk utama dari aplikasi Campaignku.
package main

import (
	// Impor package-package yang dibutuhkan.
	"campaignku/auth"
	"campaignku/campaign"
	"campaignku/handler"
	"campaignku/helper"
	"campaignku/user"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go" // Untuk urusin JWT.
	"github.com/gin-gonic/gin"    // Gin, framework buat bikin web server.
	"github.com/joho/godotenv"    // Untuk baca file .env.
	"gorm.io/driver/mysql"        // Driver MySQL untuk GORM.
	"gorm.io/gorm"                // GORM, ORM untuk Go.
)

func main() {
	// Baca konfigurasi database dari file .env.
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Sambung ke database MySQL pake GORM.
	dbUser := os.Getenv("DB_USER")         // Username DB.
	dbPassword := os.Getenv("DB_PASSWORD") // Password DB.
	dbHost := os.Getenv("DB_HOST")         // Host DB.
	dbName := os.Getenv("DB_NAME")         // Nama DB.
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	// Buat repository untuk user dan campaign.
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)

	// Buat service untuk user, campaign, dan autentikasi.
	userService := user.NewService(userRepository)
	campaignService := campaign.NewService(campaignRepository)
	authService := auth.NewService()

	// Siapin handler buat handle request ke user dan campaign.
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// Inisialisasi router pake Gin.
	router := gin.Default()
	api := router.Group("/api/v1")

	// Set endpoint dan method yang sesuai.
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/campaigns", campaignHandler.GetCampaigns)

	// Jalankan server di port 8080.
	router.Run()
}

// Fungsi middleware buat otentikasi.
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Dapetin header Authorization dari request.
		authHeader := c.GetHeader("Authorization")

		// Cek header ada token Bearer-nya apa enggak.
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Ambil dan validasi token dari header.
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Cek claim dari token.
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Ambil userID dari claim, cari user di service.
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.ApiResponse("Tidak diizinkan", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		// Simpan informasi user di context request.
		c.Set("currentUser", user)
	}
}
