package strgo

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

const asciiMaxDec = uint8(254)
const asciiMaxDecInt32 = int32(254)

type Condition struct {
	MinLength                 int
	MaxLength                 int
	OnlyContains              []byte
	OnlyContainsPrefix        []byte
	OnlyContainsSuffix        []byte
	OnlyContainsPrefixWord    []string
	OnlyContainsSuffixWord    []string
	MustContains              []byte
	MustContainsWord          []string
	MustContainsOnce          []byte
	MustContainsWordOnce      []string
	MustNotContains           []byte
	MustNotContainsWord       []string
	MustNotContainsPrefix     []byte
	MustNotContainsSuffix     []byte
	MustNotContainsPrefixWord []string
	MustNotContainsSuffixWord []string
	MayContainsOnce           []byte
	MayContainsWordOnce       []string
	MustBeFollowedBy          [2][]byte
}

type condition struct {
	minLength   int
	maxLength   int
	iterateText bool

	hasAscii bool
	/*
		ascii iterate conditions index:
		0: OnlyContains
		1: OnlyContainsPrefix
		2: OnlyContainsSuffix
		3: MustContains
		4: MustNotContains
		5: MustNotContainsPrefix
		6: MustNotContainsSuffix
		7: MayContainsOnce
		8: MustBeFollowedBy
	*/
	ascii                 [9]bool
	asciiC                [9][255]byte
	asciiMustBeFollowedBy [255][]byte

	hasWord bool
	/*
		word iterate conditions index:
		0: OnlyContainsPrefixWord
		1: OnlyContainsSuffixWord
		2: MustContainsWord
		3: MustContainsWordOnce
		4: MustNotContainsWord
		5: MustNotContainsPrefixWord
		6: MustNotContainsSuffixWord
		7: MayContainsWordOnce
	*/
	word  [8]bool
	words [8][]string
}

// Validate matches the string based on the given conditions.
// If one doesn't match, it will return an error.
func Validate(text string, conds ...*Condition) error {
	if text == "" {
		return errors.New("the string is empty")
	}
	if len(conds) < 1 {
		return nil
	}
	textLen := len(text)
	cond := setCond(conds...)

	if cond.minLength > 0 && textLen < cond.minLength {
		return errors.New("the string length cannot be less than " + strconv.Itoa(cond.minLength))
	}
	if cond.maxLength > 0 && textLen > cond.maxLength {
		return errors.New("the string length cannot be more than " + strconv.Itoa(cond.maxLength))
	}

	// ascii

	if cond.hasAscii {
		if cond.ascii[1] {
			if text[0] > asciiMaxDec {
				return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
			}
			if cond.asciiC[1][text[0]] < 1 {
				return errors.New("the string cannot contain prefix char: " + string(text[0]))
			}
		}
		if cond.ascii[2] {
			if text[textLen-1] > asciiMaxDec {
				return errors.New("the char: " + string(text[textLen-1]) + ", is not a valid ascii format")
			}
			if cond.asciiC[2][text[textLen-1]] < 1 {
				return errors.New("the string cannot contain suffix char: " + string(text[textLen-1]))
			}
		}
		if cond.ascii[5] {
			if text[0] > asciiMaxDec {
				return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
			}
			if cond.asciiC[5][text[0]] > 0 {
				return errors.New("the string must not contain prefix: " + string(text[0]))
			}
		}
		if cond.ascii[6] {
			if text[textLen-1] > asciiMaxDec {
				return errors.New("the char: " + string(text[textLen-1]) + ", is not a valid ascii format")
			}
			if cond.asciiC[6][text[textLen-1]] > 0 {
				return errors.New("the string must not contain suffix: " + string(text[textLen-1]))
			}
		}
		if cond.iterateText {
			for i, v := range text {
				if v > asciiMaxDecInt32 {
					return errors.New("the char: " + string(v) + ", is not a valid ascii format")
				}
				if cond.ascii[0] && cond.asciiC[0][v] < 1 {
					return errors.New("the string cannot contain char: " + string(byte(v)))
				}
				if cond.ascii[4] && cond.asciiC[4][v] > 0 {
					return errors.New("the string must not contain char: " + string(byte(v)))
				}
				if cond.ascii[3] && cond.asciiC[3][v] > 0 {
					cond.asciiC[3][v] = 0
				}
				if cond.ascii[7] && cond.asciiC[7][v] > 0 {
					if cond.asciiC[7][v] > 1 {
						return errors.New("the char: " + string(byte(v)) + ", must be appeared once in the string")
					}
					cond.asciiC[7][v] += 1
				}

				if cond.ascii[8] {
					if cond.asciiMustBeFollowedBy[v] != nil {
						bf := text[i]
						af := text[i]
						if i > 0 && i < textLen {
							bf = text[i-1]
						}
						if (i + 1) < textLen {
							af = text[i+1]
						}
						if bytes.IndexByte(cond.asciiMustBeFollowedBy[v], bf) < 0 || bytes.IndexByte(cond.asciiMustBeFollowedBy[v], af) < 0 {
							return errors.New("the char: " + string(byte(v)) + ", must be followed with at least one of these characters: " + string(cond.asciiMustBeFollowedBy[v]))
						}
					}
				}
			}
			if cond.ascii[3] {
				for b, v := range cond.asciiC[3] {
					if v > 0 {
						return errors.New("the string must contain char: " + string(rune(b)))
					}
				}
			}
		}
	}

	// word

	if cond.hasWord {
		if cond.word[0] {
			for _, v := range cond.words[0] {
				if v != "" && text[:len(v)] == v {
					cond.words[0] = nil
					break
				}
			}
			if cond.words[0] != nil {
				return errors.New("the string prefix doesn't match with the given prefix words")
			}
		}
		if cond.word[1] {
			for _, v := range cond.words[1] {
				if v != "" && text[textLen-len(v):] == v {
					cond.words[1] = nil
					break
				}
			}
			if cond.words[1] != nil {
				return errors.New("the string suffix doesn't match with the given suffix words")
			}
		}
		if cond.word[2] {
			for _, v := range cond.words[2] {
				if v != "" && !strings.Contains(text, v) {
					return errors.New("the string must contain word: " + v)
				}
			}
		}
		if cond.word[3] {
			for _, v := range cond.words[3] {
				if v != "" && strings.Count(text, v) != 1 {
					return errors.New("the string must contain word: " + v + ", and it must be appeared once in the string")
				}
			}
		}
		if cond.word[4] {
			for _, v := range cond.words[4] {
				if v != "" && strings.Contains(text, v) {
					return errors.New("the string must not contain word: " + v)
				}
			}
		}
		if cond.word[5] {
			for _, v := range cond.words[5] {
				if v != "" && text[:len(v)] == v {
					return errors.New("the string must not contain prefix word: " + v)
				}
			}
		}
		if cond.word[6] {
			for _, v := range cond.words[6] {
				if v != "" && text[textLen-len(v):] == v {
					return errors.New("the string must not contain suffix word: " + v)
				}
			}
		}
		if cond.word[7] {
			for _, v := range cond.words[7] {
				if v != "" && strings.Count(text, v) > 1 {
					return errors.New("the word: " + v + ", must be appeared once in the string")
				}
			}
		}
	}
	return nil
}

