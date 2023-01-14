package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/n0c1337/linkshortener/internal/auth"
	"gorm.io/gorm"
)

type WebServer struct {
	app  *fiber.App
	db   *gorm.DB
	auth *auth.Authorization
}

func NewWebServer(db *gorm.DB, auth *auth.Authorization) (ws *WebServer) {
	ws = new(WebServer)

	ws.app = fiber.New()
	ws.db = db
	ws.auth = auth

	ws.registerHandlers()

	return
}

func (ws *WebServer) registerHandlers() {
	api := ws.app.Group("/api")
	userGroup := api.Group("/users")

	// User routes
	userGroup.Post("/", ws.CreateUser)
	userGroup.Get("/:id", ws.GetUser)
}

func (ws *WebServer) setupWebServer() {
	// Implement cors middleware
	ws.app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Implement ratelimiter
	ws.app.Use(limiter.New())

	// Implement logger middleware
	ws.app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
}

func (ws *WebServer) ListenAndServe() error {
	ws.setupWebServer()
	return ws.app.Listen("192.168.2.113:8080")
}
