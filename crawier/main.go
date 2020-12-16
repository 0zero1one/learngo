package main

import (
	"net/http"
)
import "fmt"

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	resp, err := http.Get("https://album.zhenai.com/api/profile/getObjectProfile.do?objectID=1976352655&_=16" +
		"08100599437&ua=h5%2F1.0.0%2F1%2F0%2F0%2F0%2F0%2F0%2F" +
		"%2F0%2F0%2F4450d139-9bea-4f8b-8c53-16d06c0783f3%2F0%2F0%2F1578333505")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%v", resp.Body)
}
