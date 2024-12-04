package event2

import (
	"fmt"
	"github.com/behavioral-ai/core/core"
	"net/url"
)

func ExampleGetEgress_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.LOG2EgressEntryTest, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}

func ExampleGetIngress_Test() {
	values := make(url.Values)
	//ctx := core.NewExchangeOverrideContext(nil, core.NewExchangeOverride("", testrsc.LOG2IngressEntryTest, ""))

	values.Add(core.RegionKey, "*")
	entries, status := get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	values.Set(core.RegionKey, "us-west")
	values.Add(core.SubZoneKey, "dc1")
	entries, status = get[core.Output, Entry](nil, nil, values)
	fmt.Printf("test: get() -> [status:%v] [entries:%v]\n", status, len(entries))

	//Output:
	//test: get() -> [status:OK] [entries:2]
	//test: get() -> [status:OK] [entries:1]

}
