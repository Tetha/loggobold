package main

import (
	"bufio"
	"fmt"
	"os"

	loggobold "github.com/tetha/loggobold/columns"
)

func main() {
	fmt.Println("All systems go!")

	outputColumns := make([]loggobold.Column, 0)
	remainingArgs := os.Args[1:]
	for len(remainingArgs) > 0 {
		fmt.Printf("%s\n", remainingArgs)
		found := false
		for _, plugin := range loggobold.Columns {
			if plugin.ShortArg() == remainingArgs[0][1:] {
				found = true
				newRemainingArgs, column, err := plugin.ConsumeArgs(remainingArgs[1:])
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				outputColumns = append(outputColumns, column)
				remainingArgs = newRemainingArgs
			}
		}
		if !found {
			fmt.Printf("Unknown column: %s\n", remainingArgs[0])
			os.Exit(1)
		}
	}

	reader := bufio.NewReader(os.Stdin)
	lineNumber := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			os.Exit(0)
		}

		columnCount := len(outputColumns)
		for i, column := range outputColumns {
			fmt.Printf("%s", column.Contents(lineNumber))
			if i < columnCount {
				fmt.Printf("|")
			} else {
				fmt.Printf("]")
			}
		}
		fmt.Printf("%s", line)
		lineNumber++
	}
}