func setCond(conds ...*Condition) *condition {
	cond := &condition{}
	for _, v := range conds {
		cond.minLength = v.MinLength
		cond.maxLength = v.MaxLength
		if v.OnlyContains != nil {
			cond.iterateText = true
			setASCII(cond, v.OnlyContains, 0)
		}
		if v.OnlyContainsPrefix != nil {
			setASCII(cond, v.OnlyContainsPrefix, 1)
		}
		if v.OnlyContainsSuffix != nil {
			setASCII(cond, v.OnlyContainsSuffix, 2)
		}
		if v.MustContains != nil {
			cond.iterateText = true
			setASCII(cond, v.MustContains, 3)
		}
		if v.MustContainsOnce != nil {
			cond.iterateText = true
			setASCIIDouble(cond, v.MustContainsOnce, 3, 7)
		}
		if v.MustNotContains != nil {
			cond.iterateText = true
			setASCII(cond, v.MustNotContains, 4)
		}
		if v.MustNotContainsPrefix != nil {
			setASCII(cond, v.MustNotContainsPrefix, 5)
		}
		if v.MustNotContainsSuffix != nil {
			setASCII(cond, v.MustNotContainsSuffix, 6)
		}
		if v.MayContainsOnce != nil {
			cond.iterateText = true
			setASCII(cond, v.MayContainsOnce, 7)
		}
		if v.MustBeFollowedBy[0] != nil && v.MustBeFollowedBy[1] != nil {
			cond.iterateText = true
			setASCIIIteMustBeFollowedBy(cond, v.MustBeFollowedBy[0], v.MustBeFollowedBy[1])
		}
		if v.OnlyContainsPrefixWord != nil {
			setWord(cond, v.OnlyContainsPrefixWord, 0)
		}
		if v.OnlyContainsSuffixWord != nil {
			setWord(cond, v.OnlyContainsSuffixWord, 1)
		}
		if v.MustContainsWord != nil {
			setWord(cond, v.MustContainsWord, 2)
		}
		if v.MustContainsWordOnce != nil {
			setWord(cond, v.MustContainsWordOnce, 3)
		}
		if v.MustNotContainsWord != nil {
			setWord(cond, v.MustNotContainsWord, 4)
		}
		if v.MustNotContainsPrefixWord != nil {
			setWord(cond, v.MustNotContainsPrefixWord, 5)
		}
		if v.MustNotContainsSuffixWord != nil {
			setWord(cond, v.MustNotContainsSuffixWord, 6)
		}
		if v.MayContainsWordOnce != nil {
			setWord(cond, v.MayContainsWordOnce, 7)
		}
	}
	return cond
}

func setASCII(cond *condition, c []byte, i byte) {
	cond.hasAscii = true
	cond.ascii[i] = true
	for _, b := range c {
		cond.asciiC[i][b] = 1
	}
}

func setASCIIDouble(cond *condition, c []byte, i, i2 byte) {
	cond.hasAscii = true
	cond.ascii[i] = true
	cond.ascii[i2] = true
	for _, b := range c {
		cond.asciiC[i][b] = 1
		cond.asciiC[i2][b] = 1
	}
}

func setASCIIIteMustBeFollowedBy(cond *condition, c, fc []byte) {
	cond.hasAscii = true
	cond.ascii[8] = true
	for _, b := range c {
		cond.asciiMustBeFollowedBy[b] = append(cond.asciiMustBeFollowedBy[b], fc...)
	}
}

func setWord(cond *condition, w []string, i byte) {
	cond.hasWord = true
	cond.word[i] = true
	cond.words[i] = append(cond.words[i], w...)
}
