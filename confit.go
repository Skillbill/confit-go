package confit

import (
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	_hostname = "confit.skillbill.net"
	_port     = "443"
	_proto    = "https"
)

type Client struct {
	RepoId string
	Secret string
	Client http.Client
}

func init() {
	if h := os.Getenv("CONFIT_HOST"); h != "" {
		_hostname = h
	}
	if p := os.Getenv("CONFIT_PORT"); p != "" {
		_port = p
	}
	if i := os.Getenv("CONFIT_INSECURE"); i != "" {
		_proto = "http"
	}
}

func (c Client) buildURL(s string, isAlias bool) string {
	kind := "/path/"
	if isAlias {
		kind = "/alias/"
	}
	return _proto + "://" + _hostname + ":" + _port + "/api/repo/" + c.RepoId + kind + s
}

func (c Client) load(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "secret "+c.Secret)
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	return ioutil.ReadAll(res.Body)
}

func (c Client) LoadByPath(p string) ([]byte, error) {
	url := c.buildURL(p, false)
	return c.load(url)
}

func (c Client) LoadByAlias(a string) ([]byte, error) {
	url := c.buildURL(a, true)
	return c.load(url)
}
