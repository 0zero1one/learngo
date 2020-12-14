package parser

import (
	"fmt"
	"learngo/crawier/engine"
	"regexp"
)

const messageRe = `<div class="m-btn purple"[^<]*>([^<]+)</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	fmt.Printf("%s", string(contents))
	re := regexp.MustCompile(messageRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		fmt.Println(m)
		//result.Items = append(result.Items, "User: "+string(m[2])+" Sex: "+string(m[3]))
		//result.Requests = append(result.Requests, engine.Request{
		//	Url:        string(m[1]),
		//	ParserFunc: engine.NilParser,
		//})
	}

	return result
}
