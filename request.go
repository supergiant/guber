package guber

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	client    *Client
	method    string
	baseurl   string
	resource  string
	namespace string
	name      string
	body      []byte
	err       error
	response  *http.Response
}

func (r *Request) error(err error) {
	if r.err == nil {
		r.err = err
	}
}

func (r *Request) url() string {
	path := ""
	if r.namespace != "" {
		path = fmt.Sprintf("namespaces/%s/", r.namespace)
	}
	path = fmt.Sprintf("%s%s", path, r.resource)

	return fmt.Sprintf("%s/%s", r.baseurl, path)
}

func (r *Request) Resource(res Resource) *Request {
	baseurl := fmt.Sprintf("https://%s", r.client.Host)
	if res.DomainName() != "" {
		baseurl = fmt.Sprintf("%s/%s", baseurl, res.DomainName())
	}
	r.baseurl = fmt.Sprintf("%s/%s/%s", baseurl, res.ApiGroup(), res.ApiVersion())
	r.resource = res.ApiName()
	return r
}

func (r *Request) Namespace(namespace string) *Request {
	r.namespace = namespace
	return r
}

func (r *Request) Name(name string) *Request {
	r.name = name
	return r
}

func (r *Request) Entity(e Entity) *Request {
	body, err := json.Marshal(e)
	r.body = body
	r.error(err)
	return r
}

func (r *Request) Do() *Request {
	req, err := http.NewRequest(r.method, r.url(), bytes.NewBuffer(r.body))
	req.SetBasicAuth(r.client.Username, r.client.Password)
	r.error(err)
	resp, err := r.client.http.Do(req)
	r.response = resp
	r.error(err)
	return r
}

// The exit point for a Request (where error is pooped out)
func (r *Request) Into(e Entity) error {
	defer r.response.Body.Close()
	resp, err := ioutil.ReadAll(r.response.Body)
	r.error(err)
	json.Unmarshal(resp, e)
	return r.err
}
