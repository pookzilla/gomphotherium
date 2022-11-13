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
	for _, character := range stringToWrap {
		characterWidth := runewidth.RuneWidth(character)
		if character == '\n' || character == '\r' {
			wordList.PushBack(word)
			out += DrainLine(*wordList, 0, 0)

			// out += word
			out += "\n"
			out += indentString

			word = ""
			wordList = list.New()
			spaceCount = 0
			committedCharacterCount = 0
			readWidth = 0
		} else if character == ' ' {
			wordList.PushBack(word)
			wordList.PushBack(StringPtr(" "))
			spaceCount++
			committedCharacterCount += characterCountOfCurrentWord

			characterCountOfCurrentWord = 0

			word = ""
			readWidth += characterWidth
		} else if readWidth + characterWidth > maximumWidth {
			// giant word - we need to print part of it
			if characterCountOfCurrentWord + characterWidth >= maximumWidth {
				wordList.PushBack(word)
				word = ""
			}
			readWidth = characterWidth
			characterCountOfCurrentWord = characterWidth //characterWidth
			word += string(character)

			leftoverSpace := 0
			if justifyText {
				leftoverSpace = int(math.Min(float64(maximumWidth-(committedCharacterCount + spaceCount)), float64(spaceCount) * 1.5))
			}
			out += DrainLine(*wordList, spaceCount, leftoverSpace)
			out += "\n"
			out += indentString
			spaceCount = 0
			committedCharacterCount = characterCountOfCurrentWord
			wordList = list.New()
		} else {
			readWidth += characterWidth
			characterCountOfCurrentWord += characterWidth
			word += string(character)
		}
	}
	out += DrainLine(*wordList, 0, 0)
	out += word
	return out
}
