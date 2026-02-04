# *__signals__* 🪢
Godot-inspired signals for any Go project

# How to use
Initialize, connect, emit, then disconnect.

*__signals__* 🪢 allows for one-to-one, many-to-many, fan-in, and fan-out style communication between nodes[^1].

### Note
I recommend using the [Event Bus](https://en.wikipedia.org/wiki/Observer_pattern) design pattern with *__signals__* 🪢. This design pattern enables "signal up, call down" paradigm, which is popular in Godot as it allows child nodes to easily communicate with their parent nodes. Alternatively, you can use the Signal Bus design pattern for communication between completely non-connected nodes.

[^1]: Nodes are any piece of code that emits or connects to a signal. *__signals__* 🪢 has no built-in notion of nodes... yet

# Future Goals
I am primarily developing *__signals__* 🪢 for use in my own ongoing game development project. If I need more features, I will add them.

# Examples
## [Basic Example](https://github.com/dhdorr/signals/blob/main/examples/ex2/main.go)

```go
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

	err := example_signal.Disconnect(rId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No more signals")

	example_signal.Emit("Dude, where's my signal?")
}

func foo(data any) {
	fmt.Printf("Data: %T\t%v\n", data, data)
}
```
## [Game Loop Example](https://github.com/dhdorr/signals/blob/main/examples/ex3/main.go)
```go
package main

import (
	"fmt"

	"github.com/dhdorr/signals/signals"
)

type Player struct{}
type Interactable struct {
	id    int
	value int
}

var interact_signal = signals.NewSignal("interact_signal")

func main() {

	player := &Player{}
	interactable := NewInteractable()

	// Game loop
	for i := range 10 {
		player.Update(i)
	}

	fmt.Printf("dispose of %T\n", interactable)
}

func (player *Player) Update(i int) {
	// let's pretend the player is picking up an item
	if i == 4 {
		interact_signal.Emit()
	}
}

func (object *Interactable) foo(data any) {
	fmt.Printf("you received %d gold!\n", object.value)
	err := interact_signal.Disconnect(object.id)
	if err != nil {
		fmt.Println(err)
	}
}

func NewInteractable() *Interactable {
	// Generate your own unique receiver id
	rId := 1
	object := &Interactable{id: rId, value: 10}
	interact_signal.Connect(rId, object.foo)
	return object
}
```

# Thank You ❤️