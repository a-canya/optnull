package optnull

import "encoding/json"

// Int64 is an optional, nullable int64. Some of its possible values are:
//
//	JSON value | Go value
//	-----------|--------------------
//	(empty)    | Int64{Value: nil}
//	null       | Int64{Value: &nil}
//	123        | Int64{Value: &&123}
//
// It may be used to unmarshal JSON values. For marshalling, use *NullInt64 with omitempty.
type Int64 struct {
	Value **int64
}

func NewInt64(v Value, i int64) Int64 {
	switch v {
	case Empty:
		return Int64{}
	case Null:
		return Int64{Value: new(*int64)}
	case HasValue:
		p := &i
		return Int64{Value: &p}
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}

func (o *Int64) UnmarshalJSON(b []byte) error {
	o.Value = new(*int64)
	return json.Unmarshal(b, o.Value)
}
func (i *Int64) Empty() bool    { return i.Value == nil }
func (i *Int64) Null() bool     { return i.Value != nil && *i.Value == nil }
func (i *Int64) HasValue() bool { return i.Value != nil && *i.Value != nil }
func (i *Int64) Int64() int64 {
	if i.Value == nil || *i.Value == nil {
		return 0
	}
	return **i.Value
}
func (i *Int64) Pointer() *int64 {
	if i.Value == nil {
		return nil
	}
	return *i.Value
}
