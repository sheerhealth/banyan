package pointer

import (
	"testing"
)

func TestOr(t *testing.T) {
	t.Parallel()

	v := 5
	p := &v

	if Or(p, 0) != v {
		t.Errorf("want %d, got %d", v, Or(p, 0))
	}

	if Or(nil, 0) != 0 {
		t.Errorf("want %d, got %d", 0, Or(nil, 0))
	}
}

func TestOrDefault(t *testing.T) {
	t.Parallel()

	want := 5
	p := &want

	if OrDefault(p) != want {
		t.Errorf("want %d, got %d", want, OrDefault(p))
	}

	var n *int64

	if OrDefault(n) != 0 {
		t.Errorf("want %d, got %d", 0, OrDefault(n))
	}
}

func TestToOrNil(t *testing.T) {
	t.Parallel()

	if ToUnlessDefault(5) == nil {
		t.Error("expected non-nil pointer")
	}

	if ToUnlessDefault(0) != nil {
		t.Error("expected nil pointer")
	}

	if ToUnlessDefault("hello") == nil {
		t.Error("expected non-nil pointer")
	}

	if ToUnlessDefault("") != nil {
		t.Error("expected nil pointer")
	}
}

func TestTo(t *testing.T) {
	t.Parallel()

	v := To(5)

	if v == nil {
		t.Fatal("expected non-nil pointer")
	}

	if *v != 5 {
		t.Errorf("want %v, got %v", 5, *v)
	}
}

func TestValuesOrDefaults(t *testing.T) {
	t.Parallel()

	const count = 10

	in := make([]*int, 0, count)

	want := make([]int, 0, count)

	for i := range count {
		in = append(in, To(i))
		want = append(want, i)
	}

	got := ValuesOrDefaults(in)

	if len(got) != count {
		t.Fatalf("want %d values, got %d", count, len(got))
	}

	for i := range count {
		if got[i] != want[i] {
			t.Errorf("index[%d]: want %d, got %d", i, want[i], got[i])
		}
	}
}
