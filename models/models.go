package models

// Stage info stage
type Stage struct {
	ID    int     `json:"id"`
	Try   float64 `json:"try"`
	Clear float64 `json:"clear"`
}
