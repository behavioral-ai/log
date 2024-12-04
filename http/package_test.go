package http

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/jsonx"
	"github.com/behavioral-ai/log/event1"
	"net/http"
)

func ExampleExchange_Invalid() {
	resp, status := Exchange(nil)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ := http.NewRequest("", "http://www.google.com/search?q=golang", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	req, _ = http.NewRequest("", "http://www.google.com/github/behavioral-ai/log", nil)
	resp, status = Exchange(req)
	fmt.Printf("test: Exchange(nil) -> [status:%v] [status-code:%v]\n", status, resp.StatusCode)

	//Output:
	//test: Exchange(nil) -> [status:Bad Request] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, domain does not match: "/search" "github/behavioral-ai/log"]] [status-code:400]
	//test: Exchange(nil) -> [status:Bad Request [error: invalid URI, path only contains a domain: "/github/behavioral-ai/log"]] [status-code:400]

}

func ExampleExchange_Domain() {
	r, _ := http.NewRequest("", "http://localhost:8083/github/behavioral-ai/log:domain", nil)
	resp, status := Exchange(r)
	if status.OK() {
		//buf, _ := io.ReadAll(resp.Body, nil)
		fmt.Printf("test: Exchange(r) -> [status:%v] [status-code:%v] [%v]\n", status, resp.StatusCode, resp.Header.Get(core.XDomain))
	}

	//Output:
	//test: Exchange(r) -> [status:OK] [status-code:200] [github/behavioral-ai/log]

}

func _ExampleExchange_Timeseries_dbClient_Error() {
	uri := "http://localhost:8081/github/behavioral-ai/observation:v1/timeseries/egress?region=*"
	req, _ := http.NewRequest(http.MethodGet, uri, nil)

	resp, status := Exchange(req)
	if !status.OK() {
		fmt.Printf("test: Exchange() -> [status:%v]\n", status)
	} else {
		entries, status1 := jsonx.New[[]event1.Entry](resp.Body, resp.Header)
		fmt.Printf("test: Exchange() -> [status:%v] [status-code:%v] [bytes:%v] [count%v]\n", status1, resp.StatusCode, resp.ContentLength, len(entries))
	}

	//Output:
	//test: Exchange() -> [status:Invalid Argument [error on PostgreSQL database query call: dbClient is nil]]

}
