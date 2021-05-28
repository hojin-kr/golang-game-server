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