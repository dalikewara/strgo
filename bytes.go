package strgo

import (
	"bytes"
	"errors"
	"strconv"
)

type BytesCondition struct {
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

// Bytes matches the bytes based on the given condition.
// If one doesn't match, it will return an error.
func Bytes(text []byte, cond *BytesCondition) error {
	bytesLen := len(text)

	if bytesLen < 1 {
		return errors.New("the bytes is empty")
	}
	if cond.MinLength > 0 && bytesLen < cond.MinLength {
		return errors.New("the bytes length cannot be less than " + strconv.Itoa(cond.MinLength))
	}
	if cond.MaxLength > 0 && bytesLen > cond.MaxLength {
		return errors.New("the bytes length cannot be more than " + strconv.Itoa(cond.MaxLength))
	}

	if cond.OnlyContainsPrefix != nil && bytes.IndexByte(cond.OnlyContainsPrefix, text[0]) < 0 {
		return errors.New("the bytes cannot contain prefix char: " + string(text[0]))
	}
	if cond.OnlyContainsSuffix != nil && bytes.IndexByte(cond.OnlyContainsSuffix, text[bytesLen-1]) < 0 {
		return errors.New("the bytes cannot contain suffix char: " + string(text[bytesLen-1]))
	}
	if cond.MustNotContainsPrefix != nil && bytes.IndexByte(cond.MustNotContainsPrefix, text[0]) >= 0 {
		return errors.New("the bytes must not contain prefix: " + string(text[0]))
	}
	if cond.MustNotContainsSuffix != nil && bytes.IndexByte(cond.MustNotContainsSuffix, text[bytesLen-1]) >= 0 {
		return errors.New("the bytes must not contain suffix: " + string(text[bytesLen-1]))
	}

	var (
		iterateText bool
		onlyContains,
		mustContains,
		mustNotContains,
		mustBeFollowedBy,
		mustBeFollowedByPairs,
		mayContainsOnce [255]byte
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
			if cond.OnlyContains != nil && onlyContains[v] < 1 {
				return errors.New("the bytes cannot contain char: " + string(byte(v)))
			}
			if cond.MustNotContains != nil && mustNotContains[v] > 0 {
				return errors.New("the bytes must not contain char: " + string(byte(v)))
			}
			if (cond.MustContains != nil || cond.MustContainsOnce != nil) && mustContains[v] > 0 {
				mustContains[v] = 0
			}
			if (cond.MayContainsOnce != nil || cond.MustContainsOnce != nil) && mayContainsOnce[v] > 0 {
				if mayContainsOnce[v] > 1 {
					return errors.New("the char: " + string(v) + ", must be appeared once in the bytes")
				}
				mayContainsOnce[v] += 1
			}
			if cond.MustBeFollowedBy[0] != nil && cond.MustBeFollowedBy[1] != nil && mustBeFollowedBy[v] > 0 {
				if i == 0 || (i+1) == bytesLen {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if i > 0 && i < bytesLen && mustBeFollowedByPairs[text[i-1]] < 1 {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
				if (i+1) < bytesLen && mustBeFollowedByPairs[text[i+1]] < 1 {
					return errors.New("the char: " + string(v) + ", must be followed with at least one of these characters: " + string(cond.MustBeFollowedBy[1]))
				}
			}
		}
		if cond.MustContains != nil || cond.MustContainsOnce != nil {
			for b, v := range mustContains {
				if v > 0 {
					return errors.New("the bytes must contain char: " + string(byte(b)))
				}
			}
		}
	}

	return nil
}
