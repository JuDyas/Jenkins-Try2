package handler

import (
	"fmt"
	"github.com/JuDyas/Jenkins-Try2/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WordCountRequest struct {
	Text string `json:"text"`
}

func CountWords(c echo.Context) error {
	var req WordCountRequest
	fmt.Println(req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid input"})
	}

	svc := service.NewWordCountService()
	wordCount := svc.CountWords(req.Text)
	//123

	return c.JSON(http.StatusOK, map[string]int{"word_count": wordCount})
}
