# optnull

optnull is a Go package that provides types that can represent optional,
nullable JSON values.

The types are designed to be used with encoding/json Marshal. They can
represent JSON values that may be empty, null, or have a value. The types
can be used to unmarshal JSON values.

Due to constraints in the encoding/json package, the types are not
marshalable back to the same JSON format. For marshalling, use the type's
Value field directly with the omitempty tag.

Example:

```go
type User struct {
    ID    string `json:"id"`
    Name  optnull.String `json:"name,omitempty"`
    Age   optnull.Int `json:"age,omitempty"`
    Birth optnull.Time `json:"birth,omitempty"`
}

var u User
_ = json.Unmarshal([]byte(`{"id":"123","name":"Alice","age":null}`), &u)
fmt.Println(u.Name.HasValue(), u.Age.Null(), u.Birth.Empty()) // true true true
fmt.Println(u.Name.String()) // Alice
```
