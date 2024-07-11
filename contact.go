package models

import "time"

type Contact struct {
    ID             uint `gorm:"primary_key"`
    PhoneNumber    string
    Email          string
    LinkedID       uint
    LinkPrecedence string
    CreatedAt      time.Time
    UpdatedAt      time.Time
    DeletedAt      *time.Time
}
