package timeseries

import (
	"context"
	"github.com/advanced-go/common/core"
)

const (
	PkgPath    = "github/advanced-go/log/timeseries"
	Route      = "timeseries"
	WestRegion = "us-west"
	WestZoneA  = "w-a"
	WestZoneB  = "w-b"

	CentralRegion = "us-central"
	CentralZoneA  = "c-a"
	CentralZoneB  = "c-b"

	EastRegion = "us-east"
	EastZoneA  = "e-a"
	EastZoneB  = "e-b"
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
