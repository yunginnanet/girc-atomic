package girc

import (
	"fmt"
	"sync/atomic"
)

type MarshalableAtomicValue struct {
	*atomic.Value
}

func (m *MarshalableAtomicValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", m.Value.Load())), nil
}

func (m *MarshalableAtomicValue) UnmarshalJSON(b []byte) error {
	m.Value.Store(string(b))
	return nil
}

func (m *MarshalableAtomicValue) String() string {
	return m.Value.Load().(string)
}

func NewAtomicString(s string) *MarshalableAtomicValue {
	obj := &atomic.Value{}
	obj.Store(s)
	return &MarshalableAtomicValue{Value: obj}
}
