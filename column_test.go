package nanokv

import "testing"

func TestNumericColumn(t *testing.T) {
	c := &NumericColumn[uint64]{Data: []uint64{10, 20, 30}}
	if c.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", c.Len())
	}
	if c.Kind() != KindUint64 {
		t.Errorf("Kind() = %d, want KindUint64", c.Kind())
	}
}

func TestStringColumn(t *testing.T) {
	c := NewStringColumn([]string{"Ada", "Bob", "Eve"})
	if c.Len() != 3 {
		t.Fatalf("Len() = %d, want 3", c.Len())
	}
	for i, w := range []string{"Ada", "Bob", "Eve"} {
		if got := c.Get(RowPosition(i)); got != w {
			t.Errorf("Get(%d) = %q, want %q", i, got, w)
		}
	}
}
