package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalString(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.String
	}{
		{`{}`, optnull.NewString(optnull.Empty, "")},
		{`{"x": null}`, optnull.NewString(optnull.Null, "")},
		{`{"x": ""}`, optnull.NewString(optnull.HasValue, "")},
		{`{"x": "string"}`, optnull.NewString(optnull.HasValue, "string")},
	}

	var s struct {
		X optnull.String `json:"x"`
	}

	for _, tt := range tests {
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullString(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.String
	}{
		{`{}`, optnull.NewString(optnull.Empty, "")},
		{`{"x": null}`, optnull.NewString(optnull.Null, "")},
		{`{"x": ""}`, optnull.NewString(optnull.HasValue, "")},
		{`{"x": "string"}`, optnull.NewString(optnull.HasValue, "string")},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.String `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalStringPointer(b *testing.B) {
	tests := []struct {
		input string
		want  *string
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": ""}`, ptr("")},
		{`{"x": "string"}`, ptr("string")},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *string `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
