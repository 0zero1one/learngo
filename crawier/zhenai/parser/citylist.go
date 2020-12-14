package parser

import (
	"learngo/crawier/engine"
	"regexp"
)

const cityListRe = `<a[^h]*href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	limit := 1
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}

	return result
}
