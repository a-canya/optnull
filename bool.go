package optnull

import "encoding/json"

// Bool is an optional, nullable bool. Its possible values are:
//
//	JSON value | Go value
//	-----------|---------------------
//	(empty)    | Bool{Value: nil}
//	null       | Bool{Value: &nil}
//	true       | Bool{Value: &&true}
//	false      | Bool{Value: &&false}
//
// It may be used to unmarshal JSON values. For marshalling, use Value directly:
//
//	Go value (**bool) | JSON value after marshalling (omitempty)
//	------------------|-----------------------------------------
//	nil               | (empty)
//	&nil              | null
//	&&true            | true
//	&&false           | false
type Bool struct {
	Value **bool
}

func NewBool(v Value, b bool) Bool {
	switch v {
	case Empty:
		return Bool{}
	case Null:
		return Bool{Value: new(*bool)}
	case HasValue:
		p := &b
		return Bool{Value: &p}
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}

func (b *Bool) UnmarshalJSON(j []byte) error {
	b.Value = new(*bool)
	return json.Unmarshal(j, b.Value)
}
func (b *Bool) Empty() bool    { return b.Value == nil }
func (b *Bool) Null() bool     { return b.Value != nil && *b.Value == nil }
func (b *Bool) HasValue() bool { return b.Value != nil && *b.Value != nil }
func (b *Bool) Bool() bool     { return b.Value != nil && *b.Value != nil && **b.Value }
func (b *Bool) Pointer() *bool {
	if b.Value == nil {
		return nil
	}
	return *b.Value
}
