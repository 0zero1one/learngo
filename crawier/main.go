package main

import (
	"fmt"
	"learngo/crawier/fetcher"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	contents, err := fetcher.Fetch("http://album.zhenai.com/u/1976352655")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(contents))
}
