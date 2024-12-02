package http

import (
	"fmt"
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/httpx"
	"net/http"
)

const (
	versionFmt = "{\n \"version\": \"%v\"\n}"
	domainFmt  = "{\n \"domain\": \"%v\"\n}"
)

func NewVersionResponse(version string) *http.Response {
	content := fmt.Sprintf(versionFmt, version)
	resp, _ := httpx.NewResponse(http.StatusOK, httpx.SetHeader(nil, httpx.ContentType, httpx.ContentTypeText), content)
	return resp
}

func NewDomainResponse(domain string) *http.Response {
	content := fmt.Sprintf(domainFmt, domain)
	resp, _ := httpx.NewResponse(http.StatusOK, httpx.SetHeader(nil, core.XDomain, domain), content)
	return resp
}
