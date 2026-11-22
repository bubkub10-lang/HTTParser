package htmlRead

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func Read(url string) []byte {
	req, err := http.Get(url)
	if err != nil {
		fmt.Println("GetFATAL")
		return nil
	} else {
		defer req.Body.Close()
		B, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Println("ReadFATAL")
			return nil
		} else {
			return B
		}
	}
}

func Sort(x []byte) {
	var sliced []string
	stay := [29]string{
		"p",
		"h1",
		"h2",
		"h3",
		"h4",
		"h5",
		"h6",
		"article",
		"section",
		"main",
		"blockquote",
		"q",
		"pre",
		"code",
		"li",
		"dt",
		"dd",
		"td",
		"th",
		"caption",
		"div",
		"span",
		"a",
		"b",
		"strong",
		"i",
		"em",
		"mark",
		"small",
	}
	htmlstr := string(x)
	tokenizer := html.NewTokenizer(strings.NewReader(htmlstr))

	for {
		tagType := tokenizer.Next()
		if tagType == html.ErrorToken {
			break
		}
		tag := html.StartTagToken
		strTag := string(tag)
		for _, stayed := range stay {
			if strings.ContainsAny(strTag, stayed) {
				sliced = append(sliced, strTag)
			}
		}
	}
}
