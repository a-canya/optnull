package optnull_test

import (
	"encoding/json"
	"fmt"
	"optnull"
	"testing"
)

type Nontrivial struct {
	A int    `json:"a"`
	B string `json:"b"`
}

type IncomingPatch struct {
	X   optnull.OptNull[int]        `json:"x,omitempty"`
	Y   optnull.OptNull[string]     `json:"y,omitempty"`
	Z   optnull.OptNull[bool]       `json:"z,omitempty"`
	NT1 optnull.OptNull[Nontrivial] `json:"nt1,omitempty"`
	NT2 optnull.OptNull[Nontrivial] `json:"nt2,omitempty"`
	NT3 optnull.OptNull[Nontrivial] `json:"nt3,omitempty"`
}
type SymmetricResponse struct {
	X   **int        `json:"x,omitempty"`
	Y   **string     `json:"y,omitempty"`
	Z   **bool       `json:"z,omitempty"`
	NT1 **Nontrivial `json:"nt1,omitempty"`
	NT2 **Nontrivial `json:"nt2,omitempty"`
	NT3 **Nontrivial `json:"nt3,omitempty"`
}

const (
	incomingJSON = `{
		"x": 1,
		"y": null,
		"nt1": {
			"a": 5,
			"b": "hello"
		},
		"nt2": null
	}`
	minimizedJSONWithNulls = `{"x":1,"y":null,"z":null,"nt1":{"a":5,"b":"hello"},"nt2":null,"nt3":null}`
	minimizedJSON          = `{"x":1,"y":null,"nt1":{"a":5,"b":"hello"},"nt2":null}`
)

func TestUnmarshal(t *testing.T) {
	patch := IncomingPatch{}
	err := json.Unmarshal([]byte(incomingJSON), &patch)
	if err != nil {
		t.Fatal(err)
	}
	if !patch.X.HasValue {
		t.Error("X should have a value")
	}
	if !patch.Y.IsNull() {
		t.Error("Y should be null")
	}
	if !patch.Z.IsOmitted() {
		t.Error("Z should be omitted")
	}
	if !patch.NT1.HasValue {
		t.Error("NT1 should have a value")
	}
	if !patch.NT2.IsNull() {
		t.Error("NT2 should be null")
	}
	if !patch.NT3.IsOmitted() {
		t.Error("NT3 should be omitted")
	}

	fmt.Printf("%#v\n", patch)
}

func BenchmarkUnmarshal(b *testing.B) {
	input := []byte(incomingJSON)
	for i := 0; i < b.N; i++ {
		patch := IncomingPatch{}
		_ = json.Unmarshal(input, &patch)
	}
}

func TestMarshalOptNull(t *testing.T) {
	patch := IncomingPatch{
		X:   optnull.WithValue(1),
		Y:   optnull.Null[string](),
		Z:   optnull.Omitted[bool](),
		NT1: optnull.WithValue(Nontrivial{5, "hello"}),
		NT2: optnull.Null[Nontrivial](),
		NT3: optnull.Omitted[Nontrivial](),
	}

	b, err := json.Marshal(patch)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != minimizedJSONWithNulls {
		t.Errorf("got %q, want %q", b, minimizedJSONWithNulls)
	}
	fmt.Println(string(b))
}

func BenchmarkMarshalOptNull(b *testing.B) {
	patch := IncomingPatch{
		X:   optnull.WithValue(1),
		Y:   optnull.Null[string](),
		Z:   optnull.Omitted[bool](),
		NT1: optnull.WithValue(Nontrivial{5, "hello"}),
		NT2: optnull.Null[Nontrivial](),
		NT3: optnull.Omitted[Nontrivial](),
	}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(patch)
	}
}

func TestMarshalSymmetric(t *testing.T) {
	patch := SymmetricResponse{
		X:   optnull.WithValue(1).DoublePointer(),
		Y:   optnull.Null[string]().DoublePointer(),
		Z:   optnull.Omitted[bool]().DoublePointer(),
		NT1: optnull.WithValue(Nontrivial{5, "hello"}).DoublePointer(),
		NT2: optnull.Null[Nontrivial]().DoublePointer(),
		NT3: optnull.Omitted[Nontrivial]().DoublePointer(),
	}

	b, err := json.Marshal(patch)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != minimizedJSON {
		t.Errorf("got %q, want %q", b, minimizedJSON)
	}
	fmt.Println(string(b))
}

func BenchmarkMarshalSymmetric(b *testing.B) {
	patch := SymmetricResponse{
		X:   optnull.WithValue(1).DoublePointer(),
		Y:   optnull.Null[string]().DoublePointer(),
		Z:   optnull.Omitted[bool]().DoublePointer(),
		NT1: optnull.WithValue(Nontrivial{5, "hello"}).DoublePointer(),
		NT2: optnull.Null[Nontrivial]().DoublePointer(),
		NT3: optnull.Omitted[Nontrivial]().DoublePointer(),
	}
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(patch)
	}
}
