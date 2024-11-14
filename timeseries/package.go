package timeseries

import (
	"context"
	"github.com/advanced-go/common/core"
)

const (
	PkgPath = "github/advanced-go/log/timeseries"
	Route   = "timeseries"
)

//type TimeUTC time.Time

/*
// PercentileThresholdSLO - ingress host, pre-calculated percentile thresholds
func PercentileThresholdSLO(ctx context.Context, origin core.Origin) (PercentileThreshold, *core.Status) {
	return NewPercentileThreshold(), core.StatusOK()
}

// PercentileThresholdQuery - ingress host, queryable percentile thresholds
func PercentileThresholdQuery(ctx context.Context, origin core.Origin, from, to TimeUTC) (PercentileThreshold, *core.Status) {
	return NewPercentileThreshold(), core.StatusOK()
}

// StatusCodeThresholdQuery - egress route, queryable status code thresholds
func StatusCodeThresholdQuery(ctx context.Context, origin core.Origin, from, to TimeUTC, statusCodes string) (StatusCodeThreshold, *core.Status) {
	return NewStatusCodeThreshold(), core.StatusOK()
}

// GetProfile - retrieve traffic profile
func GetProfile(ctx context.Context) (*Profile, *core.Status) {
	return NewProfile(), core.StatusOK()
}


*/

func Query(ctx context.Context, origin core.Origin) (Entry, *core.Status) {
	return Entry{}, core.StatusOK()
}
