package textutils

import (
	"regexp"
	"strings"

	"github.com/neurosnap/sentences/english"
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

func SplitToBatches(text string, batchSize int) []string {
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		panic(err)
	}
	sents := tokenizer.Tokenize(text)

	var (
		batches         = make([]string, 0)
		batchCharsCount = 0
		batch           string
	)
	for i := 0; i < len(sents); i++ {
		sent := strings.TrimSpace(sents[i].Text)
		sentCharsCount := strings.Count(sent, "")

		if batchCharsCount+sentCharsCount > batchSize {
			batches = append(batches, batch)
			batch = ""
			batchCharsCount = 0
		}
		batch = batch + " " + sent
		batchCharsCount += sentCharsCount

		if i == len(sents)-1 {
			batches = append(batches, batch)
		}
	}

	return batches
}
