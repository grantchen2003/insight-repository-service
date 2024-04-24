package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateRepositoryId(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]bool{
		"repository_id_is_valid": true,
	})
}
