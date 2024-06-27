package optnull

import "encoding/json"

// Int is an optional, nullable int. Some of its possible values are:
//
//	JSON value | Go value
//	-----------|------------------
//	(empty)    | Int{Value: nil}
//	null       | Int{Value: &nil}
//	123        | Int{Value: &&123}
//
// It may be used to unmarshal JSON values. For marshalling, use *NullInt with omitempty.
type Int struct {
	Value **int
}

func NewInt(v Value, i int) Int {
	switch v {
	case Empty:
		return Int{}
	case Null:
		return Int{Value: new(*int)}
	case HasValue:
		p := &i
		return Int{Value: &p}
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}

func (i *Int) UnmarshalJSON(b []byte) error {
	i.Value = new(*int)
	return json.Unmarshal(b, i.Value)
}
func (i *Int) Empty() bool    { return i.Value == nil }
func (i *Int) Null() bool     { return i.Value != nil && *i.Value == nil }
func (i *Int) HasValue() bool { return i.Value != nil && *i.Value != nil }
func (i *Int) Int() int {
	if i.Value == nil || *i.Value == nil {
		return 0
	}
	return **i.Value
}
func (i *Int) Pointer() *int {
	if i.Value == nil {
		return nil
	}
	return *i.Value
}
