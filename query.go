package nanokv

func FindByPK(s Store, pk uint64) (RowPosition, bool) {
	col := s.Column(s.Schema().PrimaryKey).(*NumericColumn[uint64])

	for i, v := range col.Data {
		if v == pk {
			return RowPosition(i), true
		}
	}
	return 0, false
}

func ScanRange[T Numeric](col *NumericColumn[T], min, max T) []RowPosition {
	var out []RowPosition

	for i, v := range col.Data {
		if v > min && v < max {
			out = append(out, RowPosition(i))
		}
	}

	return out
}

func ScanRangeFilter[K, F Numeric](
	rangeCol *NumericColumn[K], min, max K,
	filterCol *NumericColumn[F], filterMax F,
) []RowPosition {
	candidates := ScanRange(rangeCol, min, max)

	survivors := candidates[:0] // reuse candidate sbacking array and only shirking
	for _, p := range candidates {
		if filterCol.Data[p] < filterMax {
			survivors = append(survivors, p)
		}
	}
	return survivors
}
