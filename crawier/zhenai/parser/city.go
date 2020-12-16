package parser

import (
	"learngo/crawier/engine"
	"regexp"
)

const cityRe = `<a href="http://album.zhenai.com/u/([0-9]+)"[^>]+>([^<]+)</a>[^性]*性别：</span>[^<]+</td>`
const dataUrl = "https://album.zhenai.com/api/profile/getObjectProfile.do"

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User: "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: dataUrl + "?objectID=" + string(m[1]) + "&_=1608100599437&ua=h5%2F1.0.0%" +
				"2F1%2F0%2F0%2F0%2F0%2F0%2F%2F0%2F0%2F4450d139-9bea-4f8b-8c53-16d06c0783f3%2F0%2F0%2F1578333505",
			ParserFunc: ParseProfile,
		})
	}

	return result
}
