package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalTime(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Time
	}{
		{`{}`, optnull.NewTime(optnull.Empty, time.Time{})},
		{`{"x": null}`, optnull.NewTime(optnull.Null, time.Time{})},
		{`{"x": "2022-02-02T02:02:02Z"}`, optnull.NewTime(optnull.HasValue, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))},
	}

	for _, tt := range tests {
		var s struct {
			X optnull.Time `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}

func BenchmarkUnmarshalOptNullTime(b *testing.B) {
	tests := []struct {
		input string
		want  optnull.Time
	}{
		{`{}`, optnull.NewTime(optnull.Empty, time.Time{})},
		{`{"x": null}`, optnull.NewTime(optnull.Null, time.Time{})},
		{`{"x": "2022-02-02T02:02:02Z"}`, optnull.NewTime(optnull.HasValue, time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X optnull.Time `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}

func BenchmarkUnmarshalTimePointer(b *testing.B) {
	tests := []struct {
		input string
		want  *time.Time
	}{
		{`{}`, nil},
		{`{"x": null}`, nil},
		{`{"x": "2022-02-02T02:02:02Z"}`, ptr(time.Date(2022, 2, 2, 2, 2, 2, 0, time.UTC))},
	}

	for _, tt := range tests {
		b.Run(tt.input, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				var s struct {
					X *time.Time `json:"x"`
				}
				_ = json.Unmarshal([]byte(tt.input), &s)
			}
		})
	}
}
