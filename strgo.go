package strgo

import (
	"errors"
	"fmt"
	"strings"
)

type StrGo interface {
	// Validate evaluates all conditions that have been set before and will return an error
	// if one of the conditions is not match.
	Validate() error

	// MinLength sets the minimum length condition for the string
	// and will be evaluated when Validate is called.
	MinLength(n int) StrGo

	// MaxLength sets the maximum length condition for the string
	// and will be evaluated when Validate is called.
	MaxLength(n int) StrGo

	// OnlyContainChars sets condition that the string can only contain characters
	// from the given `c` and will be evaluated when Validate is called.
	OnlyContainChars(c []string) StrGo

	// OnlyContainPrefixChars sets condition that the string can only contain prefix characters
	// from the given `pc` and will be evaluated when Validate is called.
	OnlyContainPrefixChars(pc []string) StrGo

	// OnlyContainSuffixChars sets condition that the string can only contain suffix characters
	// from the given `sc` and will be evaluated when Validate is called.
	OnlyContainSuffixChars(sc []string) StrGo

	// OnlyContainPrefixWords sets condition that the string can only contain prefix words
	// from the given `pw` and will be evaluated when Validate is called.
	OnlyContainPrefixWords(pw []string) StrGo

	// OnlyContainSuffixWords sets condition that the string can only contain suffix words
	// from the given `sw` and will be evaluated when Validate is called.
	OnlyContainSuffixWords(sw []string) StrGo

	// MustContainChars sets condition that the string must contain characters
	// from the given `c` and will be evaluated when Validate is called.
	MustContainChars(c []string) StrGo

	// MustContainWords sets condition that the string must contain words
	// from the given `w` and will be evaluated when Validate is called.
	MustContainWords(w []string) StrGo

	// MustContainCharsOnce sets condition that the string must contain characters
	// from the given `c` and each of them must be appeared once in the string.
	// It will be evaluated when Validate is called.
	MustContainCharsOnce(c []string) StrGo

	// MustContainWordsOnce sets condition that the string must contain words
	// from the given `w` and each of them must be appeared once in the string.
	// It will be evaluated when Validate is called.
	MustContainWordsOnce(w []string) StrGo

	// MustNotContainChars sets condition that the string must not contain characters
	// from the given `c` and will be evaluated when Validate is called.
	MustNotContainChars(c []string) StrGo

	// MustNotContainWords sets condition that the string must not contain words
	// from the given `w` and will be evaluated when Validate is called.
	MustNotContainWords(w []string) StrGo

	// MustNotContainPrefixChars sets condition that the string must not contain prefix characters
	// from the given `pc` and will be evaluated when Validate is called.
	MustNotContainPrefixChars(pc []string) StrGo

	// MustNotContainSuffixChars sets condition that the string must not contain suffix characters
	// from the given `sc` and will be evaluated when Validate is called.
	MustNotContainSuffixChars(sc []string) StrGo

	// MustNotContainPrefixWords sets condition that the string must not contain prefix words
	// from the given `pw` and will be evaluated when Validate is called.
	MustNotContainPrefixWords(pw []string) StrGo

	// MustNotContainSuffixWords sets condition that the string must not contain suffix words
	// from the given `sw` and will be evaluated when Validate is called.
	MustNotContainSuffixWords(sw []string) StrGo

	// MustBeFollowedByChars sets condition that the given characters `c` in the string must be followed by
	// at least one of the characters from the given `fc` and will be evaluated when Validate is called.
	MustBeFollowedByChars(c, fc []string) StrGo

	// MayContainCharsOnce sets condition that the string may contain characters
	// from the given `c`, but they must be appeared once in the string.
	// It will be evaluated when Validate is called.
	MayContainCharsOnce(c []string) StrGo

	// MayContainWordsOnce sets condition that the string may contain words
	// from the given `w`, but they must be appeared once in the string.
	// It will be evaluated when Validate is called.
	MayContainWordsOnce(w []string) StrGo
}

type strGo struct {
	str       string
	length    int
	minLength int
	maxLength int
	chars     map[string]map[string]bool
	words     map[string]map[string]bool
	conds     map[string][]string
}

// New generates new strgo validator.
func New(str string) StrGo {
	return &strGo{
		str:       str,
		length:    len(str),
		maxLength: len(str),
		chars:     make(map[string]map[string]bool),
		words:     make(map[string]map[string]bool),
		conds:     make(map[string][]string),
	}
}

