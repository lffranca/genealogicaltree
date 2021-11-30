package ggin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type personService service

func (pkg *personService) listGET(c *gin.Context) {
	c.Status(http.StatusNoContent)
	c.Done()
}
