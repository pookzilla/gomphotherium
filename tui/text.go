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

func WrapWithIndent(
	stringToWrap string, 
	maximumWidth int, 
	indentString string, 
	justifyText bool,
	) string {

	wordList := list.New()

	spaceCount := 0
	committedCharacterCount := 0
	characterCountOfCurrentWord := 0

	readWidth := 0
	out := indentString
	word := ""
	for _, r := range stringToWrap {
		characterWidth := runewidth.RuneWidth(r)
		if r == '\n' || r == '\r' {
			out += DrainLine(*wordList, 0, 0)

			out += word
			out += string('\n')
			out += indentString

			word = ""
			wordList = list.New()
			spaceCount = 0
			committedCharacterCount = 0
			readWidth = 0
			continue
		} else if r == ' ' {
			wordList.PushBack(word)
			wordList.PushBack(StringPtr(" "))
			spaceCount++
			committedCharacterCount += characterCountOfCurrentWord

			characterCountOfCurrentWord = 0

			word = ""
			readWidth += characterWidth
			continue
		} else if readWidth + characterWidth > maximumWidth {
			characterCountOfCurrentWord += characterWidth
			// giant word - we need to print part of it
			if characterCountOfCurrentWord >= maximumWidth {
				wordList.PushBack(word[:maximumWidth-1])
				word = word[maximumWidth-1:]
				characterCountOfCurrentWord = characterWidth
			}

			leftoverSpace := 0
			if justifyText {
				leftoverSpace = int(math.Min(float64(maximumWidth-(committedCharacterCount + spaceCount)), float64(spaceCount) * 1.5))
			}
			out += DrainLine(*wordList, spaceCount, leftoverSpace)

			out += "\n"
			out += indentString
			spaceCount = 0
			committedCharacterCount = characterCountOfCurrentWord
			readWidth = characterCountOfCurrentWord
			wordList = list.New()

			word += string(r)

			readWidth += characterWidth
			continue

		} else {
			readWidth += characterWidth
			word += string(r)
			characterCountOfCurrentWord += characterWidth
			continue
		}
	}
	out += DrainLine(*wordList, 0, 0)
	out += word
	return out
}
