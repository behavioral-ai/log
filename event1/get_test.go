package event1

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
	"net/url"
)

func ExampleGetIngress_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.LOG1IngressEntry, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:4]
	//test: get() -> [status:OK] [entries:2]

}
