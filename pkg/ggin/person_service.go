package ggin

import (
	"github.com/gin-gonic/gin"
	"github.com/lffranca/genealogicaltree/pkg/ggin/presenter"
	"net/http"
)

type personService service

func (pkg *personService) itemJSONGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/json" {
		c.Next()
		return
	}

	var personNameURI presenter.PersonNameURI
	if err := c.ShouldBindUri(&personNameURI); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	items, err := pkg.Client.personRepository.ByName(c.Request.Context(), personNameURI.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.PersonToPresenters(items))
}

func (pkg *personService) listJSONGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/json" {
		c.Next()
		return
	}

	items, err := pkg.Client.personRepository.All(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.PersonToPresenters(items))
}

func (pkg *personService) itemXMLGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/xml" {
		c.Next()
		return
	}

	var personNameURI presenter.PersonNameURI
	if err := c.ShouldBindUri(&personNameURI); err != nil {
		c.XML(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	items, err := pkg.Client.personRepository.ByName(c.Request.Context(), personNameURI.Name)
	if err != nil {
		c.XML(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.XML(http.StatusOK, presenter.XMLEnvelope{
		Data: presenter.PersonToPresenters(items),
	})
}

func (pkg *personService) listXMLGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/xml" {
		c.Next()
		return
	}

	items, err := pkg.Client.personRepository.All(c.Request.Context())
	if err != nil {
		c.XML(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.XML(http.StatusOK, presenter.XMLEnvelope{
		Data: presenter.PersonToPresenters(items),
	})
}

func (pkg *personService) shortestPathJSONGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/json" {
		c.Next()
		return
	}

	var shortestPathURI presenter.ShortestPathURI
	if err := c.ShouldBindUri(&shortestPathURI); err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	items, err := pkg.Client.personRepository.ShortestPath(
		c.Request.Context(),
		shortestPathURI.Name1,
		shortestPathURI.Name2,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, presenter.PersonToPresenters(items))
}

func (pkg *personService) shortestPathXMLGET(c *gin.Context) {
	if value := c.Request.Header.Get("accept"); value != "application/xml" {
		c.Next()
		return
	}

	var shortestPathURI presenter.ShortestPathURI
	if err := c.ShouldBindUri(&shortestPathURI); err != nil {
		c.XML(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	items, err := pkg.Client.personRepository.ShortestPath(
		c.Request.Context(),
		shortestPathURI.Name1,
		shortestPathURI.Name2,
	)
	if err != nil {
		c.XML(http.StatusBadRequest, presenter.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	c.XML(http.StatusOK, presenter.XMLEnvelope{
		Data: presenter.PersonToPresenters(items),
	})
}
