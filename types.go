package nanokv

// RowPosition is the physical location of a row within the column arrays
type RowPosition uint32

// ColumnID identifies a column by its offset in the schema.
type ColumnID int

// ColumnKind is the physical type of columns values. It behaves like a runtime type.
type ColumnKind uint8

const (
	KindUint64 ColumnKind = iota
	KindUint8
	KindString
)
