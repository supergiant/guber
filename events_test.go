package guber

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	tevents = &Events{
		client:    tClient,
		Namespace: "test",
	}
)

func TestDomainName(t *testing.T) {
	resp := tevents.DomainName()
	expected := ""

	if expected != resp {
		t.Error("ERROR .DomainName(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiGroup(t *testing.T) {
	resp := tevents.ApiGroup()
	expected := "api"

	if expected != resp {
		t.Error("ERROR .ApiGroup(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiVersion(t *testing.T) {
	resp := tevents.ApiVersion()
	expected := "v1"

	if expected != resp {
		t.Error("ERROR .ApiVersion(): expected,", expected, "-- But got,", resp)
	}
}

func TestApiName(t *testing.T) {
	resp := tevents.ApiName()
	expected := "events"

	if expected != resp {
		t.Error("ERROR .ApiName(): expected,", expected, "-- But got,", resp)
	}
}

func TestKind(t *testing.T) {
	resp := tevents.Kind()
	expected := "Event"

	if expected != resp {
		t.Error("ERROR .Kind(): expected,", expected, "-- But got,", resp)
	}
}

func TestCreate(t *testing.T) {
	// Setup our test server
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Much Success")
	}))
	defer ts.Close()

	// test client
	url := strings.Replace(ts.URL, "https://", "", -1)
	tsClient := &Client{
		Host:     url,
		Username: "test",
		Password: "test",
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}

	tsEvents := &Events{
		client:    tsClient,
		Namespace: "test",
	}
	resp, _ := tsEvents.Create(&Event{
		Message: "test",
		Count:   1,
	})

	expected := &Event{
		Message: "test",
		Count:   1,
	}

	if !reflect.DeepEqual(expected, resp) {
		t.Error("ERROR .Create(): expected,", expected, "-- But got,", resp)
	}
}

func TestCreateError(t *testing.T) {
	_, err := tevents.Create(&Event{
		Message: "test",
		Count:   1,
	})

	if err == nil {
		t.Error("ERROR .Create(): event create to fail.. but it did not. ")
	}
}
