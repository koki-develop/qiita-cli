package qiita

import "net/http"

type ListItemsParameters struct {
	Page    *int    `url:"page,omitempty"`
	PerPage *int    `url:"per_page,omitempty"`
	Query   *string `url:"query,omitempty"`
}

func (cl *Client) ListItems(p *ListItemsParameters) (Items, error) {
	req, err := cl.newRequest(http.MethodGet, "items", p, nil)
	if err != nil {
		return nil, err
	}

	var items Items
	if err := cl.doRequest(req, &items); err != nil {
		return nil, err
	}

	return items, nil
}
