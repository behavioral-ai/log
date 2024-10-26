package test

import (
	"github.com/advanced-go/common/core"
	"github.com/advanced-go/common/test"
	"github.com/advanced-go/log/event2"
	http2 "github.com/advanced-go/log/http"
	"github.com/advanced-go/log/testrsc"
	//httpt "github.com/advanced-go/stdlib/httpx/httpxtest"
	"net/http"
	"reflect"
	"testing"
)

func TestLog2(t *testing.T) {
	tests := []struct {
		name   string
		req    *http.Request
		resp   *http.Response
		status *core.Status
	}{
		{name: "ingress-get-all", req: test.NewRequestTest(testrsc.LOG2IngressGetAllReq, t), resp: test.NewResponseTest(testrsc.LOG2IngressGetAllResp, t), status: core.StatusOK()},

		//
	}
	for _, tt := range tests {
		ok := true
		t.Run(tt.name, func(t *testing.T) {
			resp, status := http2.Exchange(tt.req)
			if tt.status != nil && status.Code != tt.status.Code {
				t.Errorf("Exchange() got status : %v, want status : %v, error : %v", status.Code, tt.status.Code, status.Err)
				ok = false
			}
			if ok && resp.StatusCode != tt.resp.StatusCode {
				t.Errorf("Exchange() got status code : %v, want status code : %v", resp.StatusCode, tt.resp.StatusCode)
				ok = false
			}
			var gotT []event2.Entry
			var wantT []event2.Entry
			if ok {
				gotT, wantT, ok = test.Deserialize[test.Output, []event2.Entry](resp.Body, tt.resp.Body, t)
			}
			if ok && len(gotT) != len(wantT) {
				t.Errorf("got length : %v, want length : %v", len(gotT), len(wantT))
				ok = false
			}
			if ok {
				for i := 0; i < len(wantT); i++ {
					if !reflect.DeepEqual(gotT[i], wantT[i]) {
						t.Errorf("\ngot content  = %v,\nwant content = %v\n", gotT[i], wantT[i])
						ok = false
					}
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
