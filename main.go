package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize URL shortener service
	s := NewShortener() // This should match the definition in shortener.go

	// Routes
	e.POST("/shorten", s.ShortenURL)
	e.GET("/:shortURL", s.RedirectURL)
	e.GET("/metrics", s.GetMetrics)

	// Start server
	log.Fatal(e.Start(":8080"))
}
