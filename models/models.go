package models

// Stage info stage
type Stage struct {
	ID       int     `json:"id"`
	TryCnt   float64 `json:"try_cnt"`
	ClearCnt float64 `json:"clear_cnt"`
}

// balance
type Balance struct {
	Data string `json:"data"`
}

type BalanceData struct {
	Key string `json:"key"`
	Hp float64 `json:"hp"`
	Power float64 `json:"power"`
	Defense float64 `json:"defense"`
	Speed float64 `json:"speed"`
	Range float64 `json:"range"`
}

type Balances struct {
	Unit  []BalanceData `json:"unit"`
	Poker []BalanceData `json:"poker"`
}