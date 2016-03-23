package guber

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	client    *Client
	method    string
	baseurl   string
	query     string
	resource  string
	namespace string
	name      string
	body      []byte

	// NOTE this is used distinct from err, because a 404 is not technically an
	// error, except to the end-user who expects a resource to be there.
	// Without this, we don't have a way to determine if an err was a 404 or
	// something lower-level without inspecting the error message.
	found bool

	err      error
	response *http.Response
}

func (r *Request) error(err error) {
	if err != nil && r.err == nil {
		fmt.Println("REQUEST ERROR", err)
		r.err = err
	}
}

func (r *Request) url() string {
	path := ""
	if r.namespace != "" {
		path = fmt.Sprintf("namespaces/%s/", r.namespace)
	}
	path = path + r.resource
	if r.name != "" {
		path = path + "/" + r.name
	}
	if r.query != "" {
		path = path + "?" + r.query
	}
	return r.baseurl + "/" + path
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

	// TODO
	fmt.Println("Req body: ", string(body))

	r.body = body
	r.error(err)
	return r
}

func (r *Request) Query(q *QueryParams) *Request {
	if q == nil {
		return r
	}

	// v, err := query.Values(q)
	// if err != nil {
	// 	panic(err) // TODO should use r.error() here probably
	// }
	// queryStr := v.Encode()

	// TODO  -- we went with this terribly rigid strategy because of how query pkg encodes the = chars
	if ls := q.LabelSelector; ls != "" {
		r.query = "labelSelector=" + ls
	}

	return r
}

func (r *Request) Do() *Request {
	req, err := http.NewRequest(r.method, r.url(), bytes.NewBuffer(r.body))
	if err != nil {
		panic(err) // TODO
	}

	req.SetBasicAuth(r.client.Username, r.client.Password)
	r.error(err)

	// TODO
	// fmt.Println(r.url())
	fmt.Println(*req)

	resp, err := r.client.http.Do(req)
	r.error(err)

	// TODO
	if resp != nil {
		r.response = resp

		r.readBody()

		if resp.StatusCode == 404 {
			r.found = false
		} else if status := resp.Status; status[:2] != "20" {
			errMsg := fmt.Sprintf("Status: %s, Body: %s", status, string(r.body))
			r.error(errors.New(errMsg))
		} else {
			r.found = true // NOTE this only really matters for lookups, but we set it true here anyhow
		}
	}
	return r
}

func (r *Request) readBody() {
	if r.response == nil {
		r.error(errors.New("Response is nil"))
		return
	}
	defer r.response.Body.Close()
	body, err := ioutil.ReadAll(r.response.Body)
	r.body = body
	r.error(err)
}

// The exit point for a Request (where error is pooped out)
func (r *Request) Into(e Entity) error {
	if r.body != nil {
		json.Unmarshal(r.body, e)
	}
	return r.err
}
