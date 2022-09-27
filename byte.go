package strgo

import (
	"bytes"
	"errors"
	"strconv"
)

const asciiMaxDecInt32 = int32(254)

type asciis [255]byte

type ByteCondition struct {
	MinLength             int
	MaxLength             int
	OnlyContains          []byte
	OnlyContainsPrefix    []byte
	OnlyContainsSuffix    []byte
	MustContains          []byte
	MustContainsOnce      []byte
	MustNotContains       []byte
	MustNotContainsPrefix []byte
	MustNotContainsSuffix []byte
	MayContainsOnce       []byte
	MustBeFollowedBy      [2][]byte
}

// Byte matches the string based on the ByteCondition.
// If one doesn't match, it will return an error.
// This function can only validate ASCII characters.
func Byte(text string, cond *ByteCondition) error {
	if text == "" {
		return errors.New("the string is empty")
	}

	if cond.MinLength > 0 && len(text) < cond.MinLength {
		return errors.New("the string length cannot be less than " + strconv.Itoa(cond.MinLength))
	}
	if cond.MaxLength > 0 && len(text) > cond.MaxLength {
		return errors.New("the string length cannot be more than " + strconv.Itoa(cond.MaxLength))
	}

	runes := []rune(text)
	runesLen := len(runes)

	if cond.OnlyContainsPrefix != nil {
		if runes[0] > asciiMaxDecInt32 {
			return errors.New("the char: " + string(runes[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexRune(cond.OnlyContainsPrefix, runes[0]) < 0 {
			return errors.New("the string cannot contain prefix char: " + string(runes[0]))
		}
	}
	if cond.OnlyContainsSuffix != nil {
		if runes[0] > asciiMaxDecInt32 {
			return errors.New("the char: " + string(runes[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexRune(cond.OnlyContainsSuffix, runes[runesLen-1]) < 0 {
			return errors.New("the string cannot contain suffix char: " + string(runes[runesLen-1]))
		}
	}
	if cond.MustNotContainsPrefix != nil {
		if runes[0] > asciiMaxDecInt32 {
			return errors.New("the char: " + string(runes[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexRune(cond.MustNotContainsPrefix, runes[0]) >= 0 {
			return errors.New("the string must not contain prefix: " + string(runes[0]))
		}
	}
	if cond.MustNotContainsSuffix != nil {
		if runes[0] > asciiMaxDecInt32 {
			return errors.New("the char: " + string(runes[0]) + ", is not a valid ascii format")
		}
		if bytes.IndexRune(cond.MustNotContainsSuffix, runes[runesLen-1]) >= 0 {
			return errors.New("the string must not contain suffix: " + string(runes[runesLen-1]))
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
		for i, v := range runes {
			if v > asciiMaxDecInt32 {
				return errors.New("the char: " + string(v) + ", is not a valid ascii format")
			}
			if cond.OnlyContains != nil && onlyContains[v] < 1 {
				return errors.New("the string cannot contain char: " + string(v))
			}
			if cond.MustNotContains != nil && mustNotContains[v] > 0 {
				return errors.New("the string must not contain char: " + string(v))
			}
			if (cond.MustContains != nil || cond.MustContainsOnce != nil) && mustContains[v] > 0 {
				mustContains[v] = 0
			}
			if (cond.MayContainsOnce != nil || cond.MustContainsOnce != nil) && mayContainsOnce[v] > 0 {
				if mayContainsOnce[v] > 1 {
					return errors.New("the char: " + string(v) + ", must be appeared once in the string")
				}
				mayContainsOnce[v] += 1
			}
			if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil && mustBeFollowedBy[v] > 0 {
				if i == 0 || (i+1) == runesLen {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if i > 0 && i < runesLen && mustBeFollowedByPairs[runes[i-1]] < 1 {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if (i+1) < runesLen && mustBeFollowedByPairs[runes[i+1]] < 1 {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
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

	return nil
}
