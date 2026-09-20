package nanokv

import (
	"reflect"
	"testing"
)

func TestFindByPK(t *testing.T) {
	s := productsFixture()

	// id 3 is the third row, so it sits at position 2
	pos, ok := FindByPK(s, 3)
	if !ok || pos != 2 {
		t.Errorf("FindByPK(3) = (%d, %v), want (2, true)", pos, ok)
	}

	// check the very first and very last rows to catch off-by-one mistakes
	if pos, ok := FindByPK(s, 1); !ok || pos != 0 {
		t.Errorf("FindByPK(1) = (%d, %v), want (0, true)", pos, ok)
	}
	if pos, ok := FindByPK(s, 5); !ok || pos != 4 {
		t.Errorf("FindByPK(5) = (%d, %v), want (4, true)", pos, ok)
	}

	// a key that doesn't exist should say "not found", not return a random position
	if _, ok := FindByPK(s, 999); ok {
		t.Errorf("FindByPK(999) ok = true, want false (missing key)")
	}
}

func TestScanRange(t *testing.T) {
	s := productsFixture()
	ids := s.Column(0).(*NumericColumn[uint64])

	// id > 1 and id < 4  →  ids 2 and 3, which are at positions 1 and 2
	got := ScanRange(ids, 1, 4)
	if want := []RowPosition{1, 2}; !reflect.DeepEqual(got, want) {
		t.Errorf("ScanRange(1,4) = %v, want %v", got, want)
	}

	// bounds are exclusive, so ids 1 and 5 should NOT show up
	got = ScanRange(ids, 1, 5)
	if want := []RowPosition{1, 2, 3}; !reflect.DeepEqual(got, want) { // ids 2,3,4
		t.Errorf("ScanRange(1,5) exclusive = %v, want %v", got, want)
	}

	// 0 and 6 are outside the data, so every row falls inside this range
	got = ScanRange(ids, 0, 6)
	if want := []RowPosition{0, 1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("ScanRange(0,6) = %v, want %v", got, want)
	}

	// nothing sits between 10 and 20, so we expect an empty result
	if got := ScanRange(ids, 10, 20); len(got) != 0 {
		t.Errorf("ScanRange(10,20) = %v, want empty", got)
	}

	// a backwards range (min bigger than max) should also return nothing
	if got := ScanRange(ids, 4, 2); len(got) != 0 {
		t.Errorf("ScanRange(4,2) = %v, want empty", got)
	}
}

func TestScanRangeThenFilter(t *testing.T) {
	s := productsFixture()
	ids := s.Column(0).(*NumericColumn[uint64])
	prices := s.Column(1).(*NumericColumn[uint64])

	// id > 1 and id < 5 and price < 200
	//   range step  → positions 1,2,3 (ids 2,3,4 with prices 250,90,400)
	//   filter step → only position 2 stays (price 90 is under 200)
	got := ScanRangeFilter(ids, 1, 5, prices, 200)
	if want := []RowPosition{2}; !reflect.DeepEqual(got, want) {
		t.Errorf("range+filter = %v, want %v", got, want)
	}

	// no price is under 50, so the filter removes everything
	if got := ScanRangeFilter(ids, 0, 6, prices, 50); len(got) != 0 {
		t.Errorf("range+filter (price<50) = %v, want empty", got)
	}

	// every price is under 1000, so the filter keeps every row
	got = ScanRangeFilter(ids, 0, 6, prices, 1000)
	if want := []RowPosition{0, 1, 2, 3, 4}; !reflect.DeepEqual(got, want) {
		t.Errorf("range+filter (price<1000) = %v, want %v", got, want)
	}
}
