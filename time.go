package optnull

import (
	"encoding/json"
	"time"
)

// Time is an optional, nullable time.Time. Some of its possible values are:
//
//	JSON value             | Go value
//	-----------------------|----------------------
//	(empty)                | Time{Value: nil}
//	null                   | Time{Value: &nil}
//	"2022-02-22T02:02:02Z" | Time{Value: &&time.Time{...}}
//
// It may be used to unmarshal JSON values. For marshalling, use Value directly:
//
//	Go value (**time.Time) | JSON value after marshalling (omitempty)
//	-----------------------|-----------------------------------------
//	nil                    | (empty)
//	&nil                   | null
//	&&time.Time{...}       | "2022-02-22T02:02:02Z"
type Time struct {
	Value **time.Time
}

func NewTime(v Value, t time.Time) Time {
	switch v {
	case Empty:
		return Time{Value: nil}
	case Null:
		return Time{Value: new(*time.Time)}
	case HasValue:
		p := &t
		return Time{Value: &p}
	}
	panic("invalid optnull.Value; must be one of Empty, Null, or HasValue")
}

func (t *Time) UnmarshalJSON(b []byte) error {
	t.Value = new(*time.Time)
	return json.Unmarshal(b, t.Value)
}
func (t *Time) Empty() bool    { return t.Value == nil }
func (t *Time) Null() bool     { return t.Value != nil && *t.Value == nil }
func (t *Time) HasValue() bool { return t.Value != nil && *t.Value != nil }
func (t *Time) Time() time.Time {
	if t.Value == nil || *t.Value == nil {
		return time.Time{}
	}
	return **t.Value
}
func (t *Time) Pointer() *time.Time {
	if t.Value == nil {
		return nil
	}
	return *t.Value
}
