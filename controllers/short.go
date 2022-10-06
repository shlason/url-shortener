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
	ShortID string `json:"shortId"`
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
				ShortID: u.ShortID,
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
			ShortID: u.ShortID,
		},
	})
}

// GetShortIDRediect godoc
// @Summary      Redirect by short ID
// @Description  Use 301 redirect by short ID
// @Tags         short
// @Accept       json
// @Produce      json
// @Param        shortID    query    string  true  "shortID redirect use"
// @Router       /{shortID} [get]
func GetShortIDRediect(c *gin.Context) {
	shortID := c.Param("shortID")
	u := &models.URL{
		ShortID: shortID,
	}
	result := u.ReadByShortID()

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.String(http.StatusBadRequest, "404 Not found.")
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error,
		})
		return
	}

	c.Redirect(http.StatusMovedPermanently, u.LongURL)
}
