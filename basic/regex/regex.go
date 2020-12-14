package main

import (
	"fmt"
	"regexp"
)

var text string = `
i am three emails,
first is 2819586985@qq.com,
second is rjxszbh@163.com,
last is rjxszbh@gmail.com。
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
