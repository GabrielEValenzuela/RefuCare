package main

import (
	"github.com/GabrielEValenzuela/RefuCare/internal/adapters/messaging"
	"github.com/GabrielEValenzuela/RefuCare/internal/adapters/repository"
	"github.com/GabrielEValenzuela/RefuCare/internal/core/service"
	grpcport "github.com/GabrielEValenzuela/RefuCare/internal/ports/grpc"
	"github.com/GabrielEValenzuela/RefuCare/internal/ports/rest"
	"github.com/GabrielEValenzuela/RefuCare/pkg/config"
	"github.com/GabrielEValenzuela/RefuCare/pkg/logger"
	"github.com/GabrielEValenzuela/RefuCare/pkg/metrics"
	"github.com/GabrielEValenzuela/RefuCare/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration and initialize logger
	config.LoadConfig()
	logger.Init(config.Cfg.Env)
	defer logger.Sync()

	logger.Log.Infow("starting service",
		"port", config.Cfg.Port,
		"env", config.Cfg.Env,
	)

	// Fiber app setup
	app := fiber.New()

	// Metrics
	prometheusMetrics := metrics.NewPrometheus("vital_service")
	prometheusMetrics.Use(app)

	// Middlewares
	app.Use(middleware.Recovery())
	app.Use(middleware.RequestLogger())
	app.Use(middleware.CORS(config.Cfg.AllowedOrigin))
	app.Use(middleware.RequestID())
	//app.Use(middleware.JWTAuth()) // optional: apply only to protected routes

	// Routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/health", func(ctx *fiber.Ctx) error {
		logger.Log.Debugw("health check hit", "path", ctx.Path())
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "pass",
		})
	})

	redisRepo := repository.NewRedisVitalsRepository() // implement this next
	grpcClient, err := grpcport.NewPatientRecordsClient("localhost:50051")
	if err != nil {
		logger.Log.Fatalw("failed to connect to patient gRPC", "error", err)
	}

	publisher, err := messaging.NewPublisher(config.Cfg.RabbitMQURL, "vitals_queue")
	if err != nil {
		logger.Log.Fatalw("failed to init RabbitMQ publisher", "error", err)
	}
	defer publisher.Close()

	vitalsSvc := service.NewVitalsService(redisRepo, grpcClient, publisher)
	rest.RegisterRoutes(app, vitalsSvc)

	// Start server
	err = app.Listen(":" + config.Cfg.Port)
	if err != nil {
		logger.Log.Fatalw("failed to start server", "error", err)
	}
}
