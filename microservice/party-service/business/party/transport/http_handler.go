package transport

import (
	"log"
	"net/http"
	"party-service/business/party"
	"party-service/business/party/model"
	"strconv"

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

	// api := e.Group("/api-party")
	e.Use(middleware.Logger())
	public := e.Group("/public/api-party")
	public.GET("/party", handler.GetParty)
	// Added authentication
	private := e.Group("/private/api-party")

	private.POST("/party", handler.CreateParty)
	private.GET("/party/:id", handler.GetPartyByID)
	private.POST("/joinParty", handler.JoinParty)
}

func (h *PartyHTTPHandler) CreateParty(c echo.Context) error {
	var p model.Party
	if err := c.Bind(&p); err != nil {
		log.Println("CreateParty", err)
		return c.JSON(http.StatusBadRequest, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	_, err := h.partySerivce.CreateParty(p)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Return{
		Data:     nil,
		ErrorMsg: "",
	})
}

func (h *PartyHTTPHandler) JoinParty(c echo.Context) error {
	var m model.Member
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusBadRequest, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	err := h.partySerivce.JoinParty(m)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Return{
		Data:     nil,
		ErrorMsg: "",
	})
}

func (h *PartyHTTPHandler) GetParty(c echo.Context) error {
	result, err := h.partySerivce.GetParty()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Return{
		Data:     result,
		ErrorMsg: "",
	})
}

func (h *PartyHTTPHandler) GetPartyByID(c echo.Context) error {
	id := c.Param("id")
	i, _ := strconv.Atoi(id)
	result, err := h.partySerivce.GetPartyByID(i)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Return{
		Data:     result,
		ErrorMsg: "",
	})
}
