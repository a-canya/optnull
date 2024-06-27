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

	for _, tt := range tests {
		var s struct {
			X optnull.String `json:"x"`
		}
		err := json.Unmarshal([]byte(tt.input), &s)
		assert.NoError(t, err, "input = %q", tt.input)
		assert.Equal(t, tt.want, s.X, "input = %q", tt.input)
	}
}
