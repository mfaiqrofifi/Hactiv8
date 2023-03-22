package service

import (
	"Task_Microservice/domain"
	"errors"
	"fmt"
)

func CreateBook(book domain.BookDomain) domain.BookDomain {
	book.BookId = fmt.Sprintf("B%d", len(domain.Books)+1)
	domain.Books = append(domain.Books, book)
	return book
}

func GetDataBook(idBook string) (domain.BookDomain, error) {
	// var bookData domain.BookDomain
	for _, v := range domain.Books {
		if idBook == v.BookId {
			return v, nil
		}
	}
	return domain.BookDomain{}, errors.New("Notfound")
}

func AllBook() ([]domain.BookDomain, error) {
	if len(domain.Books) == 0 {
		return nil, errors.New("Data is emphty")
	}
	return domain.Books, nil
}
