package real

import (
	"net/http"
	"net/http/httputil"
)

type Download struct {
	Contents string
}

func (d Download) Get(url string) string {
	reap, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(reap, true)
	reap.Body.Close()
	if err != nil {
		panic(err)
	}

	return d.Contents + string(result)
}
