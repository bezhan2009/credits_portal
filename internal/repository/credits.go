package repository

import (
	"Credits/internal/app/models"
	"Credits/pkg/db"
	"Credits/pkg/logger"
	"strconv"
)

func GetAllCredits(month, year int, afterID, statusID string) (credits []models.Credits, err error) {
	query := db.GetDBConn().Model(&models.Credits{}).
		Where("EXTRACT(MONTH FROM created_at) = ? AND EXTRACT(YEAR FROM created_at) = ?", month, year).
		Where("credit_status_id = ?", statusID).
		Order("id ASC").
		Limit(10)

	if afterID != "" && afterID != "0" {
		if id, err := strconv.ParseUint(afterID, 10, 64); err == nil && id > 0 {
			query = query.Where("id > ?", id)
		} else {
			logger.Error.Printf("[repository.GetAllCredits] Invalid afterID: %s\n", afterID)
		}
	}

	if err = query.Find(&credits).Error; err != nil {
		logger.Error.Printf("[repository.GetAllCredits] Error finding all credits: %v", err)

		return nil, TranslateGormError(err)
	}

	return credits, nil
}

func GetCreditById(creditID int) (credit models.Credits, err error) {
	if err = db.GetDBConn().Model(&models.Credits{}).First(&credit, creditID).Error; err != nil {
		logger.Error.Printf("[repository.GetCreditsById] Error finding credits: %v", err)

		return credit, TranslateGormError(err)
	}

	return credit, nil
}

func CreateCredit(credit *models.Credits) (err error) {
	if err = db.GetDBConn().Model(&models.Credits{}).Create(credit).Error; err != nil {
		logger.Error.Printf("[repository.CreateCredit] Error creating credits: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func UpdateCredit(credit models.Credits) (err error) {
	if err = db.GetDBConn().Model(&models.Credits{}).Where("id = ?", credit.ID).Updates(&credit).Error; err != nil {
		logger.Error.Printf("[repository.UpdateCredit] Error updating credits: %v", err)

		return TranslateGormError(err)
	}

	return nil
}

func DeleteCredit(creditID uint) (err error) {
	if err = db.GetDBConn().Model(&models.Credits{}).Delete(&models.Credits{}, creditID).Error; err != nil {
		logger.Error.Printf("[repository.DeleteCredit] Error deleting credits: %v", err)
		return TranslateGormError(err)
	}

	return nil
}
