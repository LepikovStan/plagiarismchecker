package textutils

import (
	"regexp"
	"strings"
)

var (
	reSentenceDivider = regexp.MustCompile("[.?!\n]+")
)

func SplitTextToSentences(text string) []string {
	txt := strings.ReplaceAll(text, "\\n", ".")
	var (
		ss     = reSentenceDivider.Split(txt, -1)
		result = make([]string, 0)
	)

	for i := 0; i < len(ss); i++ {
		result = append(result, strings.TrimSpace(ss[i]))
	}

	return result
}
