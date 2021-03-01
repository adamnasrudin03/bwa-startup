package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}
type Repository interface {


}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

