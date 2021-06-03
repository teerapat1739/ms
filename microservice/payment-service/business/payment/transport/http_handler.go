package transport

import (
	"net/http"
	"payment-service/business/payment"
	"payment-service/business/payment/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// paymentHTTPHandler - payment handler for http request

type paymentHTTPHandler struct {
	paymentSerivce payment.Service
}

// NewpaymentHTTPHandler - init payment
func NewpaymentHTTPHandler(e *echo.Echo, paymentSerivce payment.Service) {
	handler := paymentHTTPHandler{paymentSerivce}

	e.Use(middleware.Logger())
	// Added authentication
	private := e.Group("/private/api-payment")

	private.POST("/payment", handler.Payment)
}

func (h *paymentHTTPHandler) Payment(c echo.Context) error {
	var d model.Request
	err := c.Bind(&d)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}

	err = h.paymentSerivce.Payment(d.Email, d.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Return{
			Data:     nil,
			ErrorMsg: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, model.Return{
		Data:     "result",
		ErrorMsg: "",
	})
}
