package transport

import (
	"log"
	"net/http"
	"party-service/business/party"
	"party-service/business/party/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// PartyHTTPHandler - party handler for http request

type PartyHTTPHandler struct {
	partySerivce party.Service
}

// NewPartyHTTPHandler - init party
func NewPartyHTTPHandler(e *echo.Echo, partySerivce party.Service) {
	handler := PartyHTTPHandler{partySerivce}

	api := e.Group("/api")
	api.Use(middleware.Logger())

	// Added authentication
	api.GET("/party/:id", handler.FindByID)
	api.POST("/party", handler.CreateParty)
}

func (h *PartyHTTPHandler) CreateParty(c echo.Context) error {
	var p model.Party
	if err := c.Bind(&p); err != nil {
		log.Println("CreateParty", err)
		return c.JSON(http.StatusBadRequest, "err")
	}
	_, err := h.partySerivce.CreateParty(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, "success")
}
func (h *PartyHTTPHandler) FindByID(c echo.Context) error {
	// id := c.Param("id")

	// res, err := h.partySerivce.FindByID(id)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, "err")
	// }
	return c.JSON(http.StatusOK, "res")
}
