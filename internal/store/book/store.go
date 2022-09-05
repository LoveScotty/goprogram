package book

import (
	"github.com/gin-gonic/gin"

	"github.com/LoveScotty/goprogram/internal/store/dto/bookstore"
)

type Store interface {
	Add(ctx *gin.Context, book *bookstore.Book) error
	Get(ctx *gin.Context, id uint64) (*bookstore.Book, error)
	Update(ctx *gin.Context, book *bookstore.Book) error
	All(ctx *gin.Context) ([]*bookstore.Book, error)
	Delete(ctx *gin.Context, id uint64) error
}
