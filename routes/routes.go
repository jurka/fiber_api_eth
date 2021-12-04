package routes

import (
	"fiber_api_eth/handlers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	// "gorm.io/gorm/logger"
	_ "fiber_api_eth/docs"
)

// New create an instance API routes
func New(app *handlers.App) *fiber.App {
	router := fiber.New()
	router.Use(recover.New())
	router.Use(requestid.New())
	router.Use(cors.New())
	router.Use(logger.New())
	// app.Use(logger.New(logger.Config{
	// 	Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
	// 	TimeFormat: "02-Jan-2006",
	// 	TimeZone:   "Etc/UTC",
	// }))

	router.Get("/docs/*", swagger.Handler)

	api := router.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "API v1",
		})
		return c.Next()
	})

	v1.Get("/groups", app.GetAllGroups)
	v1.Get("/groups/:groupId", app.GetGroupByID)
	v1.Get("/indexes/:indexId", app.GetIndexByID)
	v1.Get("/blocks/:block", app.GetBlockInfo)

	return router
}
