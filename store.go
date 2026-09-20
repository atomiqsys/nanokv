package nanokv

type Store interface {
	RowCount() int
	Schema() Schema
	Column(c ColumnID) Column
}

type InMemoryStore struct {
	schema  Schema
	columns []Column
}

func NewInMemoryStore(schema Schema, columns []Column) *InMemoryStore {
	return &InMemoryStore{
		schema:  schema,
		columns: columns,
	}
}

func (s *InMemoryStore) RowCount() int {
	if len(s.columns) == 0 {
		return 0
	}

	return s.columns[0].Len()
}

func (s *InMemoryStore) Schema() Schema {
	return s.schema
}

func (s *InMemoryStore) Column(c ColumnID) Column {
	return s.columns[int(c)]
}
