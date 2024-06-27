package optnull

import "encoding/json"

// Float64 is an optional, nullable float64. Some of its possible values are:
//
//	JSON value | Go value
//	-----------|----------------------
//	(empty)    | Float64{Value: nil}
//	null       | Float64{Value: &nil}
//	123        | Float64{Value: &&123}
//
// It may be used to unmarshal JSON values. For marshalling, use *NullFloat64 with omitempty.
type Float64 struct {
	Value **float64
}

func (f *Float64) UnmarshalJSON(b []byte) error {
	f.Value = new(*float64)
	return json.Unmarshal(b, f.Value)
}
func (f *Float64) Empty() bool    { return f.Value == nil }
func (f *Float64) Null() bool     { return f.Value != nil && *f.Value == nil }
func (f *Float64) HasValue() bool { return f.Value != nil && *f.Value != nil }
func (f *Float64) Float64() float64 {
	if f.Value == nil || *f.Value == nil {
		return 0
	}
	return **f.Value
}
func (f *Float64) Pointer() *float64 {
	if f.Value == nil {
		return nil
	}
	return *f.Value
}
