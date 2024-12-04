package event1

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/jsonx"
	"github.com/behavioral-ai/log/testrsc"
	"net/http"
)

func ExamplePut() {
	entries, _ := jsonx.New[[]Entry](testrsc.LOG1IngressEntry, nil)

	h := make(http.Header)
	h.Add(core.XExchangeStatus, jsonx.StatusTimeoutUri)
	//ex := core.NewExchangeOverride("", "", json.StatusTimeoutUri)
	//ctx := core.NewExchangeOverrideContext(nil, ex)

	_, status := put[core.Output, Entry](context.Background(), h, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
