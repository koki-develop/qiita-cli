package printers

import (
	"fmt"
	"io"
	"sort"
)

type Printer interface {
	Print(w io.Writer, p Printable) error
	SetColumns(cols []string)
}

type Printable interface {
	TableRows() []map[string]interface{}
}

var Registry = map[string]Printer{}

func Register(format string, p Printer) Printer {
	Registry[format] = p
	return p
}

func Get(format string) (Printer, error) {
	p, ok := Registry[format]
	if !ok {
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	return p, nil
}

func ListFormats() []string {
	formats := make([]string, len(Registry))

	i := 0
	for k := range Registry {
		formats[i] = k
		i++
	}

	sort.SliceStable(formats, func(i, j int) bool { return formats[i] < formats[j] })

	return formats
}
