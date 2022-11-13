package tui

import (
	"container/list"
	"math"
	"math/rand"

	"github.com/mattn/go-runewidth"
)

func StringPtr(s string) *string {
	return &s
}

func DrainLine(wordList list.List, numberOfSpaces int, leftoverSpaces int) string {
	out := ""
	spaceList := make([]*string, 0)
	for e := wordList.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(*string)
		if ok {
			spaceList = append(spaceList, value)
		}
	}

	for i := leftoverSpaces; i > 0; i-- {
		selection := rand.Intn(numberOfSpaces)
		*spaceList[selection] += " "
	}

	for e := wordList.Front(); e != nil; e = e.Next() {
		value, ok := e.Value.(string)
		if ok {
			out += value
		} else {
			value, ok := e.Value.(*string)
			if ok {
				out += *value
			}
		}
	}
	return out
}

func WrapWithIndent(s string, w int, indentString string, justifyText bool) string {
	wordList := list.New()

	spaceCount := 0
	characterCount := 0
	characterCountOfCurrentWord := 0

	width := 0
	out := indentString
	word := ""
	for _, r := range s {
		cw := runewidth.RuneWidth(r)
		if r == '\n' || r == '\r' {
			out += DrainLine(*wordList, 0, 0)

			out += word
			out += string('\n')
			out += indentString

			word = ""
			wordList = list.New()
			spaceCount = 0
			characterCount = 0
			width = 0
			continue
		} else if r == ' ' {
			wordList.PushBack(word)
			wordList.PushBack(StringPtr(" "))
			spaceCount++
			characterCount += characterCountOfCurrentWord

			characterCountOfCurrentWord = 0

			word = ""
			width += cw
			continue
		} else if width+cw > w {
			characterCountOfCurrentWord += cw
			
			leftoverSpace := 0
			if justifyText {
				leftoverSpace = int(math.Min(float64(w-(characterCount+spaceCount)), float64(spaceCount*2) /*float64(width/5)*/))
			}
			out += DrainLine(*wordList, spaceCount, leftoverSpace)

			out += "\n"
			out += indentString
			spaceCount = 0
			characterCount = characterCountOfCurrentWord
			width = characterCountOfCurrentWord
			wordList = list.New()

			word += string(r)

			width += cw
			continue

		} else {
			width += cw
			word += string(r)
			characterCountOfCurrentWord += cw
			continue
		}
	}
	out += DrainLine(*wordList, 0, 0)
	out += word
	return out
}
