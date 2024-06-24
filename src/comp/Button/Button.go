package Button

import (
	"fmt"

	"github.com/goptos/runtime"
	"github.com/goptos/system"
)

type Scope = runtime.Scope
type IntSignal = runtime.Signal[int]
type Elem = system.Elem
type Event = system.Event

func View(cx *Scope, parentCount IntSignal) *Elem {
	var count = (*IntSignal).New(nil, cx, 0)
	var clickF = func(Event) {
		count.Set(count.Get() + 1)
		parentCount.Set(parentCount.Get() + 1)
		fmt.Println("Testing")
	}

	/* macro:generated:view:start */
	var view *Elem = (*Elem).New(nil, "div").Child((*Elem).New(nil, "button").On("click", clickF).Text("Child").Text(" ").Text("Add")).Child((*Elem).New(nil, "p").Text("Child").Text(" ").Text("Count").Text(" ").DynText(cx, func() string { return fmt.Sprintf("%v", count.Get()) }))
	/* macro:generated:view:end */

	return view
}

/* View
<div>
	<button on:click={clickF}>Child Add</button>
	<p>Child Count {count.Get()}</p>
</div>
*/
