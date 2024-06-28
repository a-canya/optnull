package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalBool(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Bool
	}{
		{`{}`, optnull.NewBool(optnull.Empty, false)},
		{`{"x": null}`, optnull.NewBool(optnull.Null, false)},
		{`{"x": true}`, optnull.NewBool(optnull.HasValue, true)},
		{`{"x": false}`, optnull.NewBool(optnull.HasValue, false)},
	}

	for _, tt := range tests {
		var s struct {
			X optnull.Bool `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullBool(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.Bool
	}{
		{`{}`, optnull.NewBool(optnull.Empty, false)},
		{`{"x": null}`, optnull.NewBool(optnull.Null, false)},
		{`{"x": true}`, optnull.NewBool(optnull.HasValue, true)},
		{`{"x": false}`, optnull.NewBool(optnull.HasValue, false)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.Bool `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalBoolPointer(b *testing.B) {
	tests := []struct {
		input string
		want  *bool
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": true}`, ptr(true)},
		{`{"x": false}`, ptr(false)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *bool `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
