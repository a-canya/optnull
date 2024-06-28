package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalFloat64(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Float64
	}{
		{`{}`, optnull.NewFloat64(optnull.Empty, 0.0)},
		{`{"x": null}`, optnull.NewFloat64(optnull.Null, 0.0)},
		{`{"x": 1}`, optnull.NewFloat64(optnull.HasValue, 1.0)},
		{`{"x": 1.1}`, optnull.NewFloat64(optnull.HasValue, 1.1)},
	}
	for _, tt := range tests {
		var s struct {
			X optnull.Float64 `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullFloat64(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.Float64
	}{
		{`{}`, optnull.NewFloat64(optnull.Empty, 0.0)},
		{`{"x": null}`, optnull.NewFloat64(optnull.Null, 0.0)},
		{`{"x": 1}`, optnull.NewFloat64(optnull.HasValue, 1.0)},
		{`{"x": 1.1}`, optnull.NewFloat64(optnull.HasValue, 1.1)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.Float64 `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalFloat64Pointer(b *testing.B) {
	tests := []struct {
		input string
		want  *float64
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": 1}`, ptr(1.0)},
		{`{"x": 1.1}`, ptr(1.1)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *float64 `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
