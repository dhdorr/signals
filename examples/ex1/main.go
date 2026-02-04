package main

import (
	"fmt"

	"github.com/dhdorr/signals/signals"
)

func main() {
	example_signal := signals.NewSignal("example_signal")

	// Generate your own unique receiver id
	rId := 1
	example_signal.Connect(rId, foo)

	// do stuff...

	example_signal.Emit("Signals are cool!")
}

func foo(data any) {
	fmt.Printf("Data: %T\t%v\n", data, data)
}
