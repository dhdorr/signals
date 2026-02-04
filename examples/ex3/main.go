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
