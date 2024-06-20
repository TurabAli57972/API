package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/time", getTime)
	e.POST("/capitalize", postCapitalizeSentence)

	e.Logger.Fatal(e.Start(":8080"))
}

func getTime(c echo.Context) error {
	currentTime := time.Now().Format("03:04 PM")
	return c.JSON(http.StatusOK, map[string]string{
		"current_time": currentTime,
	})
}

func postCapitalizeSentence(c echo.Context) error {
	var req struct {
		Sentence string `json:"sentence"`
	}

	if err := c.Bind(&req); err != nil {
		fmt.Println("Error binding request:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.Sentence == "" {
		fmt.Println("Empty sentence field")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Sentence cannot be empty"})
	}

	capitalizedSentence := strings.ToUpper(req.Sentence)
	return c.JSON(http.StatusOK, map[string]string{"capitalized_sentence": capitalizedSentence})
}
