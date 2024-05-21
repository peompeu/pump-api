package domain

import "gorm.io/gorm"

type Contry struct {
	*gorm.Model

	Name string
	Code string

	Subscriptions []Subscription
	ExRates       []ExRate
}
