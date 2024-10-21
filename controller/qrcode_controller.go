package controller

import (
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/matthewyuh246/qrcode_go/usecase"
)

type IQRCodeController interface {
	GenerateQRCode(c echo.Context) error
	GetRecentQRCodes(c echo.Context) error
}

type qrcodeController struct {
	qu usecase.IQRCodeUsecase
}

func NewQRCodeController(qu usecase.IQRCodeUsecase) IQRCodeController {
	return &qrcodeController{qu}
}

func (qc *qrcodeController) GenerateQRCode(c echo.Context) error {
	var requestBody struct {
		Text string `json:"text"`
	}

	if err := c.Bind(&requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	qrCode, err := qc.qu.GenerateQRCode(requestBody.Text)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to generate QR code"})
	}

	err = qc.qu.SaveQRCode(qrCode)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to save QR code"})
	}

	return c.Blob(http.StatusOK, "image/png", qrCode.Image)
}

func (qc *qrcodeController) GetRecentQRCodes(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"].(uint)
	limitStr := c.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 4
	}

	qrCodes, err := qc.qu.GetRecentQRCodes(limit, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to fetch recent QR codes"})
	}

	return c.JSON(http.StatusOK, qrCodes)
}
