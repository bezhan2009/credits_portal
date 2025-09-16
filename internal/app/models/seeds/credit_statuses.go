package seeds

import (
	"Credits/internal/app/models"
	"Credits/pkg/logger"
	"errors"
	"gorm.io/gorm"
)

func SeedCreditsStatuses(db *gorm.DB) error {
	creditStatuses := []models.CreditsStatus{
		{ID: 1, Name: "Принято"},
		{ID: 2, Name: "Передано в обработку"},
		{ID: 3, Name: "Обработано"},
		{ID: 4, Name: "Отказано"},
		{ID: 5, Name: "Доработка"},
	}

	for _, statuses := range creditStatuses {
		var existingStatus models.CreditsStatus
		if err := db.First(&existingStatus, "name = ?", statuses.Name).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				db.Create(&statuses)
			} else {
				logger.Error.Printf("[seeds.SeedCreditsStatuses] Error seeding credits Statuses: %v", err)

				return err
			}
		}
	}

	return nil
}
