package object

import "fmt"

type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

// integer
type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// booleans
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// null
type Null struct{}

func (b *Null) Type() ObjectType {
	return NULL_OBJ
}
func (b *Null) Inspect() string {
	return "null"
}

// return
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType {
	return RETURN_VALUE_OBJ
}
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}
