package service

import (
	"context"
	"log"

	"github.com/kiankw/petstore/model"
	"github.com/kiankw/petstore/service/dal"
)

func (s *srv) GetAllPets(ctx context.Context) ([]*model.Pet, error) {
	pets, err := s.dalClient.GetAllPets(dal.GetDB(ctx))
	if err != nil {
		log.Printf("Fail to get all pets with error %s\n", err)
	}
	return pets, err
}

func (s *srv) CreatePet(ctx context.Context, name string) (*model.Pet, error) {
	pet, err := s.dalClient.CreatePet(dal.GetDB(ctx), name)
	if err != nil {
		log.Printf("Fail to create pet %v with error %s\n", name, err)
	}
	return pet, err
}

func (s *srv) DeletePet(ctx context.Context, id int64) error {
	err := s.dalClient.DeletePet(dal.GetDB(ctx), id)
	if err != nil {
		log.Printf("Fail to delete pet %v with error %s\n", id, err)
	}
	return err
}

func (s *srv) GetPetByID(ctx context.Context, id int64) (*model.Pet, error) {
	pet, err := s.dalClient.GetPetByID(dal.GetDB(ctx), id)
	if err != nil {
		log.Printf("Fail to get pet %v with error %s\n", id, err)
	}
	return pet, err
}

func (s *srv) UpdatePet(ctx context.Context, id int64, name string) error {
	err := s.dalClient.UpdatePet(dal.GetDB(ctx), &model.Pet{Id: id, Name: name})
	if err != nil {
		log.Printf("Fail to update pet %v with error %s\n", id, err)
	}
	return err
}
