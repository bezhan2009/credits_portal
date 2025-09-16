package service

import (
	"Credits/internal/app/models"
	"Credits/internal/repository"
)

func GetAllCredits(month, year int, afterID, statusID string) (credits []models.Credits, err error) {
	if statusID == "" {
		statusID = "1"
	}

	credits, err = repository.GetAllCredits(month, year, afterID, statusID)
	if err != nil {
		return nil, err
	}

	return credits, nil
}

func GetCreditById(creditID int) (credit models.Credits, err error) {
	credit, err = repository.GetCreditById(creditID)
	if err != nil {
		return models.Credits{}, err
	}

	return credit, nil
}

func CreateCredit(credit *models.Credits) (createdCredit models.Credits, err error) {
	if err = repository.CreateCredit(credit); err != nil {
		return models.Credits{}, err
	}

	return *credit, nil
}

func UpdateCredit(credit models.Credits) (err error) {
	if err = repository.UpdateCredit(credit); err != nil {
		return err
	}

	return nil
}

func DeleteCredit(creditID uint) (err error) {
	_, err = repository.GetCreditById(int(creditID))
	if err != nil {
		return err
	}

	if err = repository.DeleteCredit(creditID); err != nil {
		return err
	}

	return nil
}
