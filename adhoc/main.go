package main

type data struct {
	ID    string
	Type  dataType
	Value any
}

func (d data) String() string {
	return d.ID
}

type dataType uint8

const (
	Invalid dataType = iota
	String
	Bool
	Int
	Float
)

func (d dataType) String() string {
	switch d {
	case String:
		return "String"
	case Bool:
		return "Bool"
	case Int:
		return "Int"
	case Float:
		return "Float"
	default:
		return ""
	}
}

const (
	x = 123
)

func main() {

}
