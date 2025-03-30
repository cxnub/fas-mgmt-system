package http

import (
	"github.com/cxnub/fas-mgmt-system/internal/adapter/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

// NewRouter creates a new HTTP router
func NewRouter(
	config *config.Config,
	applicantHandler ApplicantHandler,
	schemeHandler SchemeHandler,
	applicationHandler ApplicationHandler,
) (*Router, error) {
	// CORS
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()

	// Register custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("marital_status", validateMaritalStatus)
		v.RegisterValidation("relationship_type", validateRelationshipType)
		v.RegisterValidation("sex", validateSex)
		v.RegisterValidation("employment_status", validateEmploymentStatus)
		v.RegisterValidation("date", validateDate)
	}
	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		applicants := api.Group("/applicants")
		{
			// Applicant routes
			applicants.GET("/", applicantHandler.ListApplicants)
			applicants.GET("/:id", applicantHandler.GetApplicant)
			applicants.POST("/", applicantHandler.CreateApplicant)
			applicants.PUT("/:id", applicantHandler.UpdateApplicant)
			applicants.DELETE("/:id", applicantHandler.DeleteApplicant)
		}

		// Scheme routes
		schemes := api.Group("/schemes")
		{
			schemeIdRoutes := schemes.Group("/:scheme_id")
			{
				schemeIdRoutes.GET("/", schemeHandler.GetScheme)
				schemeIdRoutes.PUT("/", schemeHandler.UpdateScheme)
				schemeIdRoutes.DELETE("/", schemeHandler.DeleteScheme)

				schemeIdRoutes.POST("/benefits", schemeHandler.AddSchemeBenefit)

				schemeIdRoutes.POST("/criteria", schemeHandler.AddSchemeCriteria)

			}

			benefitsRoutes := schemes.Group("/benefits")
			{
				benefitsRoutes.PUT("/:id", schemeHandler.UpdateSchemeBenefit)
				benefitsRoutes.DELETE("/:id", schemeHandler.DeleteSchemeBenefit)
			}

			schemeCriteriaRoutes := schemes.Group("/criteria")
			{
				schemeCriteriaRoutes.PUT("/:id", schemeHandler.UpdateSchemeCriteria)
				schemeCriteriaRoutes.DELETE("/:id", schemeHandler.DeleteSchemeCriteria)
			}

			schemes.GET("/", schemeHandler.ListSchemes)
			schemes.GET("/eligible", schemeHandler.ListApplicantAvailableSchemes)
			schemes.POST("/", schemeHandler.CreateScheme)
		}

		// Application routes
		applications := api.Group("/applications")
		{
			applications.GET("/", applicationHandler.ListApplications)
			applications.GET("/:id", applicationHandler.GetApplication)
			applications.POST("/", applicationHandler.CreateApplication)
			applications.PUT("/:id", applicationHandler.UpdateApplication)
			applications.DELETE("/:id", applicationHandler.DeleteApplication)
		}
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
