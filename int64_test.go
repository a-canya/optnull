package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalInt64(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Int64
	}{
		{`{}`, optnull.NewInt64(optnull.Empty, 0)},
		{`{"x": null}`, optnull.NewInt64(optnull.Null, 0)},
		{`{"x": 1}`, optnull.NewInt64(optnull.HasValue, 1)},
		{`{"x": -42}`, optnull.NewInt64(optnull.HasValue, -42)},
	}
	for _, tt := range tests {
		var s struct {
			X optnull.Int64 `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullInt64(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.Int64
	}{
		{`{}`, optnull.NewInt64(optnull.Empty, 0)},
		{`{"x": null}`, optnull.NewInt64(optnull.Null, 0)},
		{`{"x": 1}`, optnull.NewInt64(optnull.HasValue, 1)},
		{`{"x": -42}`, optnull.NewInt64(optnull.HasValue, -42)},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.Int64 `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalInt64Pointer(b *testing.B) {
	tests := []struct {
		input string
		want  *int64
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": 1}`, ptr(int64(1))},
		{`{"x": -42}`, ptr(int64(-42))},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *int64 `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
