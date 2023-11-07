package loggobold

type Column interface {
	Contents(lineNumber int) string
	MaxWidth() int
}

type ColumnPlugin interface {
	ShortArg() string
	ConsumeArgs(args []string) ([]string, Column, error)
}

var Columns []ColumnPlugin

func Register(p ColumnPlugin) {
	Columns = append(Columns, p)
}
