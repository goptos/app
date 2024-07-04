package App

import (
	"fmt"

	"github.com/goptos/app/comp/Button"
	_ "github.com/goptos/app/comp/Button" /* macro:import */
	"github.com/goptos/app/comp/Li"
	_ "github.com/goptos/app/comp/Li" /* macro:import */
	"github.com/goptos/runtime"
	"github.com/goptos/system"
)

type Scope = runtime.Scope
type SignalInt = runtime.Signal[int]
type Elem = system.Elem
type Event = system.Event

type SomeRecordType = Li.SomeRecordType

func View(cx *Scope) *Elem {
	var count = (*SignalInt).New(nil, cx, 0)
	var double_count = func() int { return count.Get() * 2 }
	var showF = func() bool {
		if count.Get() > 2 {
			return true
		} else {
			return false
		}
	}
	var clickAddF = func(Event) {
		count.Set(count.Get() + 1)
	}
	var clickSubF = func(Event) {
		count.Set(count.Get() - 1)
	}
	var eachCollectF = func() []SomeRecordType {
		var arr = []SomeRecordType{}
		if count.Get() < 0 {
			return arr
		}
		// call to API etc..
		for i, v := range []string{"Border Collie", "German Shepherd", "Golden Retriever"} {
			arr = append(arr, SomeRecordType{Id: i, Name: v})
		}
		if len(arr) < count.Get() {
			return arr
		}
		return arr[0:count.Get()]
	}
	var eachKeyF = func(item SomeRecordType) int {
		return item.Id
	}

	cx.CreateEffect(func() {
		fmt.Printf("::: Dom node #0 value is: %d\n", count.Get())
	})
	cx.CreateEffect(func() {
		fmt.Printf("::: Dom node #1 value is: %d\n", count.Get())
	})
	cx.CreateEffect(func() {
		fmt.Printf("::: Derived signal value is: %d\n", double_count())
	})

	/* macro:generated:view:start */
	var view *Elem = (*Elem).New(nil, "div").Child((*Elem).New(nil, "p").Child((*Elem).New(nil, "b").Text("Goptos"))).Child((*Elem).New(nil, "p").Text("A fine grained reactive web framework in Go")).Child((*Elem).New(nil, "p").Child((*Elem).New(nil, "i").Text("Inspired by Leptos and SolidJS"))).Child((*Elem).New(nil, "button").Attr("type", "button").On("click", clickAddF).Text("Parent Add")).Child((*Elem).New(nil, "button").Attr("type", "button").On("click", clickSubF).Text("Parent Sub")).Child((*Elem).New(nil, "p").Text("Parent Count ").DynText(cx, func() string { return fmt.Sprintf("%v", count.Get()) })).Child((*Elem).New(nil, "div").Child(Button.View(cx, count)).Attr("count", "")).DynChild(cx, showF, (*Elem).New(nil, "div").Child((*Elem).New(nil, "a").Attr("href", "http://localhost:8080").Attr("id", "link-id").Text("Reload"))).Child(system.Each((*Elem).New(nil, "ul"), cx, eachCollectF, eachKeyF, Li.View))
	/* macro:generated:view:end */

	return view
}

/* View
<div>
	<p><b>Goptos</b></p>
	<p>A fine grained reactive web framework in Go</p>
	<p><i>Inspired by Leptos and SolidJS</i></p>
	<button type="button" on:click={clickAddF}>Parent Add</button>
	<button type="button" on:click={clickSubF}>Parent Sub</button>
	<p>Parent Count {count.Get()}</p>
	<div>
		<Button   count  />
	</div>
	<div if={showF}>
		<a href="http://localhost:8080" id="link-id">Reload</a>
	</div>
	<ul each={eachCollectF} key={eachKeyF}>
		<Li />
	</ul>
</div>
*/
