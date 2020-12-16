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

	contents, err := fetcher.Fetch("https://album.zhenai.com/u/1210022635")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", string(contents))
}
