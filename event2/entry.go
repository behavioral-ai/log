package event2

import (
	"github.com/advanced-go/common/core"
	"time"
)

// Entry - timeseries access log struct
type Entry struct {
	CustomerId string    `json:"customer-id"`
	StartTime  time.Time `json:"start-time"`
	Duration   int64     `json:"duration"`
	Traffic    string    `json:"traffic"`
	CreatedTS  time.Time `json:"created-ts"`

	Origin core.Origin `json:"origin"`

	RequestId string `json:"request-id"`
	RelatesTo string `json:"relates-to"`
	Location  string `json:"location"`
	Protocol  string `json:"proto"`
	Method    string `json:"method"`
	From      string `json:"from"`
	To        string `json:"to"`
	Uri       string `json:"uri"`
	Path      string `json:"path"`
	Query     string `json:"query"`

	StatusCode int32  `json:"status-code"`
	Encoding   string `json:"encoding"`
	Bytes      int64  `json:"bytes"`

	Timeout        int32   `json:"timeout"`
	RateLimit      float64 `json:"rate-limit"`
	RateBurst      int32   `json:"rate-burst"`
	ControllerCode string  `json:"cc"`

	RouteTo      string `json:"route-to"`
	RoutePercent int    `json:"route-percent"`
	RouteCode    string `json:"rc"`
}
