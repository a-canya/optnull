package optnull

import "encoding/json"

// Any is an optional, nullable any. Some of its possible values are:
//
//	JSON value | Go value
//	-----------|----------------------
//	(empty)    | Any{Value: nil}
//	null       | Any{Value: &nil}
//	"string"   | Any{Value: &"string"}
//
// It may be used to unmarshal JSON values. For marshalling, use *NullAny with omitempty.
type Any struct {
	Value *any
}

func (a *Any) UnmarshalJSON(b []byte) error {
	a.Value = new(any)
	return json.Unmarshal(b, a.Value)
}
func (a *Any) Empty() bool    { return a.Value == nil }
func (a *Any) Null() bool     { return a.Value != nil && *a.Value == nil }
func (a *Any) HasValue() bool { return a.Value != nil && *a.Value != nil }
func (a *Any) Any() any {
	if a.Value == nil {
		return nil
	}
	return *a.Value
}
