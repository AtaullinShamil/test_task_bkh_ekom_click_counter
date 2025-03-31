package counter

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type StatsRequest struct {
	From time.Time `json:"tsFrom"`
	To   time.Time `json:"tsTo"`
}

type StatsResponse struct {
	Stats []Stat `json:"stats"`
}

type Stat struct {
	Timestamp time.Time `json:"ts" db:"timestamp"`
	Value     int       `json:"v" db:"count"`
}
