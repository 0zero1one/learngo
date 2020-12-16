package parser

import (
	"fmt"
	"learngo/crawier/engine"
)

const messageRe = `<div class="m-btn purple"[^<]*>([^<]+)</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	//fmt.Printf("%s", string(contents))
	//re := regexp.MustCompile(messageRe)
	//matches := re.FindAllSubmatch(contents, -1)
	//
	//result := engine.ParseResult{}
	//for _, m := range matches {
	//	fmt.Println(m)
	//}
	fmt.Println(string(contents))
	return engine.ParseResult{}
}
