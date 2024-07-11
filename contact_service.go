package services

import (
	"bitespeed/models"
	"time"

	"github.com/jinzhu/gorm"
)

type ContactService struct {
    db *gorm.DB
}

func NewContactService(db *gorm.DB) *ContactService {
    return &ContactService{db: db}
}

func (s *ContactService) Identify(input *models.Contact) (*models.Contact, error) {
    var contact models.Contact

    // Check if contact already exists
    if err := s.db.Where("email = ? OR phone_number = ?", input.Email, input.PhoneNumber).First(&contact).Error; err != nil {
        if gorm.IsRecordNotFoundError(err) {
            // Create new contact if not found
            input.LinkPrecedence = "primary"
            input.CreatedAt = time.Now()
            input.UpdatedAt = time.Now()
            if err := s.db.Create(&input).Error; err != nil {
                return nil, err
            }
            return input, nil
        }
        return nil, err
    }

    // Handle existing contact logic
    contact.UpdatedAt = time.Now()
    if err := s.db.Save(&contact).Error; err != nil {
        return nil, err
    }

    return &contact, nil
}
