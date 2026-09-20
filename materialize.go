package nanokv

func Materialize(s Store, positions []RowPosition, cols []ColumnID) [][]Value {
	chosen := make([]Column, len(cols))

	// get the columns caller asks for using columnID
	for j, c := range cols {
		chosen[j] = s.Column(c)
	}

	rows := make([][]Value, len(positions))
	for i, p := range positions {
		row := make([]Value, len(chosen))
		for j, col := range chosen {
			row[j] = col.ValueAt(p)
		}
		rows[i] = row
	}
	return rows
}
