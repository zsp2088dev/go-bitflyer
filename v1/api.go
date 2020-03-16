package v1

import (
	"github.com/pkg/errors"
	"net/url"
)

type API struct {
	URL string
}

func NewAPI(c *Client, endpoint string) *API {
	return &API{URL: c.BaseURL + endpoint}
}

func (api *API) ToURL() (*url.URL, error) {
	u, err := url.ParseRequestURI(api.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot parse url: %s", u)
	}
	return u, nil
}
