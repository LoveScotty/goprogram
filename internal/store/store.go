package store

import "github.com/LoveScotty/goprogram/internal/store/book"

type Factory interface {
	Book() book.Store
}
