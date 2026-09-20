package nanokv

import (
	"reflect"
	"testing"
)

func TestMaterialize(t *testing.T) {
	s := productsFixture()
	ids := s.Column(0).(*NumericColumn[uint64])
	prices := s.Column(1).(*NumericColumn[uint64])

	// full pipeline: "id > 1 and id < 5 and price < 300", then fetch id + name
	//   range+filter → positions [1, 2] (ids 2,3 with prices 250,90 both < 300)
	positions := ScanRangeFilter(ids, 1, 5, prices, 300)

	rows := Materialize(s, positions, []ColumnID{0, 2}) // want id and name only

	want := [][]Value{
		{{Kind: KindUint64, Num: 2}, {Kind: KindString, Str: "notebook"}},
		{{Kind: KindUint64, Num: 3}, {Kind: KindString, Str: "eraser"}},
	}
	if !reflect.DeepEqual(rows, want) {
		t.Errorf("Materialize = %v, want %v", rows, want)
	}
}

func TestMaterializeEmpty(t *testing.T) {
	s := productsFixture()

	// no surviving positions → no rows, not a crash
	rows := Materialize(s, []RowPosition{}, []ColumnID{0, 1})
	if len(rows) != 0 {
		t.Errorf("Materialize(empty) = %v, want empty", rows)
	}
}
