package event2

import (
	"context"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/log/common"
	"github.com/advanced-go/log/module"
	"github.com/advanced-go/log/testrsc"
	"github.com/advanced-go/postgresql/pgxsql"
	"net/http"
	"net/url"
)

func testOverride(h http.Header) http.Header {
	if h == nil {
		h = make(http.Header)
	}
	if h.Get(core.XExchangeResponse) == "" && h.Get(core.XExchangeStatus) == "" {
		h.Add(core.XExchangeResponse, testrsc.LOG2IngressEntry)
	}
	return h
}

func get[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, values url.Values) (entries []T, status *core.Status) {
	var e E

	if values == nil {
		return nil, core.StatusNotFound()
	}
	// Testing only
	h = testOverride(h)

	// Set XFrom so that PostgreSQL logging is correct.
	h = httpx.SetHeader(h, core.XFrom, module.Domain)
	entries, status = pgxsql.QueryT[T](ctx, h, common.AccessLogResource, common.AccessLogSelect, values)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
		return nil, status
	}
	if values != nil && len(values) > 0 {
		entries = filter[T](entries, values)
	}
	if len(entries) == 0 {
		status = core.NewStatus(http.StatusNotFound)
	}
	return
}

func filter[T pgxsql.Scanner[T]](entries []T, values url.Values) (result []T) {
	match := core.NewOrigin(values)
	customer := values.Get("customer")
	switch p := any(&result).(type) {
	case *[]Entry:
		if p != nil {
		}
		if entries2, ok := any(entries).([]Entry); ok {
			for _, e := range entries2 {
				if customer != "" && customer != e.CustomerId {
					continue
				}
				if core.OriginMatch(e.Origin, match) {
					*p = append(*p, e)
				}
			}
		}
	default:
	}
	return result
}
