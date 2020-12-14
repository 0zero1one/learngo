package parser

import (
	"learngo/crawier/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>[^性]*性别：</span>([^<]+)</td>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User: "+string(m[2])+" Sex: "+string(m[3]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseProfile,
		})
	}

	return result
}
