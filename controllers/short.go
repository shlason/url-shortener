package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shlason/url-shortener/configs"
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

// CreatShortURL godoc
// @Summary      建立短網址
// @Description  藉由 timestamp 轉 base62 的方式產生 unique ID 來作為短網址的 ID
// @Tags         short
// @Accept       json
// @Produce      json
// @Param        url   body      string  true  "Original URL"
// @Success      200  {object}  createShortURLResponsePayload
// @Router       /short [post]
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

// GetShortIDRediect godoc
// @Summary      Redirect by short ID
// @Description  Use 301 redirect by short ID
// @Tags         short
// @Accept       json
// @Produce      json
// @Param        shortID    path    string  true  "shortID redirect use"
// @Router       /{shortID} [get]
func GetShortIDRediect(c *gin.Context) {
	shortID := c.Param("shortID")
	u := &models.URL{
		ShortID: shortID,
	}
	result := u.ReadByShortID()

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.Redirect(http.StatusMovedPermanently, fmt.Sprintf("https://%s/notfound", configs.Server.Host))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, u.LongURL)
}
