package repository

import (

	// "ShowCase/repository/query"
	"ShowCase/model"
	"errors"
	"fmt"
)

type BookInterface interface {
	CreateBook(in model.BookDomain) (res model.BookDomain, err error)
	GetBookAll() (res []model.BookDomain, err error)
	GetBookbyID(id int) (res model.BookDomain, err error)
	UpdateBook(in model.BookDomain) (res model.BookDomain, err error)
	DeleteBook(id int) (msg string, err error)
}

func (r Repo) CreateBook(in model.BookDomain) (res model.BookDomain, err error) {
	// err = r.db.QueryRow(query.AddBook, in.Title, in.Author, in.Desc).Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	err = r.gorm.Create(&in).Scan(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) GetBookAll() (res []model.BookDomain, err error) {
	// var data model.BookDomain
	err = r.gorm.Find(&res).Error
	if err != nil {
		return res, err
	}
	// if len(res) == 0 {
	// 	err = errors.New("kosong")
	// 	return res, err
	// }
	// defer row.Close()
	// for row.Next() {
	// 	var book = model.BookDomain{}
	// 	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Desc)
	// 	res = append(res, book)
	// }
	return res, nil

}

func (r Repo) GetBookbyID(id int) (res model.BookDomain, err error) {
	// err = r.db.QueryRow(query.GetBookId, id).Scan(&res.ID, &res.Title, &res.Author, &res.Desc)
	err = r.gorm.First(&res, id).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

func (r Repo) UpdateBook(in model.BookDomain) (res model.BookDomain, err error) {
	// res, err := r.db.Exec(query.UpdateBook, in.ID, in.Title, in.Author, in.Desc)
	// if err != nil {
	// 	return "", err
	// }
	// count, err := res.RowsAffected()
	data := r.gorm.Model(&res).Where("id=?", in.ID).Updates(model.BookDomain{
		Title:  in.Title,
		Author: in.Author,
	}).Scan(&res)
	err = data.Error
	if err != nil {
		return res, err
	}
	if data.RowsAffected < 1 {
		return res, errors.New("Not Found")
	}
	return res, err

}
func (r Repo) DeleteBook(id int) (msg string, err error) {

	book := model.BookDomain{}
	data := r.gorm.Where("id=?", id).Delete(&book)

	fmt.Println(data.RowsAffected)
	err = data.Error
	if err != nil {
		return "", err
	}
	if data.RowsAffected < 1 {
		return "", errors.New("Not Found")
	}
	msg = fmt.Sprintf("Book deleted succesfully")
	return msg, err
}