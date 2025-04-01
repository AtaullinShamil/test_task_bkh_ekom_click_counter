package counter

import "time"

type ErrorResponse struct {
	Error string `json:"error"`
}

type StatsRequest struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type StatsResponse struct {
	Stats []Stat `json:"stats"`
}

type Stat struct {
	Timestamp time.Time `json:"ts" db:"timestamp" bson:"timestamp"`
	Value     int       `json:"v" db:"count" bson:"count"`
	BannerID  int       `json:"-" db:"banner_id" bson:"banner_id"`
}
