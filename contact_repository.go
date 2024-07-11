package repositories

import (
	"bitespeed/models"

	"github.com/jinzhu/gorm"
)

type ContactRepository struct {
    db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
    return &ContactRepository{db: db}
}

func (r *ContactRepository) FindByEmailOrPhone(email string, phone string) (*models.Contact, error) {
    var contact models.Contact
    if err := r.db.Where("email = ? OR phone_number = ?", email, phone).First(&contact).Error; err != nil {
        return nil, err
    }
    return &contact, nil
}

func (r *ContactRepository) Create(contact *models.Contact) error {
    return r.db.Create(contact).Error
}
