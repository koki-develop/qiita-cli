package qiita

import "github.com/koki-develop/qiita-cli/internal/printers"

var (
	_ printers.Printable = (Item)(nil)
	_ printers.Printable = (Items)(nil)
)

type Item map[string]interface{}
type Items []Item

func (item Item) TableRows() []map[string]interface{} {
	item["user"] = item["user"].(map[string]interface{})["id"]

	tags := item["tags"].([]interface{})
	tagNames := make([]string, len(tags))
	for i, t := range tags {
		tagNames[i] = t.(map[string]interface{})["name"].(string)
	}
	item["tags"] = tagNames

	return []map[string]interface{}{item}
}

func (items Items) TableRows() []map[string]interface{} {
	rows := make([]map[string]interface{}, len(items))
	for i, item := range items {
		rows[i] = item.TableRows()[0]
	}
	return rows
}

type Tag struct {
	Name string `json:"name"`
}

type Tags []*Tag

func (tags Tags) Names() []string {
	names := make([]string, len(tags))
	for i, t := range tags {
		names[i] = t.Name
	}
	return names
}

func TagsFromStrings(ss []string) Tags {
	tags := make(Tags, len(ss))
	for i, s := range ss {
		tags[i] = &Tag{Name: s}
	}
	return tags
}
