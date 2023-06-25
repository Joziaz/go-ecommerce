package middlewares

import (
	shared "ecommerce/shared/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		var errorsList []string
		for _, err := range ctx.Errors {
			switch e := err.Err.(type) {
			case *shared.DomainError:
				errorsList = append(errorsList, e.Error())
			default:
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		if errorsList != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"errors": errorsList,
			})
		}
	}
}
