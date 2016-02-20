package guber

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host     string
	Username string
	Password string
}

func (client *Client) url(path string) string {
	return fmt.Sprintf("https://%s/api/v1/%s", client.Host, path)
}

func (client *Client) Request(method string, path string) []byte {
	req, err := http.NewRequest(method, client.url(path), nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(client.Username, client.Password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return body
}

func (client *Client) getItems(resourceType string, query string, data interface{}) {
	path := resourceType
	if query != "" {
		path = fmt.Sprintf("%s?%s", path, query)
	}
	resp := client.Request("GET", path)
	if err := json.Unmarshal(resp, &data); err != nil {
		panic(err)
	}
}

func (client *Client) Nodes(query string) []*Node {
	var data struct{ Items []*Node }
	client.getItems("nodes", query, &data)

	// TODO is there a better way to do this so I don't have to repeat for each method below?
	for _, item := range data.Items {
		item.Client = client
	}

	return data.Items
}

func (client *Client) Events(query string) []*Event {
	var data struct{ Items []*Event }
	client.getItems("events", query, &data)
	return data.Items
}

func (client *Client) Pods(query string) []*Pod {
	var data struct{ Items []*Pod }
	client.getItems("pods", query, &data)
	return data.Items
}
