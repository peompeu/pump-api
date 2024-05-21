package domain

type Subscription struct {
	WantDealBaseRate float64

	Notifications []Notification

	UserID    uint
	CountryID uint
}
