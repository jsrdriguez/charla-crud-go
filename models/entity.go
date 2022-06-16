package models

type Entity struct {
	EntityId            uint   `json:"entity_id" gorm:"primarykey;<-:create"`
	Name                string `json:"name"`
	IdetificationNumber string `json:"idetificationNumber"`
	Email               string `json:"email"`
}

type Entities []Entity
