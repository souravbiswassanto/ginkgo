package ginkgo

import (
	"strings"
)

const (
	CategoryNovel        string = "Novel"
	CategoryShortStory   string = "ShortStory"
	CategoryThreeHundred string = "ThreeHundred"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

func (b *Book) Category() string {
	if b.Pages > 300 {
		return CategoryNovel
	} else if b.Pages < 300 {
		return CategoryShortStory
	}
	return CategoryThreeHundred
}

func (b *Book) AuthorLastName() string {
	if b.Author == "" {
		return ""
	}
	words := strings.Split(b.Author, " ")
	//fmt.Println(words)
	return words[len(words)-1] // last word
}

func (b *Book) IsValid() bool {
	if b.Author == "" {
		return false
	}
	return true
}

func (b *Book) AuthorFirstName() string {
	if b.Author == "" {
		return ""
	}
	words := strings.Split(b.Author, " ")
	if len(words) == 1 {
		return ""
	}
	return words[0]
}
