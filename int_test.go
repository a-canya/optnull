package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalInt(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Int
	}{
		{`{}`, optnull.NewInt(optnull.Empty, 0)},
		{`{"x": null}`, optnull.NewInt(optnull.Null, 0)},
		{`{"x": 1}`, optnull.NewInt(optnull.HasValue, 1)},
		{`{"x": -42}`, optnull.NewInt(optnull.HasValue, -42)},
	}
	for _, tt := range tests {
		var s struct {
			X optnull.Int `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullInt(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.Int
	}{
		{`{}`, optnull.NewInt(optnull.Empty, 0)},
		{`{"x": null}`, optnull.NewInt(optnull.Null, 0)},
		{`{"x": 1}`, optnull.NewInt(optnull.HasValue, 1)},
		{`{"x": -42}`, optnull.NewInt(optnull.HasValue, -42)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.Int `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalIntPointer(b *testing.B) {
	tests := []struct {
		input string
		want  *int
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": 1}`, ptr(1)},
		{`{"x": -42}`, ptr(-42)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *int `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
