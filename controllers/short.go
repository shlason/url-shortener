package controllers

import (
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/models"
	"github.com/shlason/url-shortener/utils"
	"gorm.io/gorm"
)

type createShortURLRequestPayload struct {
	URL string `json:"url"`
}

type createShortURLResponsePayload struct {
	URL string `json:"url"`
}

func CreateShortURL(c *gin.Context) {
	var r createShortURLRequestPayload
	err := c.BindJSON(&r)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Content-Type is not JSON format",
			"data":    nil,
		})
		return
	}
	_, err = url.ParseRequestURI(r.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid URL",
			"data":    nil,
		})
		return
	}
	u := &models.URL{
		LongURL: r.URL,
	}
	result := u.ReadByLongURL()
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"data": createShortURLResponsePayload{
				URL: u.ShortID,
			},
		})
		return
	}
	u.ShortID = utils.Base62.Encode(time.Now().UnixMilli())
	result = u.Create()
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": createShortURLResponsePayload{
			URL: u.ShortID,
		},
	})
}

func GetShortIDRediect(c *gin.Context) {
	shortID := c.Param("shortID")
	u := &models.URL{
		ShortID: shortID,
	}
	result := u.ReadByShortID()
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.String(http.StatusBadRequest, "404 Not found.")
		return
	}
	c.Redirect(http.StatusMovedPermanently, u.LongURL)
}
