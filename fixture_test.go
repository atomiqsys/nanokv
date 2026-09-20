package nanokv

import "testing"

// productsFixture builds a small, hand-written products table used as the oracle
// across query tests. Rows are in insertion order (unsorted) on purpose.
func productsFixture() *InMemoryStore {
	schema := Schema{
		Cols: []ColumnDef{
			{Name: "id", Kind: KindUint64},
			{Name: "price", Kind: KindUint64},
			{Name: "name", Kind: KindString},
		},
		PrimaryKey: 0, // id
	}
	ids := &NumericColumn[uint64]{Data: []uint64{1, 2, 3, 4, 5}}
	prices := &NumericColumn[uint64]{Data: []uint64{100, 250, 90, 400, 150}}
	names := NewStringColumn([]string{"pen", "notebook", "eraser", "backpack", "marker"})
	return NewInMemoryStore(schema, []Column{ids, prices, names})
}

func TestFixtureReadsBack(t *testing.T) {
	s := productsFixture()
	if s.RowCount() != 5 {
		t.Fatalf("RowCount() = %d, want 5", s.RowCount())
	}
	// This type-assert-once pattern is exactly what the executor will do.
	ids := s.Column(0).(*NumericColumn[uint64])
	if ids.Data[0] != 1 {
		t.Errorf("id[0] = %d, want 1", ids.Data[0])
	}
	names := s.Column(2).(*StringColumn)
	if got := names.Get(3); got != "backpack" {
		t.Errorf("name[3] = %q, want backpack", got)
	}
}
