package domain

import "gorm.io/gorm"

type Subscription struct {
	gorm.Model

	WantDealBaseRate float64

	Notifications []Notification

	UserID    uint
	CountryID uint
}
