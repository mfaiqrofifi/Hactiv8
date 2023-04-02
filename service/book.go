package service

import "ShowCase/model"

type BookInterface interface {
	CreateBook(in model.BookDomain) (res model.BookDomain, err error)
	GetBook() (res []model.BookDomain, err error)
	GetBookbyID(id int) (res model.BookDomain, err error)
	UpdateBook(in model.BookDomain) (res model.BookDomain, err error)
	DeleteBook(id int) (msg string, err error)
}

func (s *Service) CreateBook(in model.BookDomain) (res model.BookDomain, err error) {
	res, err = s.repo.CreateBook(in)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBook() (res []model.BookDomain, err error) {
	res, err = s.repo.GetBookAll()
	if err != nil {
		return res, err
	}
	return res, nil
}
func (s *Service) GetBookbyID(id int) (res model.BookDomain, err error) {
	res, err = s.repo.GetBookbyID(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UpdateBook(in model.BookDomain) (res model.BookDomain, err error) {
	res, err = s.repo.UpdateBook(in)
	if err != nil {
		return res, err
	}
	return res, err
}
func (s *Service) DeleteBook(id int) (msg string, err error) {
	msg, err = s.repo.DeleteBook(id)
	if err != nil {
		return msg, err
	}
	return msg, err
}
