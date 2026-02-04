package signals

import (
	"errors"
	"fmt"
)

type iSignal interface {
	Emit(...any)
	Connect(int, func(any))
	Disconnect(int) error
}

type receiver struct {
	callback func(any)
	id       int
}

type signal struct {
	receivers []*receiver
	name      string
}

func (s *signal) Emit(data ...any) {
	for _, r := range s.receivers {
		r.callback(data)
	}
}

func (s *signal) Connect(id int, fn func(any)) {
	r := &receiver{id: id, callback: fn}
	s.receivers = append(s.receivers, r)
}

func (s *signal) Disconnect(id int) error {
	for i := range s.receivers {
		if s.receivers[i].id == id {
			tmp := s.receivers[i]
			s.receivers[i] = s.receivers[len(s.receivers)-1]
			s.receivers[len(s.receivers)-1] = tmp
			s.receivers = s.receivers[:len(s.receivers)-1]
			return nil
		}
	}

	msg := fmt.Sprintf("Failed to disconnect signal: %s. No receiver with ID: %d found", s.name, id)
	return errors.New(msg)
}

// Call this function when creating a new Signal
func NewSignal(name string) iSignal {
	return iSignal(&signal{name: name, receivers: make([]*receiver, 0, 10)})
}
