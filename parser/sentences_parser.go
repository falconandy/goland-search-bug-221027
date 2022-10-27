package parser

type SentencesParser interface {
	Parse(text string) ([]string, error)
}
