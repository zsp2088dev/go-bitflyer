package httpclient

import (
	"bytes"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	api "github.com/zsp2088dev/go-bitflyer"
	"github.com/zsp2088dev/go-bitflyer/auth"
	"io/ioutil"
	"net/http"
	"time"
)

// A high-performance 100% compatible drop-in replacement of "encoding/json"
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type HttpClient struct {
	auth *auth.Auth
}

func New() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) Auth(auth *auth.Auth) *HttpClient {
	c.auth = auth
	return c
}

func (c *HttpClient) Request(api api.API, req api.Request, result interface{}) (*http.Response, error) {
	u, err := api.ToURL()
	if err != nil {
		return nil, errors.Wrap(err, "set base url")
	}
	u.RawQuery = req.Query()

	r, err := http.NewRequest(req.Method(), u.String(), bytes.NewReader(req.Payload()))
	if err != nil {
		return nil, errors.Wrap(err, "create new request")
	}

	if c.auth != nil {
		r.Header = *auth.CreateHeaders(c.auth, req, u)
	}

	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(r)
	if err != nil {
		return nil, errors.Wrapf(err, "request to %s", req.Endpoint())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("response status: " + res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read data")
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, errors.Wrapf(err, "unmarshal data: %s", string(data))
	}

	return res, nil
}
