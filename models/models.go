package models

// User basic info
type User struct {
	ID         int    `json:"id"`
	PlatformID string `json:"platform_id"`
	Platform   string `json:"platform"`
	DeviceID   string `json:"device_id"`
}

// Rank rank info
type Rank struct {
	ID    int     `json:"id"`
	Score float64 `json:"score"`
}

// Stage info stage
type Stage struct {
	ID    int     `json:"id"`
	Try   float64 `json:"try"`
	Clear float64 `json:"clear"`
}
