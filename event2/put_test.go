package event2

import (
	"context"
	"fmt"
	"github.com/behavioral-ai/core/core"
	"github.com/behavioral-ai/core/jsonx"
	"github.com/behavioral-ai/log/testrsc"
	"net/http"
)

func ExamplePut() {
	entries, _ := jsonx.New[[]Entry](testrsc.LOG2IngressEntry, nil)

	//ex := core.NewExchangeOverride("", "", json.StatusTimeoutUri)
	//ctx := core.NewExchangeOverrideContext(nil, ex)
	h := make(http.Header)
	h.Set(core.XExchangeStatus, jsonx.StatusTimeoutUri)
	_, status := put[core.Output, Entry](context.Background(), h, entries)
	fmt.Printf("test: put() -> [status:%v]\n", status)

	//Output:
	//test: put() -> [status:Timeout]

}
