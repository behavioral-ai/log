package event2

import (
	"context"
	"errors"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/httpx"
	"github.com/behavioral-ai/log/common"
	"github.com/behavioral-ai/log/module"
	"github.com/behavioral-ai/postgres/pgxsql"
	"net/http"
)

// put - function to Put a set of entries into a datastore
func put[E core.ErrorHandler, T pgxsql.Scanner[T]](ctx context.Context, h http.Header, body []T) (h2 http.Header, status *core.Status) {
	var e E

	h2 = httpx.SetHeader(h2, httpx.ContentType, httpx.ContentTypeText)
	if len(body) == 0 {
		status = core.NewStatusError(core.StatusInvalidContent, errors.New("error: no entries found"))
		return nil, status
	}
	h = httpx.SetHeader(h, core.XFrom, module.Domain)
	_, status = pgxsql.InsertT[T](ctx, h, common.AccessLogResource, common.AccessLogInsert, body)
	if !status.OK() {
		e.Handle(status.WithRequestId(h))
	}
	return
}
