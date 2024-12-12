package dbg

import (
	"fmt"
	"io"
	"os"
)

var dbg = debugger(0)
var noop = &noopWriter{}

type debug struct {
	writer io.Writer
	level  int
}

func (d *debug) println(level int, content ...any) {
	if level >= d.level {
		_, _ = fmt.Fprintln(d.writer, content...)
	}
}

func (d *debug) printf(level int, format string, args ...any) {
	if level >= d.level {
		_, _ = fmt.Fprintf(d.writer, format, args...)
	}
}

func (d *debug) print(level int, content ...any) {
	if level >= d.level {
		_, _ = fmt.Fprint(d.writer, content...)
	}
}

func (d *debug) setLevel(level int) {
	if level > 0 {
		d.writer = os.Stdout
	} else {
		d.writer = noop
	}
	d.level = level
}

func debugger(level int) *debug {
	dbg := debug{}
	dbg.setLevel(level)
	return &dbg
}

type noopWriter struct {
}

func (w *noopWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func Println(level int, content ...any) {
	dbg.println(level, content...)
}

func Printf(level int, format string, args ...any) {
	dbg.printf(level, format, args...)
}

func Print(level int, content ...any) {
	dbg.print(level, content...)
}

func SetLevel(level int) {
	dbg.setLevel(level)
}
