# optnull

**optnull** is a Go package that provides types that can represent optional,
nullable JSON values.

The types are designed to be used with encoding/json Marshal. They can
represent JSON values that may be empty, null, or have a value. The types can
be used to unmarshal JSON values.

Due to constraints in the encoding/json package, the types are not marshalable
back to the same JSON format. For marshalling, use the type's Value field
directly with the omitempty tag.

## Example

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

## Motivation

The idea to create this package was to be used in PATCH HTTP requests. Imagine
you have a `/users/:userID` endpoint which accepts the PATCH method. A user is
represented with a JSON such as
`{"id": "123", "name": "Alice", "age": 23, "birth": "2001-06-21T00:00:00Z"}`.
Fields other than id are optional. A PATCH request may, for each optional
field, (a) modify its value, (b) unset its value, or (c) do nothing. For
example, a PATCH request to `/users/123` with body `{"name": "John"}` sets the
name to "John", a body of `{"name": null}` sets the name to null and
`{"age": 33}` does not modify the name.

Using native types and the encoding/json package it is not possible to
distinguish in a Go backend service between the omitted and null values. This
package offers a way to do that.

## Implementation details

### Detailed problem

Usually, with encoding/json one uses pointers to distinguish omitted values and
nulls from a zero value.

```go
var j = []byte(`{"a": "string", "b": "", "c": null}`)

type S struct {
    A *string `json:"a"`
    B *string `json:"b"`
    C *string `json:"c"`
    D *string `json:"d"`
}
var s S
_ = json.Unmarshal(j, &s)
fmt.Printf("A: %q, B: %q, C %v, D: %v\n", *s.A, *s.B, s.C, s.D) // A: "string", B: "", C <nil>, D: <nil>
```

However this does not allow us to distinguish omitted from null. This is
usually not a problem, but in some occasions you may want to. One idea may be
to use double pointers, but it soes not work:

```go
var j = []byte(`{"a": "string", "b": "", "c": null}`)

type S struct {
    A *string `json:"a"`
    B *string `json:"b"`
    C *string `json:"c"`
    D *string `json:"d"`
}
var s S
_ = json.Unmarshal(j, &s)
fmt.Printf("A: %q, B: %q, C %v, D: %v\n", *s.A, *s.B, s.C, s.D) // A: "string", B: "", C <nil>, D: <nil>
```

### Sample implementation

The only way to change how a type value is decoded from JSON is to implement
the [json.Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler) interface.
The simplest implementation that works for a string type is:

```go
package optnull

type String struct {
	Value **string
}

func (s *String) UnmarshalJSON(b []byte) error {
	s.Value = new(*string)
	return json.Unmarshal(b, s.Value)
}
```

The previous example using this implementation would be written as:

```go
var j = []byte(`{"a": "string", "b": "", "c": null}`)

type S struct {
    A optnull.String `json:"a"`
    B optnull.String `json:"b"`
    C optnull.String `json:"c"`
    D optnull.String `json:"d"`
}
var s S
_ = json.Unmarshal(j, &s)
fmt.Printf("A: %q, B: %q, C %v, D: %v\n", **s.A.Value, **s.B.Value, *s.C.Value, s.D.Value) // A: "string", B: "", C <nil>, D: <nil>
```

- When the field is omitted, UnmarshalJSON is never called. Therefore the
  final value is the types zero value: `optnull.String{Value: nil}`.
- When the value is `null`, UnmarshalJSON is called, s.Value is assigned to
  `new(*string)` and then json.Unmarshal does not modify the value any more.
  The final value is: `optnull.String{Value: &nil}`.
- When the value is any string, UnmarshalJSON is called, s.Value is assigned to
  `new(*string)` and then json.UnmarshalJSON fills *s.Value with a pointer to
  the string. final value: `optnull.String{Value: &&"string"}`.

### Why not generics?

The package exposes many similar types (`String`, `Time`, `Int`, etc.). I tried
implementing this with a generic type. However, since the solution requires us
to imlement the json.Unmarshaler interface and it's not possible ot implement
interfaces for generic types, this solution is not viable:

```go
type OptNull[T any] struct{
    Value **T
}

// this is not allowed
func (o *OptNull[T any]) UnmarshalJSON(b []byte) error {
	o.Value = new(*OptNull[T])
	return json.Unmarshal(b, o.Value)
}
```
