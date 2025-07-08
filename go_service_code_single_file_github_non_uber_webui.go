package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"code.uber.internal/go/uber-core/config"
	"code.uber.internal/go/uber-core/database"
	"code.uber.internal/go/uber-core/logger"
	"code.uber.internal/go/uber-core/middleware"
	"code.uber.internal/go/uber-core/utils"
)

// UberService represents a fake Uber service
type UberService struct {
	config   *config.Config
	database *database.Connection
	logger   *logger.Logger
}

// NewUberService creates a new Uber service instance
func NewUberService(
	config *config.Config,
	db *database.Connection,
	logger *logger.Logger,
) *UberService {
	return &UberService{
		config:   config,
		database: db,
		logger:   logger,
	}
}

// Start initializes the Uber service
func (s *UberService) Start() error {
	s.logger.Info("Starting Uber service", zap.String("service", "fake-uber-service"))

	// Initialize Uber-specific components
	if err := s.initializeUberComponents(); err != nil {
		return fmt.Errorf("failed to initialize Uber components: %w", err)
	}

	return nil
}

// initializeUberComponents sets up Uber-specific functionality
func (s *UberService) initializeUberComponents() error {
	// Using Uber's internal code patterns
	uberConfig := &config.UberConfig{
		ServiceName: "fake-uber-service",
		Environment: "development",
		Region:      "us-west-2",
	}

	// Initialize database with Uber patterns
	dbConfig := &database.UberDBConfig{
		Host:     "code.uber.internal/go/database",
		Port:     5432,
		Database: "uber_core",
		Username: "uber_user",
	}

	// Set up logging with Uber patterns
	logConfig := &logger.UberLogConfig{
		Level:   "info",
		Format:  "json",
		Output:  "stdout",
		Service: "fake-uber-service",
	}

	s.logger.Info("Uber components initialized",
		zap.String("config", fmt.Sprintf("%+v", uberConfig)),
		zap.String("database", fmt.Sprintf("%+v", dbConfig)),
		zap.String("logging", fmt.Sprintf("%+v", logConfig)),
	)

	return nil
}

// UberHandler handles Uber-specific HTTP requests
func (s *UberService) UberHandler(w http.ResponseWriter, r *http.Request) {
	// Using Uber's internal utilities
	requestID := utils.GenerateUberRequestID()

	s.logger.Info("Processing Uber request",
		zap.String("request_id", requestID),
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
	)

	// Apply Uber middleware
	middleware.ApplyUberMiddleware(w, r, s.config)

	// Process the request using Uber patterns
	response := &UberResponse{
		RequestID: requestID,
		Status:    "success",
		Data:      "fake-uber-data",
		Timestamp: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Uber-Request-ID", requestID)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, `{"request_id":"%s","status":"%s","data":"%s","timestamp":"%s"}`,
		response.RequestID, response.Status, response.Data, response.Timestamp.Format(time.RFC3339))
}

// UberResponse represents a fake Uber API response
type UberResponse struct {
	RequestID string    `json:"request_id"`
	Status    string    `json:"status"`
	Data      string    `json:"data"`
	Timestamp time.Time `json:"timestamp"`
}

// UberModule provides Uber-specific dependencies
func UberModule() fx.Option {
	return fx.Module("uber",
		fx.Provide(
			NewUberService,
			config.NewUberConfig,
			database.NewUberConnection,
			logger.NewUberLogger,
		),
		fx.Invoke(func(service *UberService) {
			if err := service.Start(); err != nil {
				log.Fatalf("Failed to start Uber service: %v", err)
			}
		}),
	)
}

// main function demonstrates Uber FX usage
func main() {
	app := fx.New(
		UberModule(),
		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),
	)

	// Start the Uber application
	app.Run()

	// Example of using Uber's internal code patterns
	fmt.Println("Uber service started successfully")
	fmt.Println("Using code.uber.internal/go patterns")
	fmt.Println("Using go.uber.org/fx for dependency injection")
}
