package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/internal-tx/config"
	"github.com/sirupsen/logrus"

	txAPI "github.com/internal-tx/api/tx"
)

func StartServer() {
	engine := fiber.New()
	engine.Use(cors.New())
	engine.Use(logger.New())
	engine.Use(recover.New())

	registerAPIs(engine)

	logrus.Fatal(engine.Listen(config.GetConfig().RestServer.Port))
}

func registerAPIs(app *fiber.App) {
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	txAPIs(app)
}

func txAPIs(app *fiber.App) {
	group := app.Group("/v1/txs")

	group.Get("/internal", txAPI.GetInternalTx)
}
