package optnull_test

import (
	"encoding/json"
	"optnull"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalAny(t *testing.T) {
	tests := []struct {
		input string
		want  optnull.Any
	}{
		{`{}`, optnull.NewAny(optnull.Empty, nil)},
		{`{"x": null}`, optnull.NewAny(optnull.Null, nil)},
		{`{"x": 1}`, optnull.NewAny(optnull.HasValue, 1.0)},
		{`{"x": "string"}`, optnull.NewAny(optnull.HasValue, "string")},
		{`{"x": {"y": 2}}`, optnull.NewAny(optnull.HasValue, map[string]any{"y": 2.0})},
		{`{"x": [1, 2, 3]}`, optnull.NewAny(optnull.HasValue, []any{1.0, 2.0, 3.0})},
	}

	for _, tt := range tests {
		var s struct {
			X optnull.Any `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}
