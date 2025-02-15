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
		c.HTML(http.StatusOK, "index.html", gin.H{})
	}
}
