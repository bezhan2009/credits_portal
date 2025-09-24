package models

import (
	"gorm.io/gorm"
	"time"
)

type CreditsStatus struct {
	ID   uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"type:varchar(40)" json:"name"`
}

type Credits struct {
	gorm.Model

	Name       string `json:"name" form:"name" gorm:"varchar(100)"`
	Surname    string `json:"surname" form:"surname" gorm:"varchar(100)"`
	Patronymic string `json:"patronymic" form:"patronymic" gorm:"varchar(100)"`

	Phone      string `gorm:"type:varchar(20);not null" json:"phone" form:"phone"`
	INN        string `gorm:"type:varchar(50)" json:"inn" form:"inn"`
	ClientCode string `gorm:"type:varchar(50)" json:"client_code" form:"client_code"`

	LoanType     string `gorm:"type:varchar(100)" json:"loan_type" form:"loan_type"`
	Address      string `gorm:"type:text" json:"address" form:"address"`
	BranchOffice string `gorm:"type:varchar(200)" json:"branch_office" form:"branch_office"`

	FrontSideOfThePassport string `gorm:"type:varchar(255)" json:"front_side_of_the_passport" form:"front_side_of_the_passport"`
	BackSideOfThePassport  string `gorm:"type:varchar(255)" json:"back_side_of_the_passport" form:"back_side_of_the_passport"`
	SelfieWithPassport     string `gorm:"type:varchar(255)" json:"selfie_with_passport" form:"selfie_with_passport"`

	Workplace           string    `gorm:"type:varchar(200)" json:"workplace" form:"workplace"`
	EmploymentDate      time.Time `json:"employment_date" form:"employment_date"`
	Salary              float64   `gorm:"type:decimal(15,2)" json:"salary" form:"salary"`
	IncomeProofDocument string    `gorm:"type:varchar(500)" json:"income_proof_document" form:"income_proof_document"`

	AdditionalIncomeSource string  `gorm:"type:text" json:"additional_income_source" form:"additional_income_source"`
	AdditionalIncomeAmount float64 `gorm:"type:decimal(15,2)" json:"additional_income_amount" form:"additional_income_amount"`

	LoanPurpose string  `gorm:"type:text" json:"loan_purpose" form:"loan_purpose"`
	LoanTerm    int     `json:"loan_term" form:"loan_term"`
	LoanAmount  float64 `gorm:"type:decimal(15,2)" json:"loan_amount" form:"loan_amount"`

	RequestCreator string `json:"request_creator" form:"request_creator"`

	CreditStatusID uint          `json:"credit_status_id" form:"credit_status_id"  gorm:"default:1"`
	CreditsStatus  CreditsStatus `gorm:"foreignKey:CreditStatusID;references:ID" json:"credits_status" form:"credits_status"`
}

type CreditsComment struct {
	gorm.Model

	Status      bool   `gorm:"type:boolean;not null" json:"status"`
	Description string `gorm:"type:text" json:"description"`

	CreditsID uint    `gorm:"type:integer;not null" json:"credits_id"`
	Credits   Credits `gorm:"foreignKey:CreditsID;references:ID" json:"credits"`
}
