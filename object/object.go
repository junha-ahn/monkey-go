package object

import "fmt"

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ = "NULL"
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


type BOOLEAN struct {
	Value bool
}
func (b *BOOLEAN) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *BOOLEAN) Type() ObjectType { return BOOLEAN_OBJ }

type NULL struct {}
func (n *NULL) Inspect() string { return "null" }
func (n *NULL) Type() ObjectType { return NULL_OBJ }
