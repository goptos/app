package main

import (
	"github.com/goptos/app/comp/App"
	"github.com/goptos/runtime"
	"github.com/goptos/system"
)

type Scope = runtime.Scope
type IntSignal = runtime.Signal[int]
type Elem = system.Elem
type Event = system.Event

func main() {
	system.Run(App.View)
}

//go:generate goptos genview
