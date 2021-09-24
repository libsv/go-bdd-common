package diff

import (
	"fmt"
	"reflect"
)

// ErrUnsupported is returned when an unsupported type is encountered (func, struct ...).
type ErrUnsupported struct {
	LHS reflect.Type
	RHS reflect.Type
}

func (e ErrUnsupported) Error() string {
	return "unsupported types: " + e.LHS.String() + ", " + e.RHS.String()
}

type errInvalidType struct {
	Value interface{}
	For   string
}

func (e errInvalidType) Error() string {
	return fmt.Sprintf("%T is not a valid type for %s", e.Value, e.For)
}

// ErrLHSNotSupported is returned when calling diff.LHS on a Differ that does not contain
// an LHS value (i.e Ignore)
type ErrLHSNotSupported struct {
	Diff Differ
}

func (e ErrLHSNotSupported) Error() string {
	return fmt.Sprintf("%T does not contain an LHS value", e.Diff)
}

// ErrRHSNotSupported is returned when calling diff.EHS on a Differ that does not contain
// an RHS value (i.e Ignore)
type ErrRHSNotSupported struct {
	Diff Differ
}

func (e ErrRHSNotSupported) Error() string {
	return fmt.Sprintf("%T does not contain an RHS value", e.Diff)
}

type errInvalidStream struct {
	Value interface{}
}

func (e errInvalidStream) Error() string {
	return fmt.Sprintf("%T does not implement the Stream interface", e.Value)
}
