package object

import "fmt"

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ = "ERROR"
)

type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }


type Boolean struct {
	Value bool
}
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type NULL struct {}
func (n *NULL) Inspect() string { return "null" }
func (n *NULL) Type() ObjectType { return NULL_OBJ }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }