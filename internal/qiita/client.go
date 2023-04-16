package qiita

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/google/go-querystring/query"
)

type Client struct {
	token      string
	httpClient *http.Client
}

func New(token string) *Client {
	return &Client{
		token:      token,
		httpClient: new(http.Client),
	}
}

func (cl *Client) newRequest(method, pathname string, params, body interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: "https", Path: "qiita.com/api/v2"}
	u.Path = path.Join(u.Path, pathname)
	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = v.Encode()
	}

	var p io.Reader = nil
	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		p = bytes.NewReader(j)
	}

	req, err := http.NewRequest(method, u.String(), p)
	if err != nil {
		return nil, err
	}
	if cl.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	}
	if p != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (cl *Client) doRequest(req *http.Request, dst interface{}) error {
	resp, err := cl.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var s strings.Builder
		if _, err := io.Copy(&s, resp.Body); err != nil {
			return err
		}
		return errors.New(s.String())
	}

	if dst != nil {
		if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
			return err
		}
	}

	return nil
}
