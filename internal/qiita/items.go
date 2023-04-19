package qiita

import (
	"fmt"
	"net/http"
)

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

type ListAuthenticatedUserItemsParameters struct {
	Page    *int `url:"page,omitempty"`
	PerPage *int `url:"per_page,omitempty"`
}

func (cl *Client) ListAuthenticatedUserItems(p *ListAuthenticatedUserItemsParameters) (Items, error) {
	req, err := cl.newRequest(http.MethodGet, "authenticated_user/items", p, nil)
	if err != nil {
		return nil, err
	}

	var items Items
	if err := cl.doRequest(req, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (cl *Client) GetItem(id string) (*Item, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("items/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := cl.doRequest(req, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

type CreateItemParameters struct {
	Title   *string `json:"title,omitempty"`
	Tags    *Tags   `json:"tags,omitempty"`
	Body    *string `json:"body,omitempty"`
	Private *bool   `json:"private,omitempty"`
	Tweet   *bool   `json:"tweet,omitempty"`
}

func (cl *Client) CreateItem(params *CreateItemParameters) (*Item, error) {
	req, err := cl.newRequest(http.MethodPost, "items", nil, params)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := cl.doRequest(req, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

type UpdateItemParameters struct {
	Title   *string `json:"title,omitempty"`
	Tags    *Tags   `json:"tags,omitempty"`
	Body    *string `json:"body,omitempty"`
	Private *bool   `json:"private,omitempty"`
}

func (cl *Client) UpdateItem(id string, params *UpdateItemParameters) (*Item, error) {
	req, err := cl.newRequest(http.MethodPatch, fmt.Sprintf("items/%s", id), nil, params)
	if err != nil {
		return nil, err
	}

	var item Item
	if err := cl.doRequest(req, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (cl *Client) DeleteItem(id string) error {
	req, err := cl.newRequest(http.MethodDelete, fmt.Sprintf("items/%s", id), nil, nil)
	if err != nil {
		return err
	}

	if err := cl.doRequest(req, nil); err != nil {
		return err
	}

	return nil
}
