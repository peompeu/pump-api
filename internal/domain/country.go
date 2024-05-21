package domain

type Country struct {
	Name string
	Code string

	Subscriptions []Subscription
	ExRates       []ExRate
}
