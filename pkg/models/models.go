package models

import (
	"database/sql/driver"
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UniqueID struct for UUIDs
type UniqueID struct {
	UUID string `gorm:"type:varchar(36);primary_key"`
}

// Scan implements the Scanner interface
func (u *UniqueID) Scan(value interface{}) error {
	if value == nil {
		u.UUID = ""
		return nil
	}
	
	str, ok := value.(string)
	if !ok {
		return errors.New("failed to scan UniqueID")
	}
	u.UUID = str
	return nil
}

// Value implements the Valuer interface
func (u UniqueID) Value() (driver.Value, error) {
	return u.UUID, nil
}

// Client struct represents the client table
type Client struct {
	ID        UniqueID `gorm:"embedded;primary_key"`
	Name      string   `gorm:"type:varchar(100)"`
	Email     string   `gorm:"type:varchar(100);uniqueIndex"`
	Phone     string   `gorm:"type:varchar(20)"`
	CreatedAt time.Time
}
type ClientRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type MySQLRepo struct {
	db *gorm.DB
}

func NewMySQLRepo(dsn string) (*MySQLRepo, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("NewMySQLRepo: Failed to open database: %v", err)
		return nil, err
	}
	return &MySQLRepo{db: db}, nil
}