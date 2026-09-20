package nanokv

type Value struct {
	Kind ColumnKind
	Num  uint64
	Str  string
}
