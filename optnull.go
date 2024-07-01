// optnull package provides the [OptNull] type that can represent optional,
// nullable JSON values.
//
// [OptNull] is designed to be used with encoding/json Marshal. It can
// represent JSON values that may be omitted, null, or have a value.
//
// Due to constraints in the encoding/json package, [OptNull] is not
// marshalable back to the same JSON format. For marshalling, use the
// [DoublePointer] method with the omitempty tag.
//
// Example:
//
//	type User struct {
//	    ID    string             `json:"id"`
//	    Name  optnull[string]    `json:"name,omitempty"`
//	    Age   optnull[int]       `json:"age,omitempty"`
//	    Birth optnull[time.Time] `json:"birth,omitempty"`
//	}
//
//	var u User
//	_ = json.Unmarshal([]byte(`{"id":"123","name":"Alice","age":null}`), &u)
//	fmt.Println(u.Name.HasValue, u.Age.IsNull(), u.Birth.IsOmitted()) // true true true
//	fmt.Println(u.Name.String()) // Alice
package optnull

import (
	"encoding/json"
)

// OptNull represents an optional, nullable value especially designed to be
// unmarshalled from JSON. Its possible values are:
//
//	JSON Value | IsPresent | HasValue | Value
//	-----------|-----------|----------|------
//	(omitted)  | false     | false    | nil
//	null       | true      | false    | nil
//	value      | true      | true     | value
//
// If marshalled to JSON, both null and omitted values are marshalled as null
// (due to limitations in encoding/json). For a symmetric marshalling, use the
// [DoublePointer] method with the omitempty tag.
type OptNull[T any] struct {
	IsPresent bool // IsPresent indicates that the value is not omitted.
	HasValue  bool // HasValue indicates that the value is not omitted nor null.
	Value     T    // if HasValue is true, Value contains the value.
}

// Omitted returns an OptNull that represents an omitted value.
func Omitted[T any]() OptNull[T] {
	return OptNull[T]{}
}

// Null returns an OptNull that represents a null value.
func Null[T any]() OptNull[T] {
	return OptNull[T]{IsPresent: true}
}

// WithValue returns an OptNull that represents a value.
func WithValue[T any](v T) OptNull[T] {
	return OptNull[T]{IsPresent: true, HasValue: true, Value: v}
}

func (o *OptNull[T]) UnmarshalJSON(b []byte) error {
	o.IsPresent = true
	if string(b) == "null" {
		return nil
	}
	o.HasValue = true
	return json.Unmarshal(b, &o.Value)
}

func (o OptNull[T]) MarshalJSON() ([]byte, error) {
	if !o.HasValue {
		return []byte("null"), nil
	}
	return json.Marshal(o.Value)
}

// IsOmitted returns true if the value is omitted.
func (o OptNull[T]) IsOmitted() bool {
	return !o.IsPresent
}

// IsNull returns true if the value is null.
func (o OptNull[T]) IsNull() bool {
	return o.IsPresent && !o.HasValue
}

// Pointer returns a pointer to the value, or nil if the value is null or
// omitted.
func (o OptNull[T]) Pointer() *T {
	if !o.HasValue {
		return nil
	}
	return &o.Value
}

// DoublePointer returns a pointer to a pointer to the value, or nil if the
// omitted, or a pointer to nil if the value is null.
//
// This is useful to marshal the value with the omitempty tag:
//
//	type S struct {
//		X **string `json:"x,omitempty"`
//		Y **string `json:"y,omitempty"`
//		Z **string `json:"z,omitempty"`
//	}
//	s := S{
//		X: optnull.WithValue("value").DoublePointer(),
//		Y: optnull.Null[string]().DoublePointer(),
//		Z: optnull.Omitted[string]().DoublePointer(),
//	}
//	b, _ := json.Marshal(s)
//	fmt.Println(string(b)) // {"x":"value","y":null}
func (o OptNull[T]) DoublePointer() **T {
	if !o.IsPresent {
		return nil
	}

	var p *T
	if !o.HasValue {
		return &p
	}

	p = &o.Value
	return &p
}
