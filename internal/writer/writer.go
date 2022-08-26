package writer

import (
	"fmt"

	"github.com/gosuri/uilive"
)

type Writer struct {
	writer *uilive.Writer
}

func NewWriter() *Writer {
	return &Writer{
		writer: uilive.New()}
}

func (w *Writer) Print(context string) {
	w.writer.Start()
	w.clear()
	fmt.Fprintf(w.writer, context+"\n")
}

func (w *Writer) Stop() {
	w.writer.Stop()
}

func (w *Writer) Update(context string) {
	fmt.Fprintln(w.writer, context)
}

func (w *Writer) clear() {
	fmt.Print("\033[H\033[2J")
}
