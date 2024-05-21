package domain

import "gorm.io/gorm"

type Notification struct {
	gorm.Model

	Readed bool

	UserID         uint
	SubscriptionID uint
}
