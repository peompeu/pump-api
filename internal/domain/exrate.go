package domain

type ExRate struct {
	// Telegraphic Transfer Buying의 약자로, 송금을 통해 외화를 살 때 적용되는 환율을 의미합니다.
	Ttb float64

	// Telegraphic Transfer Selling의 약자로, 송금을 통해 외화를 팔 때 적용되는 환율을 의미합니다.
	Tts float64

	// 외환 거래의 기준 환율을 의미합니다.
	DealBaseRate float64

	CountryID uint
}
