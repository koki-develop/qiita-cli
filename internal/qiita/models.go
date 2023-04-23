package qiita

import (
	"github.com/koki-develop/qiita-cli/internal/printers"
	"github.com/koki-develop/qiita-cli/internal/util"
)

var (
	_ printers.Printable = (Item)(nil)
	_ printers.Printable = (Items)(nil)
)

type ItemFrontMatter struct {
	ID      *string   `yaml:"id,omitempty"`
	Title   *string   `yaml:"title,omitempty"`
	Tags    *[]string `yaml:"tags,flow,omitempty"`
	Private *bool     `yaml:"private,omitempty"`
}

func (fm ItemFrontMatter) QiitaTags() *Tags {
	if fm.Tags == nil {
		return nil
	}

	var tags Tags
	for _, t := range *fm.Tags {
		tags = append(tags, &Tag{Name: t})
	}
	return &tags
}

type Item map[string]interface{}
type Items []Item

func (item Item) ID() string {
	return item["id"].(string)
}

func (item Item) Title() string {
	return item["title"].(string)
}

func (item Item) Tags() Tags {
	tags := item["tags"].([]interface{})
	rtn := make(Tags, len(tags))
	for i, t := range tags {
		rtn[i] = &Tag{
			Name: t.(map[string]interface{})["name"].(string),
		}
	}

	return rtn
}

func (item Item) FrontMatter() ItemFrontMatter {
	return ItemFrontMatter{
		ID:      util.Ptr(item.ID()),
		Title:   util.Ptr(item.Title()),
		Tags:    util.Ptr(item.Tags().Names()),
		Private: util.Ptr(item.Private()),
	}
}

func (item Item) Body() string {
	return item["body"].(string)
}

func (item Item) Private() bool {
	return item["private"].(bool)
}

func (item Item) TableRows() []map[string]interface{} {
	item["user"] = item["user"].(map[string]interface{})["id"]
	item["tags"] = item.Tags().Names()

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
