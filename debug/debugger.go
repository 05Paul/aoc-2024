package dbg

import (
	"fmt"
	"io"
	"os"
)

var dbg = debugger(false)
var noop = &noopWriter{}

type debug struct {
	writer io.Writer
}

func (d *debug) println(content ...any) {
	_, _ = fmt.Fprintln(d.writer, content...)
}

func (d *debug) printf(format string, args ...any) {
	_, _ = fmt.Fprintf(d.writer, format, args...)
}

func (d *debug) print(content ...any) {
	_, _ = fmt.Fprint(d.writer, content...)
}

func (d *debug) setMode(mode bool) {
	if mode {
		d.writer = os.Stdout
	} else {
		d.writer = noop
	}
}

func debugger(mode bool) *debug {
	dbg := debug{}
	dbg.setMode(mode)
	return &dbg
}

type noopWriter struct {
}

func (w *noopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func Println(content ...any) {
	dbg.println(content...)
}

func Printf(format string, args ...any) {
	dbg.printf(format, args...)
}

func Print(content ...any) {
	dbg.print(content...)
}

func SetMode(active bool) {
	dbg.setMode(active)
}
