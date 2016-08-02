package guber

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"
)

type Error404 struct{}

func (e *Error404) Error() string {
	return "Resource not found"
}

type Error409 struct{}

func (e *Error409) Error() string {
	return "Resource already exists"
}

type Request struct {
	client    *RealClient
	method    string
	headers   map[string]string
	basePath  string
	query     string
	path      string
	resource  string
	namespace string
	name      string
	body      []byte

	err          error
	response     *http.Response
	responseBody []byte
}

// Implement Stringer interface
func (r *Request) String() string {
	obj := struct {
		Method       string
		Headers      map[string]string
		URL          string
		Status       int
		RequestBody  string
		ResponseBody string
	}{
		r.method,
		r.headers,
		r.url(),
		r.response.StatusCode,
		string(r.body),
		string(r.responseBody),
	}
	out, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(out)
}

func (r *Request) error(err error) {
	if err != nil && r.err == nil {
		r.err = err
	}
}

func (r *Request) url() string {
	resourcePath := path.Join(r.resource, r.name, r.path)

	if r.namespace != "" {
		resourcePath = path.Join("namespaces", r.namespace, resourcePath)
	}
	if r.query != "" {
		resourcePath += "?" + r.query
	}
	return "https://" + path.Join(r.client.Host, r.basePath, resourcePath)
}

func (r *Request) Collection(c Collection) *Request {
	m := c.Meta()
	r.basePath = path.Join(m.DomainName, m.APIGroup, m.APIVersion)
	r.resource = m.APIName
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

func (r *Request) Query(q *QueryParams) *Request {
	if q == nil {
		return r
	}

	var segments []string
	if ls := q.LabelSelector; ls != "" {
		segments = append(segments, "labelSelector="+ls)
	}
	if fs := q.FieldSelector; fs != "" {
		segments = append(segments, "fieldSelector="+fs)
	}
	r.query = strings.Join(segments, "&")

	return r
}

func (r *Request) Path(path string) *Request {
	r.path = path
	return r
}

func (r *Request) Do() *Request {
	req, err := http.NewRequest(r.method, r.url(), bytes.NewBuffer(r.body))
	if err != nil {
		panic(err) // TODO
	}

	req.SetBasicAuth(r.client.Username, r.client.Password)
	r.error(err)

	for k, v := range r.headers {
		req.Header.Set(k, v)
	}

	resp, err := r.client.http.Do(req)
	r.error(err)

	// TODO
	if resp != nil {
		r.response = resp

		r.readBody()

		if resp.StatusCode == 404 {
			r.error(new(Error404))
		} else if resp.StatusCode == 409 {
			r.error(new(Error409))
		} else if status := resp.Status; status[:2] != "20" {
			r.error(fmt.Errorf("Status: %s, Body: %s", status, string(r.responseBody)))
		}

		Log.Debug(r)
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
	r.responseBody = body
	r.error(err)
}

func (r *Request) Body() (string, error) {
	return string(r.responseBody), r.err
}

// The exit point for a Request (where error is pooped out)
func (r *Request) Into(e Entity) error {
	if r.responseBody != nil {
		json.Unmarshal(r.responseBody, e)
	}
	return r.err
}