func (str *strGo) MinLength(n int) StrGo {
	str.minLength = n
	return str
}

func (str *strGo) MaxLength(n int) StrGo {
	str.maxLength = n
	return str
}

func (str *strGo) OnlyContainChars(c []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, c, "OnlyContainChar")
	return str
}

func (str *strGo) OnlyContainPrefixChars(pc []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, pc, "OnlyContainPrefixChars")
	return str
}

func (str *strGo) OnlyContainSuffixChars(sc []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, sc, "OnlyContainSuffixChars")
	return str
}

func (str *strGo) OnlyContainPrefixWords(pw []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, pw, "OnlyContainPrefixWords")
	return str
}

func (str *strGo) OnlyContainSuffixWords(sw []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, sw, "OnlyContainSuffixWords")
	return str
}

func (str *strGo) MustContainChars(c []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, c, "MustContainChars")
	return str
}

func (str *strGo) MustContainWords(w []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, w, "MustContainWords")
	return str
}

func (str *strGo) MustContainCharsOnce(c []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, c, "MustContainChars")
	str.chars, str.conds = setChars(str.chars, str.conds, c, "MayContainCharsOnce")
	return str
}

func (str *strGo) MustContainWordsOnce(w []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, w, "MustContainWords")
	str.words, str.conds = setWords(str.words, str.conds, w, "MayContainWordsOnce")
	return str
}

func (str *strGo) MustNotContainChars(c []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, c, "MustNotContainChars")
	return str
}

func (str *strGo) MustNotContainWords(w []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, w, "MustNotContainWords")
	return str
}

func (str *strGo) MustNotContainPrefixChars(pc []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, pc, "MustNotContainPrefixChars")
	return str
}

func (str *strGo) MustNotContainSuffixChars(sc []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, sc, "MustNotContainSuffixChars")
	return str
}

func (str *strGo) MustNotContainPrefixWords(pw []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, pw, "MustNotContainPrefixWords")
	return str
}

func (str *strGo) MustNotContainSuffixWords(sw []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, sw, "MustNotContainSuffixWords")
	return str
}

func (str *strGo) MustBeFollowedByChars(c, fc []string) StrGo {
	str.chars, str.conds = setMustBeFollowedByChars(str.chars, str.conds, c, fc, "MustBeFollowedByChars")
	return str
}

func (str *strGo) MayContainCharsOnce(c []string) StrGo {
	str.chars, str.conds = setChars(str.chars, str.conds, c, "MayContainCharsOnce")
	return str
}

func (str *strGo) MayContainWordsOnce(w []string) StrGo {
	str.words, str.conds = setWords(str.words, str.conds, w, "MayContainWordsOnce")
	return str
}

