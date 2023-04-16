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

func (*JSONPrinter) Print(w io.Writer, cols []string, p Printable) error {
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(p); err != nil {
		return err
	}
	return nil
}
