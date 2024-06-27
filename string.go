package optnull

import "encoding/json"

// String is an optional, nullable string. Some of its possible values are:
//
//	JSON value | Go value after unmarshalling
//	-----------|--------------------------
//	(empty)    | String{Value: nil}
//	null       | String{Value: &nil}
//	"string"   | String{Value: &&"string"}
//
// It may be used to unmarshal JSON values. For marshalling, use Value directly:
//
//	Go value (**string) | JSON value after marshalling (omitempty)
//	--------------------|-----------------------------------------
//	nil                 | (empty)
//	&nil                | null
//	&&"string"          | "string"
type String struct {
	Value **string
}

func NewString(v Value, s string) String {
	switch v {
	case Empty:
		return String{}
	case Null:
		return String{Value: new(*string)}
	case HasValue:
		p := &s
		return String{Value: &p}
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}

func (s *String) UnmarshalJSON(b []byte) error {
	s.Value = new(*string)
	return json.Unmarshal(b, s.Value)
}
func (s *String) Empty() bool    { return s.Value == nil }
func (s *String) Null() bool     { return s.Value != nil && *s.Value == nil }
func (s *String) HasValue() bool { return s.Value != nil && *s.Value != nil }
func (s *String) String() string {
	if s.Value == nil || *s.Value == nil {
		return ""
	}
	return **s.Value
}
func (s *String) Pointer() *string {
	if s.Value == nil {
		return nil
	}
	return *s.Value
}
