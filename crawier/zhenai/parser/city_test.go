package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err != nil {
		panic(err)
	}

	parseCity := ParseCity(contents)
	expectedUrls := []string{
		"http://album.zhenai.com/u/1976352655",
		"http://album.zhenai.com/u/1195899308",
		"http://album.zhenai.com/u/1807074256",
	}
	expectedNames := []string{
		"User: 天堂人满地狱打烊 Sex: 男士",
		"User: 唯一 Sex: 女士",
		"User: 余生有你 Sex: 女士",
	}
	requestSize := 20
	if requestSize != len(parseCity.Requests) {
		t.Errorf("result should have %d requests; but had %d", requestSize, len(parseCity.Requests))
	}
	for i, url := range expectedUrls {
		if parseCity.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, parseCity.Requests[i].Url)
		}
	}

	if requestSize != len(parseCity.Requests) {
		t.Errorf("result should have %d items; but had %d", requestSize, len(parseCity.Items))
	}
	for i, name := range expectedNames {
		if parseCity.Items[i].(string) != name {
			t.Errorf("expected url #%d: %s; but was %s", i, name, parseCity.Items[i].(string))
		}
	}
}
