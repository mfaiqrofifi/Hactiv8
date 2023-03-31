package handler

import "DDD/service"

type Handler struct {
	app service.ServiceInterface
}

func NewServer(service service.ServiceInterface) Handler {
	return Handler{app: service}
}
