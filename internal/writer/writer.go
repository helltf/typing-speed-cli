package writer

import (
	"fmt"

	"github.com/gosuri/uilive"
)

var writer = uilive.New()

func Print(context string) {
	writer.Start()
	clear()
	fmt.Fprintf(writer, context+"\n")
}

func Stop() {
	writer.Stop()
}

func Update(context string) {
	fmt.Fprintln(writer, context)
}

func clear() {
	fmt.Print("\033[H\033[2J")
}
