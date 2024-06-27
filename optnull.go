// optnull package provides types that can represent optional, nullable JSON
// values.
//
// The types are designed to be used with encoding/json Marshal. They can
// represent JSON values that may be empty, null, or have a value. The types
// can be used to unmarshal JSON values.
//
// Due to constraints in the encoding/json package, the types are not
// marshalable back to the same JSON format. For marshalling, use the type's
// Value field directly with the omitempty tag.
//
// Example:
//
//	type User struct {
//		ID    string `json:"id"`
//		Name  optnull.String `json:"name,omitempty"`
//		Age   optnull.Int `json:"age,omitempty"`
//		Birth optnull.Time `json:"birth,omitempty"`
//	}
//
//	var u User
//	_ = json.Unmarshal([]byte(`{"id":"123","name":"Alice","age":null}`), &u)
//	fmt.Println(u.Name.HasValue(), u.Age.Null(), u.Birth.Empty()) // true true true
//	fmt.Println(u.Name.String()) // Alice
package optnull

// Value represents the possible states of an optional, nullable value: Empty,
// Null, or HasValue.
type Value uint8

const (
	Empty    Value = iota // Empty represents an empty value.
	Null                  // Null represents a null value.
	HasValue              // HasValue represents a value which is neither empty nor null.
)

func (v Value) String() string {
	switch v {
	case Empty:
		return "Empty"
	case Null:
		return "Null"
	case HasValue:
		return "HasValue"
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}
