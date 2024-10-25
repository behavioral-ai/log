package test

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/test"
	"github.com/advanced-go/log/event1"
	http2 "github.com/advanced-go/log/http"
	"github.com/advanced-go/log/testrsc"
	"net/http"
	"reflect"
	"testing"
)

func TestLog1(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		{name: "ingress-get-all", req: test.NewRequestTest(testrsc.LOG1IngressGetAllReq, t), resp: test.NewResponseTest(testrsc.LOG1IngressGetAllResp, t), status: core.StatusOK()},
		{name: "ingress-get-not-found", req: test.NewRequestTest(testrsc.LOG1IngressGetNotFoundReq, t), resp: test.NewResponseTest(testrsc.NotFoundResp, t), status: nil},
		{name: "ingress-put-ok", req: test.NewRequestTest(testrsc.LOG1IngressPutReq, t), resp: test.NewResponseTest(testrsc.OKResp, t), status: nil},

		//
	}
	for _, tt := range tests {
		success := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				success = false
			}
			if success && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				success = false
			}
			var gotT []event1.Entry
			var wantT []event1.Entry
			if success {
				gotT, wantT, success = test.Deserialize[test.Output, []event1.Entry](resp.Body, tt.resp.Body, t)
			}
			if success {
				if !reflect.DeepEqual(gotT, wantT) {
					t.Errorf("Exchange() got = %v,\n want %v", gotT, wantT)
				}
			}
		})
	}
}

/*
	if !reflect.DeepEqual(got, tt.want) {
		t.Errorf("Exchange() got = %v, want %v", got, tt.want)
	}
	if !reflect.DeepEqual(got1, tt.want1) {
		t.Errorf("Exchange() got1 = %v, want %v", got1, tt.want1)
	}

*/
