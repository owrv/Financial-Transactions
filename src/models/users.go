package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Id              int     `json:"id"`
	CurrentAge      int     `json:"currente_age"`
	RetirementAge   int     `json:"retirement_age"`
	BirthYear       int     `json:"birth_year"`
	BirthMonth      int     `json:"birth_month"`
	Gender          int     `json:"gender"` // 0 para homem 1 para mulher
	Address         string  `json:"address"`
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	PerCapitaIncome int     `json:"per_capita_income"`
	YearlyIncome    int     `json:"yearly_income"`
	TotalDebt       int     `json:"total_debt"`
	CreditScore     int     `json:"credit_score"`
	NumCreditCards  int     `json:"num_credit_cards"`
}
