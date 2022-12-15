package models

type Price struct {
	Ask        string `json:"ask"`
	Bid        string `json:"bid"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	CreateDate string `json:"create_date"`
	High       string `json:"high"`
	Low        string `json:"low"`
	Name       string `json:"name"`
	PctChange  string `json:"pctChange"`
	Timestamp  string `json:"timestamp"`
	VarBid     string `json:"varBid"`
}

type CoinToReal struct {
	EuroToReal          Price `json:"EURBRL"`
	DolarToReal         Price `json:"USDBRL"`
	PoundToReal         Price `json:"GBPBRL"`
	BitcoinToReal       Price `json:"BTCBRL"`
	DogecoinToReal      Price `json:"DOGEBRL"`
	CanadianDolarToReal Price `json:"CADBRL"`
}
