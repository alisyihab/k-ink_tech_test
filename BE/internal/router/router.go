package router

import (
	"os"
	"time"

	"registration-smart-reward/internal/handler"
	"registration-smart-reward/internal/middleware"
	"registration-smart-reward/internal/repo"
	"registration-smart-reward/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Build and return Gin engine
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 🌐 CORS: izinkan request dari FE (localhost:3000) ‑‑ atur sesuai domain prod
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// User / Auth
	userRepo := repo.NewUserRepository()
	userSvc := service.NewUserService(userRepo, os.Getenv("JWT_SECRET"))
	authHdl := handler.NewAuthHandler(userSvc)

	// Paket
	paketRepo := repo.NewPaketRepository()
	paketSvc := service.NewPaketService(paketRepo)
	paketHdl := handler.NewPaketHandler(paketSvc)

	// Member
	memberRepo := repo.NewMemberRepository()
	aktivasiRepo := repo.NewPaketAktivasiRepository()
	memberSvc := service.NewMemberService(memberRepo, paketRepo, aktivasiRepo)
	memberHdl := handler.NewMemberHandler(memberSvc)

	/* -------------------------------------------------------------------------- */
	/*                           PUBLIC (no auth required)                        */
	/* -------------------------------------------------------------------------- */
	r.POST("/login", authHdl.Login)       // -> return JWT token
	r.POST("/member", memberHdl.Register) // form registrasi member  (public)
	r.GET("/paket", paketHdl.ListActive)  // dropdown paket (public)

	/* -------------------------------------------------------------------------- */
	/*                  PROTECTED ROUTE ― Require Bearer JWT                      */
	/* -------------------------------------------------------------------------- */
	api := r.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		// CRUD master Paket (tambah paket)
		api.POST("/paket", paketHdl.Create)

		// List member paginated
		api.GET("/members", memberHdl.List)

		// stats
		dashboardHdl := handler.NewDashboardHandler(memberRepo, paketRepo)
		api.GET("/dashboard", dashboardHdl.Overview)

	}

	return r
}
