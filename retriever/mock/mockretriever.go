package mock

import "fmt"

type Download struct {
	Contents string
}

func (d Download) Get(url string) string {
	return d.Contents + url
}

func (d *Download) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", d.Contents)
}

func (d *Download) Post(url string, form map[string]string) string {
	d.Contents = form["contents"]
	return "ok"
}
