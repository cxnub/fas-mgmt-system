package main

import (
	"fmt"
	_ "github.com/cxnub/fas-mgmt-system/docs"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/config"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/handler/http"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres"
	"github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres/repository"
	pg "github.com/cxnub/fas-mgmt-system/internal/adapter/storage/postgres/sqlc"
	_ "github.com/cxnub/fas-mgmt-system/internal/core/domain"
	_ "github.com/cxnub/fas-mgmt-system/internal/core/port"
	"github.com/cxnub/fas-mgmt-system/internal/core/service"
	_ "github.com/swaggo/files"       // Swagger files
	_ "github.com/swaggo/gin-swagger" // Required for Swagger documentation
	"golang.org/x/net/context"
	"log"
)

// @title FAS Management System API
// @version 1.0
// @description This is the API documentation for FAS Management System.
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
func main() {
	cfg := config.New()

	// Init Database
	ctx := context.Background()
	db, err := postgres.New(ctx, cfg)
	q := pg.New(db)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Dependency Injection
	applicantRepo := repository.NewApplicantRepository(db, q)
	applicantService := service.NewApplicantService(applicantRepo)
	applicantHandler := http.NewApplicantHandler(applicantService)

	schemeRepo := repository.NewSchemeRepository(db, q)
	schemeService := service.NewSchemeService(schemeRepo, applicantRepo)
	schemeHandler := http.NewSchemeHandler(schemeService)

	applicationRepo := repository.NewApplicationRepository(db, q)
	applicationService := service.NewApplicationService(applicationRepo, applicantRepo, schemeRepo)
	applicationHandler := http.NewApplicationHandler(applicationService)

	// Init Router
	router, err := http.NewRouter(
		cfg,
		*applicantHandler,
		*schemeHandler,
		*applicationHandler,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", cfg.ApiUrl, cfg.ApiPort)
	log.Print("Starting the HTTP Server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		log.Fatal(err)
	}
}
