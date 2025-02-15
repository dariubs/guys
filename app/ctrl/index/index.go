package index

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Index struct {
	DB *gorm.DB
}

func NewIndex(db *gorm.DB) *Index {
	return &Index{DB: db}
}

func (i *Index) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	}
}

func (i *Index) Username() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "username.html", gin.H{})
	}
}
