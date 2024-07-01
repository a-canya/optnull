# optnull

**optnull** is a Go package that provides OptNull, a type that can represent
optional, nullable JSON values.

OptNull is designed to be used with encoding/json Marshal. It can represent
JSON values that may be empty, null, or have a value.

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

