package bookstore

import "errors"

var (
	ErrNotFound     = errors.New("book is not found")
	ErrAlreadyExist = errors.New("book is already exist")
)

type Book struct {
	Id         uint64   `json:"id"`
	Name       string   `json:"name"`
	AuthorList []string `json:"author_list"`
}
