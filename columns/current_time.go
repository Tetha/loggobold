package loggobold

import "time"

type CurrentTimeColumn struct{}

type CurrentTimeColumnPlugin struct{}

func (p CurrentTimeColumnPlugin) ShortArg() string {
	return "c"
}

func (p CurrentTimeColumnPlugin) ConsumeArgs(args []string) ([]string, Column, error) {
	return args, CurrentTimeColumn{}, nil
}

func (c CurrentTimeColumn) Contents(lineNumber int) string {
	return time.Now().Format("03:04:05")
}

func (c CurrentTimeColumn) MaxWidth() int {
	return 8
}

func init() {
	Register(CurrentTimeColumnPlugin{})
}
