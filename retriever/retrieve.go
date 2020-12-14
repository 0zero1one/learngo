package main

import (
	"fmt"
	"learngo/retriever/mock"
	real2 "learngo/retriever/real"
)

type retrieve interface {
	Get(url string) string
}

type poster interface {
	 Post(url string, from map[string]string) string
}

const url = "http://www.baidu.com"

func download(r retrieve, url string) string {
	return r.Get(url)
}

func post(poster poster) {
	poster.Post(url, map[string]string{
		"name": "afei",
		"course": "golang",
	})
}

type retrieverPoster interface {
	retrieve
	poster
}

func session(s retrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faker ",
	})
	return s.Get(url)
}

func main() {
	var r retrieve

	mockRetriever := mock.Download{
		Contents: "this is a fake http://www.baidu.com",
	}
	r = &mockRetriever
	inspect(r)

	// TYpe assertime
	if mockRetriever, ok := r.(*mock.Download); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("r is not a mock retriever")
	}

	fmt.Println("Try a session with mockRetriever")
	fmt.Println(session(&mockRetriever))
}

func inspect(r retrieve) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > Type:%T Value:%v\n", r, r)
	fmt.Print(" > Type switch: ")
	switch v := r.(type) {
	case *mock.Download:
		fmt.Println("Contents: ", v.Contents)
	case *real2.Download:
		fmt.Println("UserAgent: ", v.Contents)
	}
	fmt.Println()
}
