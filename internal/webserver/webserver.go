package webserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/n0c1337/linkshortener/internal/auth"
	"github.com/n0c1337/linkshortener/internal/config"
	"gorm.io/gorm"
)

type WebServer struct {
	app    *fiber.App
	db     *gorm.DB
	auth   *auth.Authorization
	config *config.Config
}

func NewWebServer(db *gorm.DB, auth *auth.Authorization, cfg *config.Config) (ws *WebServer) {
	ws = new(WebServer)

	ws.app = fiber.New()
	ws.db = db
	ws.auth = auth
	ws.config = cfg

	ws.setupWebServer()
	ws.registerHandlers()

	return
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

func (ws *WebServer) registerHandlers() {
	api := ws.app.Group("/api")
	userRouterGroup := api.Group("/users")
	linkRouterGroup := api.Group("/links")

	// Normal routes
	ws.app.Get("/redirect/:redirect", ws.redirect)

	// User routes
	userRouterGroup.Post("/", ws.CreateUser)
	userRouterGroup.Get("/:id", ws.GetUser)

	// Link routes
	linkRouterGroup.Post("/", ws.CreateLink)
	linkRouterGroup.Get("/:discriminator", ws.GetLinkByDiscriminator)
	linkRouterGroup.Get("/byId/:id", ws.GetLinkById)
	linkRouterGroup.Delete("/:id", ws.DeleteLink)
}

func (ws *WebServer) ListenAndServe() error {
	return ws.app.Listen(ws.config.Address)
}
