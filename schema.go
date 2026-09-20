package nanokv

// ColumnDef describes one column. The `name` is for the parser to resolve the column.
type ColumnDef struct {
	Name string
	Kind ColumnKind
}

// Schema is the runtime defination of a table.
type Schema struct {
	Cols       []ColumnDef
	PrimaryKey ColumnID // will implented clustered index on this later.
}

func (s Schema) ColumnByName(name string) (ColumnID, bool) {
	for i, c := range s.Cols {
		if c.Name == name {
			return ColumnID(i), true
		}
	}

	return 0, false
}
