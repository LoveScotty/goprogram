package mem

import (
	"sync"

	"github.com/LoveScotty/goprogram/internal/store"
	"github.com/LoveScotty/goprogram/internal/store/book"
	"github.com/LoveScotty/goprogram/internal/store/factory"
)

type memStore struct{}

var (
	memFactory store.Factory
	once       sync.Once
)

func init() {
	once.Do(func() {
		memFactory = &memStore{}
		factory.Register("bookstore", memFactory)
	})
}

func (mem *memStore) Book() book.Store {
	return newBook()
}
