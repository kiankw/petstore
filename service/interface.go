package service

import (
	"context"

	"github.com/kiankw/petstore/model"
	"github.com/kiankw/petstore/service/dal"
)

type Service interface {
	GetAllPets(ctx context.Context) ([]*model.Pet, error)
	CreatePet(ctx context.Context, name string) (*model.Pet, error)
	DeletePet(ctx context.Context, id int64) error
	GetPetByID(ctx context.Context, id int64) (*model.Pet, error)
	UpdatePet(ctx context.Context, id int64, name string) error
}

func NewService(dalClient dal.DAL) Service {
	return &srv{
		dalClient: dalClient,
	}
}

type srv struct {
	dalClient dal.DAL
}
