package qiita

import "github.com/koki-develop/qiita-cli/internal/printer"

var (
	_ printer.Printable = (Item)(nil)
	_ printer.Printable = (Items)(nil)
)

type Item map[string]interface{}
type Items []Item

func (item Item) TableRows() []map[string]interface{} {
	item["user"] = item["user"].(map[string]interface{})["id"]
	return []map[string]interface{}{item}
}

func (items Items) TableRows() []map[string]interface{} {
	rows := make([]map[string]interface{}, len(items))
	for i, item := range items {
		rows[i] = item.TableRows()[0]
	}
	return rows
}
