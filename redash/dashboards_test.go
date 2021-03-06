package redash

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDashboards(t *testing.T) {
	const listDashboardsResBody = `{"something":"something","results":[{"something":"something"}]}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards" {
			fmt.Fprint(w, listDashboardsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *ListDashboardsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: listDashboardsResBody},
		{input: &ListDashboardsInput{}, status: 200, body: listDashboardsResBody},
	}

	for _, c := range testCases {
		result := testClient.ListDashboards(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDashboard(t *testing.T) {
	const getDashboardResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards/something" {
			fmt.Fprint(w, getDashboardResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDashboardInput
		status int
		body   string
	}{
		{input: &GetDashboardInput{DashboardSlug: "something"}, status: 200, body: getDashboardResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDashboard(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetPublicDashboard(t *testing.T) {
	const getPublicDashboardResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards/public/dummy" {
			fmt.Fprint(w, getPublicDashboardResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetPublicDashboardInput
		status int
		body   string
	}{
		{input: &GetPublicDashboardInput{Token: "dummy"}, status: 200, body: getPublicDashboardResBody},
	}

	for _, c := range testCases {
		result := testClient.GetPublicDashboard(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}

func TestGetDashboardTags(t *testing.T) {
	const getDashboardTagsResBody = `{"something":"something"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/dashboards/tags" {
			fmt.Fprint(w, getDashboardTagsResBody)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	testClient := NewClient(&Config{EndpointUrl: ts.URL, ApiKey: "dummy"})

	testCases := []struct {
		input  *GetDashboardTagsInput
		status int
		body   string
	}{
		{input: nil, status: 200, body: getDashboardTagsResBody},
		{input: &GetDashboardTagsInput{}, status: 200, body: getDashboardTagsResBody},
	}

	for _, c := range testCases {
		result := testClient.GetDashboardTags(c.input)
		if result.StatusCode != c.status || result.Body != c.body {
			t.Errorf("Unexpected response: status:%+v != %+v, body:%+v != %+v", result.StatusCode, c.status, result.Body, c.body)
		}
	}
}
