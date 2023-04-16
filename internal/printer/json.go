package printer

import (
	"encoding/json"
	"io"
)

var (
	JSON = Register("json", NewJSONPrinter())
)

type JSONPrinter struct{}

func NewJSONPrinter() *JSONPrinter {
	return &JSONPrinter{}
}

func (*JSONPrinter) Print(w io.Writer, p Printable) error {
	j, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	if _, err := w.Write(j); err != nil {
		return err
	}
	if _, err := w.Write([]byte{'\n'}); err != nil {
		return err
	}
	return nil
}
