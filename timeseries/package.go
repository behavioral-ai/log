package timeseries

import (
	"context"
	"github.com/advanced-go/common/core"
)

const (
	PkgPath    = "github/advanced-go/log/timeseries"
	Route      = "timeseries"
	WestRegion = "us-west1"
	WestZoneA  = "us-west1-a"
	WestZoneB  = "us-west1-b"

	CentralRegion = "us-central1"
	CentralZoneA  = "us-central1-a"
	CentralZoneB  = "us-central1-b"

	EastRegion = "us-east1"
	EastZoneA  = "us-east1-a"
	EastZoneB  = "us-east1-b"
)

func Query(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	switch origin.Zone {
	case CentralZoneA:
		return entry(&centralAIndex, centralAService)
	case CentralZoneB:
		return entry(&centralBIndex, centralBService)
	case WestZoneA:
		return entry(&westAIndex, westAService)
	case WestZoneB:
		return entry(&westBIndex, westBService)
	}
	return Entry{}, core.StatusNotFound()
}

func Reset() {
	centralAIndex = 0
	centralBIndex = 0
	westAIndex = 0
	westBIndex = 0
}
