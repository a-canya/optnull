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
