package timeseries

import (
	"context"
	"github.com/advanced-go/common/core"
)

const (
	PkgPath       = "github/advanced-go/log/timeseries"
	Route         = "timeseries"
	centralRegion = "us-central"
	westRegion    = "us-west"
)

func Query(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	switch origin.Region {
	case centralRegion:
		if centralIndex >= len(centralData) {
			return Entry{}, core.StatusNotFound()
		}
		e := centralData[centralIndex]
		centralIndex++
		return e, core.StatusOK()
	case westRegion:
		if westIndex >= len(westData) {
			return Entry{}, core.StatusNotFound()
		}
		e := westData[westIndex]
		westIndex++
		return e, core.StatusOK()
	}
	return Entry{}, core.StatusNotFound()
}

func Reset() {
	centralIndex = 0
	westIndex = 0
}
