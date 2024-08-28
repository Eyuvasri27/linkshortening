package main

import (
	"crypto/sha1"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type Shortener struct {
	store    *URLStore
	urlMutex sync.RWMutex
}

func NewShortener() *Shortener {
	return &Shortener{
		store: NewURLStore(),
	}
}

type URLRequest struct {
	URL string `json:"url"`
}

func (s *Shortener) ShortenURL(c echo.Context) error {
	var req URLRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	originalURL := req.URL
	if originalURL == "" {
		return c.JSON(http.StatusBadRequest, "URL is required")
	}

	s.urlMutex.RLock()
	shortURL, exists := s.store.GetShortURL(originalURL)
	s.urlMutex.RUnlock()

	if exists {
		return c.JSON(http.StatusOK, map[string]string{"shortURL": shortURL})
	}

	shortURL = s.generateShortURL(originalURL)
	s.urlMutex.Lock()
	s.store.StoreURL(originalURL, shortURL)
	s.urlMutex.Unlock()

	return c.JSON(http.StatusOK, map[string]string{"shortURL": shortURL})
}

func (s *Shortener) RedirectURL(c echo.Context) error {
	shortURL := c.Param("shortURL")
	originalURL, exists := s.store.GetOriginalURL(shortURL)
	if !exists {
		return c.JSON(http.StatusNotFound, "URL not found")
	}
	return c.Redirect(http.StatusMovedPermanently, originalURL)
}

func (s *Shortener) GetMetrics(c echo.Context) error {
	metrics := s.store.GetTopDomains(3)
	return c.JSON(http.StatusOK, metrics)
}

func (s *Shortener) generateShortURL(url string) string {
	h := sha1.New()
	h.Write([]byte(url))
	return hex.EncodeToString(h.Sum(nil))[:8]
}
