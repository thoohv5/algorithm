package trie

import (
	"unicode/utf8"
)

type ISensWords interface {
	AddWord(sensWord string)
	Replace(text string) (newText string)
}

type sensWords struct {
	replace rune
	root    *Node
}

func New() ISensWords {
	return &sensWords{
		replace: '*',
		root:    &Node{},
	}
}

func (sw *sensWords) AddWord(sensWord string) {

	curNode := sw.root
	sensChar := []rune(sensWord)
	for _, sc := range sensChar {
		curNode = curNode.AddChild(sc)
	}
	curNode.FullContent = sensWord
	curNode.End = true

}

func (sw *sensWords) Replace(text string) string {

	dealText := []rune(text)

	curNode := sw.root
	for index, tc := range dealText {
		curNode = curNode.FindChild(tc)
		if nil == curNode {
			curNode = sw.root
			continue
		}
		if curNode.End {

			for i := utf8.RuneCountInString(curNode.FullContent) - 1; i >= 0; i-- {
				dealText[index-i] = sw.replace
			}

			curNode = sw.root
		}

	}

	return string(dealText)
}
