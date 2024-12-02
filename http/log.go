package http

import (
	"errors"
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"github.com/advanced-go/common/uri"
	"github.com/advanced-go/log/event1"
	"github.com/advanced-go/log/event2"
	"github.com/advanced-go/log/module"
	"net/http"
)

func logExchange(r *http.Request, p *uri.Parsed) (*http.Response, *core.Status) {
	h2 := make(http.Header)
	h2.Add(httpx.ContentType, httpx.ContentTypeText)

	if p == nil {
		p1, err := uri.ValidateURL(r.URL, module.Domain)
		if err != nil {
			return httpx.NewResponse(http.StatusBadRequest, h2, err)
		}
		p = p1
	}

	switch r.Method {
	case http.MethodGet:
		return logGet(r, p)
	case http.MethodPut:
		return logPut(r, p)
	default:
		status := core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("error invalid method: [%v]", r.Method)))
		return httpx.NewResponse(status.HttpCode(), h2, status.Err)
	}
}

func logGet(r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var buf []byte
	var h2 http.Header

	switch p.Version {
	case ver1, "":
		buf, h2, status = event1.Get(r, p.Path)
	case ver2:
		buf, h2, status = event2.Get(r, p.Path)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	if !status.OK() {
		resp, _ = httpx.NewResponse(status.HttpCode(), h2, status.Err)
		return resp, status
	}
	return httpx.NewResponse(status.HttpCode(), h2, buf)
}

func logPut(r *http.Request, p *uri.Parsed) (resp *http.Response, status *core.Status) {
	var h2 http.Header

	switch p.Version {
	case ver1, "":
		h2, status = event1.Put(r, p.Path, nil)
	case ver2:
		h2, status = event2.Put(r, p.Path, nil)
	default:
		status = core.NewStatusError(http.StatusBadRequest, errors.New(fmt.Sprintf("invalid version: [%v]", r.Header.Get(core.XVersion))))
	}
	return httpx.NewResponse(status.HttpCode(), h2, status.Err)
}
