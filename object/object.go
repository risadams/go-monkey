package object

import (
	"fmt"
)

type ObjectType string

const (
	BOOLEAN_OBJ = "BOOLEAN"
	INTEGER_OBJ = "INTEGER"
	NULL_OBJ    = "NULL"
)

// the base from which all other objects are derived
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Boolean, represents a true or false value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Integer, represents a discrete 64 bit number
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

// represents an object that is null (abscent of value)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }
