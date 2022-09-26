package strgo

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

const asciiMaxDec = uint8(254)
const asciiMaxDecInt32 = int32(254)

type asciis [asciiMaxDec + 1]byte

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

// Validate matches the string based on the given condition.
// If one doesn't match, it will return an error.
func Validate(text string, cond *Condition) error {
	if text == "" {
		return errors.New("the string is empty")
	}

	textLen := len(text)

	if cond.MinLength > 0 && textLen < cond.MinLength {
		return errors.New("the string length cannot be less than " + strconv.Itoa(cond.MinLength))
	}
	if cond.MaxLength > 0 && textLen > cond.MaxLength {
		return errors.New("the string length cannot be more than " + strconv.Itoa(cond.MaxLength))
	}

	if cond.OnlyContainsPrefix != nil {
		if text[0] > asciiMaxDec {
			return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexByte(cond.OnlyContainsPrefix, text[0]) < 0 {
			return errors.New("the string cannot contain prefix char: " + string(text[0]))
		}
	}
	if cond.OnlyContainsSuffix != nil {
		if text[0] > asciiMaxDec {
			return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexByte(cond.OnlyContainsSuffix, text[textLen-1]) < 0 {
			return errors.New("the string cannot contain suffix char: " + string(text[textLen-1]))
		}
	}
	if cond.MustNotContainsPrefix != nil {
		if text[0] > asciiMaxDec {
			return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexByte(cond.MustNotContainsPrefix, text[0]) >= 0 {
			return errors.New("the string must not contain prefix: " + string(text[0]))
		}
	}
	if cond.MustNotContainsSuffix != nil {
		if text[0] > asciiMaxDec {
			return errors.New("the char: " + string(text[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexByte(cond.MustNotContainsSuffix, text[textLen-1]) >= 0 {
			return errors.New("the string must not contain suffix: " + string(text[textLen-1]))
		}
	}
	if cond.OnlyContainsPrefixWord != nil {
		for _, v := range cond.OnlyContainsPrefixWord {
			if v != "" && text[:len(v)] == v {
				cond.OnlyContainsPrefixWord = nil
				break
			}
		}
		if cond.OnlyContainsPrefixWord != nil {
			return errors.New("the string prefix doesn't match with the given prefix words")
		}
	}
	if cond.OnlyContainsSuffixWord != nil {
		for _, v := range cond.OnlyContainsSuffixWord {
			if v != "" && text[textLen-len(v):] == v {
				cond.OnlyContainsSuffixWord = nil
				break
			}
		}
		if cond.OnlyContainsSuffixWord != nil {
			return errors.New("the string suffix doesn't match with the given suffix words")
		}
	}
	if cond.MustNotContainsPrefixWord != nil {
		for _, v := range cond.MustNotContainsPrefixWord {
			if v != "" && text[:len(v)] == v {
				return errors.New("the string must not contain prefix word: " + v)
			}
		}
	}
	if cond.MustNotContainsSuffixWord != nil {
		for _, v := range cond.MustNotContainsSuffixWord {
			if v != "" && text[textLen-len(v):] == v {
				return errors.New("the string must not contain suffix word: " + v)
			}
		}
	}

	var (
		iterateText bool
		onlyContains,
		mustContains,
		mustNotContains,
		mustBeFollowedBy,
		mustBeFollowedByPairs,
		mayContainsOnce asciis
	)

	if cond.OnlyContains != nil {
		iterateText = true
		for _, v := range cond.OnlyContains {
			onlyContains[v] = 1
		}
	}
	if cond.MustContains != nil {
		iterateText = true
		for _, v := range cond.MustContains {
			mustContains[v] = 1
		}
	}
	if cond.MustContainsOnce != nil {
		iterateText = true
		for _, v := range cond.MustContainsOnce {
			mustContains[v] = 1
			mayContainsOnce[v] = 1
		}
	}
	if cond.MustNotContains != nil {
		iterateText = true
		for _, v := range cond.MustNotContains {
			mustNotContains[v] = 1
		}
	}
	if cond.MayContainsOnce != nil {
		iterateText = true
		for _, v := range cond.MayContainsOnce {
			mayContainsOnce[v] = 1
		}
	}
	if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil {
		iterateText = true
		for _, v := range cond.MustBeFollowedBy[0] {
			mustBeFollowedBy[v] = 1
		}
		for _, v := range cond.MustBeFollowedBy[1] {
			mustBeFollowedByPairs[v] = 1
		}
	}

	if iterateText {
		for i, v := range text {
			if v > asciiMaxDecInt32 {
				return errors.New("the char: " + string(v) + ", is not a valid ascii format")
			}
			if cond.OnlyContains != nil && onlyContains[v] < 1 {
				return errors.New("the string cannot contain char: " + string(byte(v)))
			}
			if cond.MustNotContains != nil && mustNotContains[v] > 0 {
				return errors.New("the string must not contain char: " + string(byte(v)))
			}
			if (cond.MustContains != nil || cond.MustContainsOnce != nil) && mustContains[v] > 0 {
				mustContains[v] = 0
			}
			if (cond.MayContainsOnce != nil || cond.MustContainsOnce != nil) && mayContainsOnce[v] > 0 {
				if mayContainsOnce[v] > 1 {
					return errors.New("the char: " + string(byte(v)) + ", must be appeared once in the string")
				}
				mayContainsOnce[v] += 1
			}
			if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil && mustBeFollowedBy[v] > 0 {
				if i == 0 || (i+1) == textLen {
					return errors.New("the char: " + string(byte(v)) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if i > 0 && i < textLen && mustBeFollowedByPairs[text[i-1]] < 1 {
					return errors.New("the char: " + string(byte(v)) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if (i+1) < textLen && mustBeFollowedByPairs[text[i+1]] < 1 {
					return errors.New("the char: " + string(byte(v)) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
			}
		}
		if cond.MustContains != nil || cond.MustContainsOnce != nil {
			for b, v := range mustContains {
				if v > 0 {
					return errors.New("the string must contain char: " + string(rune(b)))
				}
			}
		}
	}

	if cond.MustContainsWord != nil {
		for _, v := range cond.MustContainsWord {
			if v != "" && !strings.Contains(text, v) {
				return errors.New("the string must contain word: " + v)
			}
		}
	}
	if cond.MustContainsWordOnce != nil {
		for _, v := range cond.MustContainsWordOnce {
			if v != "" && strings.Count(text, v) != 1 {
				return errors.New("the string must contain word: " + v + ", and it must be appeared once in the string")
			}
		}
	}
	if cond.MustNotContainsWord != nil {
		for _, v := range cond.MustNotContainsWord {
			if v != "" && strings.Contains(text, v) {
				return errors.New("the string must not contain word: " + v)
			}
		}
	}
	if cond.MayContainsWordOnce != nil {
		for _, v := range cond.MayContainsWordOnce {
			if v != "" && strings.Count(text, v) > 1 {
				return errors.New("the word: " + v + ", must be appeared once in the string")
			}
		}
	}

	return nil
}
