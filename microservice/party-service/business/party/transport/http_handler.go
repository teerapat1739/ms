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
	api.POST("/party", handler.CreateParty)
	api.POST("/joinParty", handler.JoinParty)
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

func (h *PartyHTTPHandler) JoinParty(c echo.Context) error {
	var m model.Member
	if err := c.Bind(&m); err != nil {
		log.Println("JoinParty", err)
		return c.JSON(http.StatusBadRequest, "err")
	}
	err := h.partySerivce.JoinParty(m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "err")
	}
	return c.JSON(http.StatusOK, "success")
}
