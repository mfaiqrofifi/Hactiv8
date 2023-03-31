package repository

import (
	"DDD/model"
	"DDD/repository/query"
	"fmt"
)

type BookInterface interface {
	CreateBook(in model.BookDomain) (res model.BookDomain, err error)
	GetBookAll() (res []model.BookDomain, err error)
	GetBookbyID(id int) (res model.BookDomain, err error)
	UpdateBook(in model.BookDomain) (msg string, err error)
	DeleteBook(id int) (msg string, err error)
}

func (r Repo) CreateBook(in model.BookDomain) (res model.BookDomain, err error) {
	err = r.db.QueryRow(query.AddBook, in.Title, in.Author, in.Desc).Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookAll() (res []model.BookDomain, err error) {
	row, err := r.db.Query(query.GetBook)
	if err != nil {
		return res, err
	}
	defer row.Close()
	for row.Next() {
		var book = model.BookDomain{}
		err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
		res = append(res, book)
	}
	return res, nil

}

func (r Repo) GetBookbyID(id int) (res model.BookDomain, err error) {
	err = r.db.QueryRow(query.GetBookId, id).Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateBook(in model.BookDomain) (msg string, err error) {
	res, err := r.db.Exec(query.UpdateBook, in.ID, in.Title, in.Author, in.Desc)
	if err != nil {
		return "", err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Updated data ammount %d", count), err

}
func (r Repo) DeleteBook(id int) (msg string, err error) {
	res, err := r.db.Exec(query.DeleteBook, id)
	if err != nil {
		return "", err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	msg = fmt.Sprintf("Deleted data ammount %d", count)
	return msg, err
}
