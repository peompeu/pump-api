package domain

import "gorm.io/gorm"

type Country struct {
	gorm.Model

	Name string
	Code string

	Subscriptions []Subscription
	ExRates       []ExRate
}