func (str *strGo) Validate() error {
	if str.length < 1 {
		return errors.New(ErrEmpty)
	}

	if str.length < str.minLength {
		return errors.New(fmt.Sprintf(ErrMinLength, str.minLength))
	}

	if str.length > str.maxLength {
		return errors.New(fmt.Sprintf(ErrMaxLength, str.maxLength))
	}

	for s, v := range str.words {
		if _, ok := str.conds["OnlyContainPrefixWords"]; ok {
			if _, ok := v["OnlyContainPrefixWords"]; ok && strings.HasPrefix(str.str, s) {
				delete(str.conds, "OnlyContainPrefixWords")
			}
		}

		if _, ok := str.conds["OnlyContainSuffixWords"]; ok {
			if _, ok := v["OnlyContainSuffixWords"]; ok && strings.HasSuffix(str.str, s) {
				delete(str.conds, "OnlyContainSuffixWords")
			}
		}

		if _, ok := v["MustContainWords"]; ok && !strings.Contains(str.str, s) {
			return errors.New(fmt.Sprintf(ErrMustContainWords, s))
		}

		if _, ok := v["MayContainWordsOnce"]; ok && strings.Count(str.str, s) > 1 {
			return errors.New(fmt.Sprintf(ErrMayContainWordsOnce, s))
		}

		if _, ok := v["MustNotContainWords"]; ok && strings.Contains(str.str, s) {
			return errors.New(fmt.Sprintf(ErrMustNotContainWords, s))
		}

		if _, ok := v["MustNotContainPrefixWords"]; ok && strings.HasPrefix(str.str, s) {
			return errors.New(fmt.Sprintf(ErrMustNotContainPrefixWords, s))
		}

		if _, ok := v["MustNotContainSuffixWords"]; ok && strings.HasSuffix(str.str, s) {
			return errors.New(fmt.Sprintf(ErrMustNotContainSuffixWords, s))
		}
	}

	if v, ok := str.conds["OnlyContainPrefixWords"]; ok {
		return errors.New(fmt.Sprintf(ErrOnlyContainPrefixWords, v))
	}

	if v, ok := str.conds["OnlyContainSuffixWords"]; ok {
		return errors.New(fmt.Sprintf(ErrOnlyContainSuffixWords, v))
	}

	if v, ok := str.chars["OnlyContainPrefixChars"]; ok {
		p := string(str.str[0])
		if _, ok := v[p]; !ok {
			return errors.New(fmt.Sprintf(ErrOnlyContainPrefixChars, p))
		}
	}

	if v, ok := str.chars["OnlyContainSuffixChars"]; ok {
		s := string(str.str[str.length-1])
		if _, ok := v[s]; !ok {
			return errors.New(fmt.Sprintf(ErrOnlyContainSuffixChars, s))
		}
	}

	if v, ok := str.chars["MustNotContainPrefixChars"]; ok {
		p := string(str.str[0])
		if _, ok := v[p]; ok {
			return errors.New(fmt.Sprintf(ErrMustNotContainPrefixChars, p))
		}
	}

	if v, ok := str.chars["MustNotContainSuffixChars"]; ok {
		s := string(str.str[str.length-1])
		if _, ok := v[s]; ok {
			return errors.New(fmt.Sprintf(ErrMustNotContainSuffixChars, s))
		}
	}

	for i, v := range str.str {
		s := string(v)

		if v, ok := str.chars["OnlyContainChar"]; ok && !v[s] {
			return errors.New(fmt.Sprintf(ErrOnlyContainChars, s))
		}

		if v, ok := str.chars["MustNotContainChars"]; ok && v[s] {
			return errors.New(fmt.Sprintf(ErrMustNotContainChars, s))
		}

		if v, ok := str.chars["MustBeFollowedByChars"+s]; ok {
			b := s
			a := s
			if i > 0 && i < str.length {
				b = string(str.str[i-1])
			}
			if (i + 1) < str.length {
				a = string(str.str[i+1])
			}
			if !v[b] || !v[a] {
				return errors.New(fmt.Sprintf(ErrMustBeFollowedByChars, s, str.conds["MustBeFollowedByChars"+s]))
			}
		}

		if v, ok := str.chars["MayContainCharsOnce"]; ok {
			if _, ok := v[s]; ok {
				if !v[s] {
					return errors.New(fmt.Sprintf(ErrMayContainCharsOnce, s))
				}
				v[s] = false
			}
		}

		if v, ok := str.chars["MustContainChars"]; ok && v[s] {
			delete(str.chars["MustContainChars"], s)
		}
	}

	for s, _ := range str.chars["MustContainChars"] {
		return errors.New(fmt.Sprintf(ErrMustContainChars, s))
	}

	return nil
}

func setChars(m map[string]map[string]bool, cm map[string][]string, c []string, cond string) (map[string]map[string]bool, map[string][]string) {
	for _, v := range c {
		if len(v) == 1 {
			if _, ok := m[cond]; !ok {
				m[cond] = make(map[string]bool)
			}
			m[cond][v] = true
			cm[cond] = append(cm[cond], v)
		}
	}
	return m, cm
}

func setMustBeFollowedByChars(m map[string]map[string]bool, cm map[string][]string, c []string, fc []string, cond string) (map[string]map[string]bool, map[string][]string) {
	for _, v := range c {
		if len(v) == 1 {
			if _, ok := m[cond+v]; !ok {
				m[cond+v] = make(map[string]bool)
			}
			for _, fv := range fc {
				m[cond+v][fv] = true
				cm[cond+v] = append(cm[cond+v], fv)
			}
		}
	}
	return m, cm
}

func setWords(m map[string]map[string]bool, cm map[string][]string, w []string, cond string) (map[string]map[string]bool, map[string][]string) {
	for _, v := range w {
		if _, ok := m[v]; !ok {
			m[v] = make(map[string]bool)
		}
		m[v][cond] = true
		cm[cond] = append(cm[cond], v)
	}
	return m, cm
}
