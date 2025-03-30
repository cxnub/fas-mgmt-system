package http

import (
	"errors"
	"github.com/cxnub/fas-mgmt-system/internal/core/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// InternalServerError sends a 500 Internal Server Error response with a generic error message in JSON format.
func InternalServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, newErrorResponse("Internal Server Error"))
}

// validationError sends an error response for some specific request validation error
func validationError(ctx *gin.Context, err error, obj interface{}) {
	errorMap := make(map[string]string)
	var ve validator.ValidationErrors

	if errors.As(err, &ve) {
		for _, fe := range ve {
			jsonKey := util.GetJSONTag(obj, fe.StructField()) // Get JSON key
			errorMap[jsonKey] = msgForTag(fe.Tag())           // Use JSON key
		}
		ctx.JSON(http.StatusBadRequest, newValidationErrorResponse(errorMap))
		return
	}
}

func handleError(ctx *gin.Context, err error) {
	if errInfo, exists := errorMap[err]; exists {
		ctx.JSON(errInfo.StatusCode, newErrorResponse(errInfo.Message))
	} else {
		InternalServerError(ctx)
	}
}

// handleSuccess sends a JSON response with the provided HTTP status code, message, and data.
// If the message is empty, a default "Success" message is used.
func handleSuccess(ctx *gin.Context, statusCode int, message string, data any) {
	if message == "" {
		message = "Success"
	}

	ctx.JSON(statusCode, newResponse(true, message, data))
}
