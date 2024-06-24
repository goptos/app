package Li

import (
	"fmt"

	"github.com/goptos/runtime"
	"github.com/goptos/system"
)

type Scope = runtime.Scope
type Elem = system.Elem
type SomeRecordType struct {
	Id   int
	Name string
}

func View(cx *Scope, item SomeRecordType) *Elem {

	/* macro:generated:view:start */
	var view *Elem = (*Elem).New(nil, "li").DynText(cx, func() string { return fmt.Sprintf("%v", item.Name) })
	/* macro:generated:view:end */

	return view
}

/* View
<li>{item.Name}</li>
*/
