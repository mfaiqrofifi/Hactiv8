package service

import "DDD/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	BookInterface
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
