package handler

import "github.com/kiankw/petstore/service"

type handlerStruct struct {
	srv service.Service
}

func NewHandler(s service.Service) *handlerStruct {
	return &handlerStruct{
		srv: s,
	}
}
