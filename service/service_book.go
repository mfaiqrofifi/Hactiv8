package service

import (
	"Task_Microservice/domain"
	"errors"
	"fmt"
	"strconv"
)

func CreateBook(book domain.BookDomain) domain.BookDomain {
	var dt string
	val := 0
	if len(domain.Books) != 0 {
		dt = domain.Books[len(domain.Books)-1].BookId
		val, _ = strconv.Atoi(string(dt[1]))
	}
	book.BookId = fmt.Sprintf("B%d", val+1)
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
func UpdateBook(idBook string, data domain.BookDomain) (string, error) {
	data.BookId = idBook
	for i, v := range domain.Books {
		if v.BookId == idBook {
			domain.Books[i] = data
			s := fmt.Sprintf("data %s has been updated", idBook)
			return s, nil
		}
	}
	return "", errors.New("Can't find the data")
}
func DeleteBook(idBook string) (string, error) {
	condition := false
	var index int
	for i, v := range domain.Books {
		if v.BookId == idBook {
			index = i
			condition = true
			break
		}
	}
	if !condition {
		return "", errors.New("Data not Found")
	}
	copy(domain.Books[index:], domain.Books[index+1:])
	domain.Books[len(domain.Books)-1] = domain.BookDomain{}
	domain.Books = domain.Books[:len(domain.Books)-1]
	return "Book has been deleted", nil
}
