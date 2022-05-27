package dal

import (
	"gorm.io/gorm"

	"github.com/kiankw/petstore/model"
)

type DAL interface {
	GetAllPets(db *gorm.DB) ([]*model.Pet, error)
	CreatePet(db *gorm.DB, name string) (*model.Pet, error)
	DeletePet(db *gorm.DB, id int64) error
	GetPetByID(db *gorm.DB, id int64) (*model.Pet, error)
	UpdatePet(db *gorm.DB, pet *model.Pet) error
}

func GetDAL() DAL {
	return &dalClient{}
}

type dalClient struct{}
