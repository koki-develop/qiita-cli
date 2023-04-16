package printer

import (
	"bytes"
	"fmt"
	"io"

	"github.com/koki-develop/qiita-cli/internal/util"
)

var (
	Table = Register("table", NewTablePrinter())
)

type TablePrinter struct{}

func NewTablePrinter() *TablePrinter {
	return &TablePrinter{}
}

func (*TablePrinter) Print(w io.Writer, cols []string, p Printable) error {
	buf := new(bytes.Buffer)

	rows := p.TableRows()

	columnLengths := make(map[string]int, len(cols))
	for _, col := range cols {
		l := 0
		for _, row := range rows {
			v := row[col]
			s := fmt.Sprint(v)
			l = util.Max(l, util.Width(s))
		}
		columnLengths[col] = util.Max(l, util.Width(col))
	}

	for i, col := range cols {
		l := columnLengths[col]
		h := util.Pad(col, l)

		buf.WriteString(h)
		if i+1 != len(cols) {
			buf.WriteRune(' ')
		}
	}
	buf.WriteRune('\n')

	for _, row := range rows {
		for i, col := range cols {
			l := columnLengths[col]
			v := row[col]
			d := util.Pad(fmt.Sprint(v), l)
			buf.WriteString(d)
			if i+1 != len(cols) {
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
