package db

import (
	models2 "Credits/internal/app/models"
	"Credits/internal/app/models/seeds"
	"Credits/pkg/logger"
	"errors"
)

func Migrate() error {
	if dbConn == nil {
		logger.Error.Printf("[db.Migrate] Error because database connection is nil")

		return errors.New("database connection is not initialized")
	}

	err := dbConn.AutoMigrate(
		&models2.CreditsStatus{},
		&models2.Credits{},
		&models2.CreditsComment{},
	)
	if err != nil {
		logger.Error.Printf("[db.Migrate] Error migrating tables: %v", err)

		return err
	}

	err = seeds.SeedCreditsStatuses(dbConn)
	if err != nil {
		return err
	}

	return nil
}
