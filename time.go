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
// It may be used to unmarshal JSON values.
type Time struct {
	Value **time.Time
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
