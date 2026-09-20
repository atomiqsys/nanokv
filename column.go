package nanokv

// Column is the contract the store exposes per column
type Column interface {
	Kind() ColumnKind
	Len() int
	ValueAt(p RowPosition) Value
}

// Numeric is the set of fixed width types a NumericColumn can hold
type Numeric interface {
	uint8 | uint64
}

// NumericColumn is a array of fixed-width values
type NumericColumn[T Numeric] struct {
	Data []T
}

func (c *NumericColumn[T]) Len() int {
	return len(c.Data)
}

func (c *NumericColumn[T]) Kind() ColumnKind {
	return kindOf[T]()
}

func (c *NumericColumn[T]) ValueAt(p RowPosition) Value {
	return Value{Kind: c.Kind(), Num: uint64(c.Data[p])}
}

// kindOf bridges compile-time type params to the runtime ColumnKind enum.
func kindOf[T Numeric]() ColumnKind {
	var zero T
	switch any(zero).(type) {
	case uint8:
		return KindUint8
	case uint64:
		return KindUint64
	default:
		panic("unsupported numeric kind")
	}
}

type StringColumn struct {
	Buf     []byte
	Offsets []uint32
}

func (c *StringColumn) Kind() ColumnKind {
	return KindString
}

func (c *StringColumn) Len() int {
	return len(c.Offsets) - 1
}

func (c *StringColumn) Get(p RowPosition) string {
	return string(c.Buf[c.Offsets[p]:c.Offsets[p+1]])
}

func (c *StringColumn) ValueAt(p RowPosition) Value {
	return Value{
		Kind: KindString,
		Str:  c.Get(p),
	}
}

func NewStringColumn(vals []string) *StringColumn {
	offsets := make([]uint32, len(vals)+1)
	var buf []byte
	for i, v := range vals {
		buf = append(buf, v...)
		offsets[i+1] = uint32(len(buf))
	}
	return &StringColumn{Buf: buf, Offsets: offsets}
}
