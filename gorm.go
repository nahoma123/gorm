package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	Open() (*gorm.DB, error)
	Migrate(models []interface{}) error
}

type gormDb struct {
	dbURL string
}

func Initialize(dbURL string) Gorm {
	return &gormDb{
		dbURL: dbURL}
}

func (grm *gormDb) Open() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(grm.dbURL), &gorm.Config{})
}

func (grm *gormDb) Migrate(models []interface{}) error {
	db, err := grm.Open()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(models...)
	return err
}

// At system initialization
// Check if the super admin exists by checking if the permission and roles table don't exist,
//
