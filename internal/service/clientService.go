package service

import (
	"fmt"
	"p1/config"
	"p1/pkg/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
)

// ClientService struct that manages client operations using GORM
type ClientService struct {
	db     *gorm.DB       // Pointer to the GORM DB instance
	client *models.Client // Pointer to a Client instance
}

// NewClientService creates a new instance of ClientService
func NewClientService() (*ClientService, error) {
	// Database connection using GORM
	MySQLDSN := config.MySQLDSN
	db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Migrator().DropTable(&models.Client{}); err != nil {
		return nil,fmt.Errorf("failed to drop tables: %w", err)
	}
	// Auto-migrate the client model to ensure table is created if it doesn't exist
	err = db.AutoMigrate(&models.Client{})
	if err != nil {
		return nil, fmt.Errorf("error migrating database: %w", err)
	}

	return &ClientService{
		db: db,
		client: &models.Client{
			//ID:        models.UniqueID{UUID: uuid.NewString()},
			CreatedAt: time.Now(),
		},
	}, nil
}

// InitClientServices initializes the client services and handles client data entry
func (cs *ClientService) InitClientServices(name string,email string,phone string) error {

	cs.client.Name = name
	cs.client.Email = email
	cs.client.Phone = phone
	rand.Seed(uint64(time.Now().UnixNano()))
	id :=  uuid.NewString()
	cs.client.ID = models.UniqueID{UUID:id}
	cs.client.CreatedAt = time.Now().UTC().Add(5*time.Hour + 30*time.Minute)
	// Insert the new client into the database using GORM
	result := cs.db.Create(&cs.client)
	if result.Error != nil {
		return fmt.Errorf("error inserting client into database: %w", result.Error)
	}

	fmt.Println("Client successfully added to the database!")

	// Optionally handle additional requirements
	return nil
}
