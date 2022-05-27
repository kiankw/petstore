package dal

import (
	"gorm.io/gorm"

	"github.com/kiankw/petstore/model"
)

func (d *dalClient) GetAllPets(db *gorm.DB) (pets []*model.Pet, err error) {
	err = db.Find(&pets).Error
	return pets, err
}
func (d *dalClient) CreatePet(db *gorm.DB, name string) (*model.Pet, error) {
	pet := &model.Pet{Name: name}
	err := db.Create(&pet).Error
	return pet, err
}
func (d *dalClient) DeletePet(db *gorm.DB, id int64) error {
	return db.Delete(&model.Pet{Id: id}).Error
}
func (d *dalClient) GetPetByID(db *gorm.DB, id int64) (*model.Pet, error) {
	pet := &model.Pet{Id: id}
	err := db.First(&pet).Error
	return pet, err
}
func (d *dalClient) UpdatePet(db *gorm.DB, pet *model.Pet) error {
	return db.Save(&pet).Error
}
