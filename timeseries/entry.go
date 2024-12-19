package timeseries

import "github.com/behavioral-ai/core/core"

const (
	westAZone    = "us-west1-a"
	westBZone    = "us-west1-b"
	centralAZone = "us-central1-a"
	centralBZone = "us-central1-b"
)

type Entry struct {
	Origin   core.Origin `json:"origin"`
	Latency  int         `json:"latency"`  // Milliseconds for a given percentile
	Gradient int         `json:"gradient"` // Rate of change
	RPS      int         `json:"rps"`      // Requests per second
}

var (
	westAIndex   = 0
	westAService = []Entry{
		{Origin: core.Origin{Region: WestRegion, Zone: WestZoneA, Host: "host1.com"}, Latency: 1500, Gradient: 5},
	}

	westBIndex   = 0
	westBService = []Entry{
		{Origin: core.Origin{Region: WestRegion, Zone: WestZoneB, Host: "host2.com"}, Latency: 1000, Gradient: 45},
	}

	centralAIndex   = 0
	centralAService = []Entry{
		{Origin: core.Origin{Region: CentralRegion, Zone: CentralZoneA, Host: "host3.com"}, Latency: 2000, Gradient: 80},
	}

	centralBIndex   = 0
	centralBService = []Entry{
		{Origin: core.Origin{Region: CentralRegion, Zone: CentralZoneB, Host: "host4.com"}, Latency: 1750, Gradient: 0},
	}
)

func entry(index *int, data []Entry) (Entry, *core.Status) {
	if *index >= len(data) {
		return Entry{}, core.StatusNotFound()
	}
	e := data[*index]
	*index++
	return e, core.StatusOK()
}
