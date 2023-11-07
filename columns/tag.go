package loggobold

type TagColumn struct {
	Tag string
}
type TagColumnPlugin struct{}

func (p TagColumnPlugin) ShortArg() string {
	return "t"
}

func (p TagColumnPlugin) ConsumeArgs(args []string) ([]string, Column, error) {
	tagString := args[0]
	remainingArgs := args[1:]

	return remainingArgs, TagColumn{Tag: tagString}, nil
}

func (c TagColumn) Contents(lineNumber int) string {
	return c.Tag
}

func (c TagColumn) MaxWidth() int {
	return len(c.Tag)
}

func init() {
	Register(TagColumnPlugin{})
}
