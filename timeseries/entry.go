package timeseries

import "github.com/advanced-go/common/core"

type Entry struct {
	Origin         core.Origin `json:"origin"`
	Duration       int64       `json:"duration"`
	TimeoutCount   int         `json:"timeout-count"`
	RateLimitCount int         `json:"rate-limit-count"`
}

var (
	westIndex = 0
	westData  = []Entry{
		{Origin: core.Origin{Region: "us-west", Host: "www.west-host1.com"}, Duration: 65, TimeoutCount: 0, RateLimitCount: 2},
		{Origin: core.Origin{Region: "us-west", Host: "www.west-host1.com"}, Duration: 65, TimeoutCount: 0, RateLimitCount: 2},
	}

	centralIndex = 0
	centralData  = []Entry{
		{Origin: core.Origin{Region: "us-central", Host: "www.central-host1.com"}, Duration: 65, TimeoutCount: 0, RateLimitCount: 2},
		{Origin: core.Origin{Region: "us-central", Host: "www.central-host1.com"}, Duration: 65, TimeoutCount: 0, RateLimitCount: 2},
	}
)
