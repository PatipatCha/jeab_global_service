package databases

import (
	"fmt"

	"github.com/PatipatCha/jeab_global_service/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Initialize connection string.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", configuration.AzureConfig().Host, configuration.AzureConfig().User, configuration.AzureConfig().Password, configuration.AzureConfig().Database)

	// Initialize connection object using GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully created connection to database")

	return db, nil
}
