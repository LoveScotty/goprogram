package mem

import (
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/LoveScotty/goprogram/internal/store/dto/bookstore"
)

type TBook struct {
	sync.RWMutex
	bookMap map[uint64]*bookstore.Book
}

var b = &TBook{
	bookMap: make(map[uint64]*bookstore.Book),
}

func newBook() *TBook {
	return b
}

func (s *TBook) Add(ctx *gin.Context, book *bookstore.Book) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.bookMap[book.Id]
	if ok {
		return bookstore.ErrAlreadyExist
	}
	s.bookMap[book.Id] = book

	return nil
}

func (s *TBook) Get(ctx *gin.Context, id uint64) (*bookstore.Book, error) {
	s.Lock()
	defer s.Unlock()
	book, ok := s.bookMap[id]
	if !ok {
		return nil, bookstore.ErrNotFound
	}

	return book, nil
}

func (s *TBook) Update(ctx *gin.Context, book *bookstore.Book) error {
	s.Lock()
	defer s.Unlock()
	oldBook, ok := s.bookMap[book.Id]
	if ok {
		return bookstore.ErrNotFound
	}
	newBook := *oldBook
	if oldBook.Name != book.Name {
		newBook.Name = book.Name
	}
	if len(book.AuthorList) > 0 {
		newBook.AuthorList = book.AuthorList
	}
	s.bookMap[book.Id] = &newBook

	return nil
}

func (s *TBook) All(ctx *gin.Context) ([]*bookstore.Book, error) {
	s.Lock()
	defer s.Unlock()
	allBookList := make([]*bookstore.Book, 0, len(s.bookMap))
	for _, book := range s.bookMap {
		allBookList = append(allBookList, book)
	}

	return allBookList, nil
}

func (s *TBook) Delete(ctx *gin.Context, id uint64) error {
	s.Lock()
	defer s.Unlock()
	_, ok := s.bookMap[id]
	if ok {
		return bookstore.ErrNotFound
	}
	delete(s.bookMap, id)

	return nil
}
