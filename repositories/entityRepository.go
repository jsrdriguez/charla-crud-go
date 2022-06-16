package repositories

import (
	"jsrdriguez/crud-go/models"
	"jsrdriguez/crud-go/requests"
)

func NewEntity(name, idetificationNumber, email string) *models.Entity {
	return &models.Entity{
		Name:                name,
		IdetificationNumber: idetificationNumber,
		Email:               email,
	}
}

func SaveEntity(entity *requests.EntityRequestInsert) {
	if err := models.DB().Model(models.Entity{}).Create(&entity).Error; err != nil {
		panic(err)
	}
}

func UpdateEntity(id int, entity *models.Entity) {
	models.DB().Model(&models.Entity{}).Where("entity_id = ?", id).Updates(entity)
}

func DeleteEntity(id int) {
	models.DB().Where("entity_id = ?", id).Delete(&models.Entity{})
}

func GetEntities() models.Entities {
	var entities models.Entities

	if err := models.DB().Find(&entities).Error; err != nil {
		panic(err)
	}

	return entities
}

func GetEntity(id int) models.Entity {
	var entity models.Entity

	if err := models.DB().Find(&entity, id).Error; err != nil {
		panic(err)
	}

	return entity
}
