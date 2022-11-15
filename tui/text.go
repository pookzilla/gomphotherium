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

	out := indentString
	word := ""
	for _, character := range stringToWrap {
		characterWidth := runewidth.RuneWidth(character)
		if character == '\n' || character == '\r' {
      if (characterCountOfCurrentWord + committedCharacterCount >= maximumWidth) {
        out += "\n"
        out += indentString
      }
			wordList.PushBack(word)
      line := DrainLine(*wordList, 0, 0)
			out += line
			out += "\n"
			out += indentString

			word = ""
      characterCountOfCurrentWord = 0
			wordList = list.New()
			spaceCount = 0
			committedCharacterCount = 0
		} else if character == ' ' {
			wordList.PushBack(word)
			wordList.PushBack(StringPtr(" "))
			spaceCount++
			committedCharacterCount += characterCountOfCurrentWord + 1

			characterCountOfCurrentWord = 0
			word = ""
		} else if committedCharacterCount+runewidth.StringWidth(word)+characterWidth > maximumWidth {

      // giant word - we need to print part of it
			if characterCountOfCurrentWord + characterWidth >= maximumWidth {
				wordList.PushBack(word)
				word = ""
        characterCountOfCurrentWord = 0
			}

			word += string(character)
      characterCountOfCurrentWord += characterWidth

			leftoverSpace := 0
			if justifyText {
				leftoverSpace = int(math.Min(float64(maximumWidth-(characterCountOfCurrentWord+spaceCount)), float64(spaceCount)*1.5))
			}
      line := DrainLine(*wordList, spaceCount, leftoverSpace)
			out += line
			out += "\n"
			out += indentString
			spaceCount = 0
			committedCharacterCount = 0
			wordList = list.New()
		} else {
			characterCountOfCurrentWord += characterWidth
			word += string(character)
		}
	}
  line := DrainLine(*wordList, 0, 0)
	out += line
	out += word
	return out
}
