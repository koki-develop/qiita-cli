package printers

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/koki-develop/qiita-cli/internal/util"
)

var (
	Table = Register("table", NewTablePrinter())
)

type TablePrinter struct {
	columns []string
}

func NewTablePrinter() *TablePrinter {
	return &TablePrinter{}
}

// TODO: refactor
func (printer *TablePrinter) Print(w io.Writer, p Printable) error {
	buf := new(bytes.Buffer)

	rows := p.TableRows()

	columnLengths := make(map[string]int, len(printer.columns))
	for _, col := range printer.columns {
		l := 0
		for _, row := range rows {
			v := row[col]
			s := printer.string(v)
			l = util.Max(l, util.Width(s))
		}
		columnLengths[col] = util.Max(l, util.Width(col))
	}

	for i, col := range printer.columns {
		l := columnLengths[col]
		h := util.Pad(col, l)

		buf.WriteString(strings.ToUpper(h))
		if i+1 != len(printer.columns) {
			buf.WriteRune(' ')
		}
	}
	buf.WriteRune('\n')

	for _, row := range rows {
		for i, col := range printer.columns {
			l := columnLengths[col]
			v := row[col]
			d := util.Pad(printer.string(v), l)
			buf.WriteString(d)
			if i+1 != len(printer.columns) {
				buf.WriteRune(' ')
			}
		}
		buf.WriteRune('\n')
	}

	if _, err := io.Copy(w, buf); err != nil {
		return err
	}
	return nil
}

func (printer *TablePrinter) string(v interface{}) string {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice:
		l := rv.Len()
		cols := make([]string, l)
		for i := 0; i < l; i++ {
			cols[i] = printer.string(rv.Index(i).Interface())
		}
		return strings.Join(cols, ", ")
	default:
		return fmt.Sprintf("%v", v)
	}
}

func (printer *TablePrinter) SetColumns(cols []string) {
	printer.columns = cols
}
