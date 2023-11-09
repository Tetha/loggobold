package loggobold

import "fmt"

type PatternColumn struct {
	Pattern []string
}

type PatternColumnPlugin struct{}

var PATTERNS = map[string][]string{
	"bouncy-ball": {
		"o....",
		".o...",
		"..o..",
		"...o.",
		"....o",
		"...o.",
		"..o..",
		".o...",
		"o....",
	},
	"disco": {
		"\\....",
		".\\...",
		"..\\..",
		"...\\.",
		"....\\",
		"..../",
		".../.",
		"../..",
		"./...",
		"/....",
	},
	"slow-fox": {
		"\\....",
		"\\....",
		"\\....",
		".\\...",
		".\\...",
		".\\...",
		".\\...",
		"..\\..",
		"..\\..",
		"..\\..",
		"...\\.",
		"...\\.",
		"...\\.",
		"....\\",
		"....\\",
		"....\\",
		"..../",
		"..../",
		"..../",
		".../.",
		".../.",
		".../.",
		"../..",
		"../..",
		"../..",
		"./...",
		"./...",
		"./...",
		"/....",
		"/....",
		"/....",
	},
	"bigball": {
		"O...........",
		".O..........",
		"..O.........",
		"...O........",
		"....O.......",
		".....O......",
		"......O.....",
		".......O....",
		"........O...",
		".........O..",
		"..........O.",
		"...........O",
		"..........O.",
		".........O..",
		"........O...",
		".......O....",
		"......O.....",
		".....O......",
		"....O.......",
		"...O........",
		"..O.........",
		".O..........",
		"O...........",
	},
	"flash": {
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"                ",
		"     O          ",
		"      O         ",
		"       O        ",
		"        O       ",
		"         O      ",
		"   O      O     ",
		"    OO     O    ",
		"     O  O   O   ",
		"      O   O OO  ",
		"       O        ",
		"    \\\\  O  //   ",
		"     \\\\   //    ",
		"      \\\\ //     ",
		"       \\V/      ",
		"        V       ",
		"                ",
		"                ",
		"                ",
	},
	"beach": {
		"~....",
		"~~...",
		"~~~..",
		"~~~~.",
		"~~~~~",
		"~~~~~",
		"~~~~~",
		"~~~~.",
		"~~~..",
		"~~...",
	},
	"disco-triangles": {
		".../\\...",
		"..//\\\\..",
		".///\\\\\\.",
		"////\\\\\\\\",
		"\\\\\\\\////",
		".\\\\\\///.",
		"..\\\\//..",
		"...\\/...",
	},
}

func (p PatternColumnPlugin) ShortArg() string {
	return "p"
}

func (p PatternColumnPlugin) ConsumeArgs(args []string) ([]string, Column, error) {
	patternArg := args[0]
	remainingArgs := args[1:]

	pattern, ok := PATTERNS[patternArg]
	if !ok {
		return remainingArgs, nil, fmt.Errorf("Unknown pattern: %s", patternArg)
	}

	return remainingArgs, PatternColumn{Pattern: pattern}, nil
}

func (c PatternColumn) Contents(lineNumber int) string {
	return c.Pattern[lineNumber%len(c.Pattern)]
}

func (c PatternColumn) MaxWidth() int {
	return len(c.Pattern[0])
}

func init() {
	Register(PatternColumnPlugin{})
}
