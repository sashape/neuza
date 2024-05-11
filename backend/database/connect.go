package database

import (
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

type Config struct {
	user string
	password string
	database string
	host string
}

func loadConfig() Config {
	return Config{
			user: os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASSWORD"),
			database: os.Getenv("POSTGRES_DB"),
			host: os.Getenv("POSTGRES_HOST"),
	}
}


func Connect() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := loadConfig()

	dsn := "host=" + config.host +
				 " user=" + config.user +
				 " password=" + config.password +
				 " dbname=" + config.database + " port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{DB: database}, err
}

func (d *Database) AutoMigrate(models ...interface{}) error {
	return d.DB.AutoMigrate(models...)
}
