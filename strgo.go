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
	str                       string
	length                    int
	minLength                 int
	maxLength                 int
	onlyContainChar           map[string]string
	onlyContainPrefixChars    map[string]string
	onlyContainSuffixChars    map[string]string
	mustContainChars          map[string]string
	mustContainWords          map[string]string
	mustContainCharsOnce      map[string]string
	mustContainCharsOnceF     map[string]string
	mustContainWordsOnce      map[string]string
	mustNotContainWords       map[string]string
	mustNotContainChars       map[string]string
	mustNotContainPrefixChars map[string]string
	mustNotContainSuffixChars map[string]string
	mustBeFollowedByChars     map[string]string
	mustBeFollowedByCharsF    map[string]string
	mustBeFollowedByCharsFC   []string
	mayContainCharsOnce       map[string]string
	mayContainWordsOnce       map[string]string
}

// New generates new strgo validator.
func New(str string) StrGo {
	return &strGo{
		str:                       str,
		length:                    len(str),
		minLength:                 0,
		maxLength:                 len(str),
		onlyContainChar:           make(map[string]string),
		onlyContainPrefixChars:    make(map[string]string),
		onlyContainSuffixChars:    make(map[string]string),
		mustContainChars:          make(map[string]string),
		mustContainWords:          make(map[string]string),
		mustContainCharsOnce:      make(map[string]string),
		mustContainCharsOnceF:     make(map[string]string),
		mustContainWordsOnce:      make(map[string]string),
		mustNotContainWords:       make(map[string]string),
		mustNotContainChars:       make(map[string]string),
		mustNotContainPrefixChars: make(map[string]string),
		mustNotContainSuffixChars: make(map[string]string),
		mustBeFollowedByChars:     make(map[string]string),
		mustBeFollowedByCharsF:    make(map[string]string),
		mustBeFollowedByCharsFC:   []string{},
		mayContainCharsOnce:       make(map[string]string),
		mayContainWordsOnce:       make(map[string]string),
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
	str.onlyContainChar = setChars(str.onlyContainChar, c)
	return str
}

func (str *strGo) OnlyContainPrefixChars(pc []string) StrGo {
	str.onlyContainPrefixChars = setChars(str.onlyContainPrefixChars, pc)
	return str
}

func (str *strGo) OnlyContainSuffixChars(sc []string) StrGo {
	str.onlyContainSuffixChars = setChars(str.onlyContainSuffixChars, sc)
	return str
}

func (str *strGo) MustContainChars(c []string) StrGo {
	str.mustContainChars = setChars(str.mustContainChars, c)
	return str
}

func (str *strGo) MustContainWords(w []string) StrGo {
	str.mustContainWords = setWords(str.mustContainWords, w)
	return str
}

func (str *strGo) MustContainCharsOnce(c []string) StrGo {
	str.mustContainCharsOnce = setChars(str.mustContainCharsOnce, c)
	str.mustContainCharsOnceF = setChars(str.mustContainCharsOnceF, c)
	return str
}

func (str *strGo) MustContainWordsOnce(w []string) StrGo {
	str.mustContainWordsOnce = setWords(str.mustContainWordsOnce, w)
	return str
}

func (str *strGo) MustNotContainChars(c []string) StrGo {
	str.mustNotContainChars = setChars(str.mustNotContainChars, c)
	return str
}

func (str *strGo) MustNotContainWords(w []string) StrGo {
	str.mustNotContainWords = setWords(str.mustNotContainWords, w)
	return str
}

func (str *strGo) MustNotContainPrefixChars(pc []string) StrGo {
	str.mustNotContainPrefixChars = setChars(str.mustNotContainPrefixChars, pc)
	return str
}

func (str *strGo) MustNotContainSuffixChars(sc []string) StrGo {
	str.mustNotContainSuffixChars = setChars(str.mustNotContainSuffixChars, sc)
	return str
}

func (str *strGo) MustBeFollowedByChars(c, fc []string) StrGo {
	str.mustBeFollowedByChars = setChars(str.mustBeFollowedByChars, c)
	str.mustBeFollowedByCharsF = setChars(str.mustBeFollowedByCharsF, fc)
	str.mustBeFollowedByCharsFC = fc
	return str
}

func (str *strGo) MayContainCharsOnce(c []string) StrGo {
	str.mayContainCharsOnce = setChars(str.mayContainCharsOnce, c)
	return str
}

func (str *strGo) MayContainWordsOnce(w []string) StrGo {
	str.mayContainWordsOnce = setWords(str.mayContainWordsOnce, w)
	return str
}

func (str *strGo) Validate() error {
	var err error

	if err = str.validateMinLength(); err != nil {
		return err
	}
	if err = str.validateMaxLength(); err != nil {
		return err
	}
	if err = str.validateOnlyContainPrefixChars(); err != nil {
		return err
	}
	if err = str.validateMustNotContainPrefixChars(); err != nil {
		return err
	}
	if err = str.validateOnlyContainSuffixChars(); err != nil {
		return err
	}
	if err = str.validateMustNotContainSuffixChars(); err != nil {
		return err
	}
	if err = str.validateMustContainWords(); err != nil {
		return err
	}
	if err = str.validateMustContainWordsOnce(); err != nil {
		return err
	}
	if err = str.validateMustNotContainWords(); err != nil {
		return err
	}
	if err = str.validateMayContainWordsOnce(); err != nil {
		return err
	}

	for i, v := range str.str {
		s := string(v)

		if err = str.validateOnlyContainChars(s); err != nil {
			return err
		}
		if err = str.validateMustNotContainChars(s); err != nil {
			return err
		}
		if err = str.validateMustContainCharsOnce(s); err != nil {
			return err
		}
		if err = str.validateMustBeFollowedByChars(i, s); err != nil {
			return err
		}
		if err = str.validateMayContainCharsOnce(s); err != nil {
			return err
		}

		str.evaluateMustContainChars(s)
	}

	if err = str.validateMustContainChars(); err != nil {
		return err
	}

	return nil
}

func (str *strGo) validateMinLength() error {
	if str.length < str.minLength {
		return errors.New(fmt.Sprintf(ErrMinLength, str.minLength))
	}
	return nil
}

func (str *strGo) validateMaxLength() error {
	if str.length > str.maxLength {
		return errors.New(fmt.Sprintf(ErrMaxLength, str.maxLength))
	}
	return nil
}

func (str *strGo) validateOnlyContainChars(s string) error {
	if _, ok := str.onlyContainChar[s]; !ok && len(str.onlyContainChar) > 0 {
		return errors.New(fmt.Sprintf(ErrOnlyContainChars, s))
	}
	return nil
}

func (str *strGo) validateOnlyContainPrefixChars() error {
	if str.length > 0 && len(str.onlyContainPrefixChars) > 0 {
		p := string(str.str[0])
		if _, ok := str.onlyContainPrefixChars[p]; !ok {
			return errors.New(fmt.Sprintf(ErrOnlyContainPrefixChars, p))
		}
	}
	return nil
}

func (str *strGo) validateOnlyContainSuffixChars() error {
	if str.length > 0 && len(str.onlyContainSuffixChars) > 0 {
		s := string(str.str[str.length-1])
		if _, ok := str.onlyContainSuffixChars[s]; !ok {
			return errors.New(fmt.Sprintf(ErrOnlyContainSuffixChars, s))
		}
	}
	return nil
}

func (str *strGo) evaluateMustContainChars(s string) {
	if _, ok := str.mustContainChars[s]; ok && len(str.mustContainChars) > 0 {
		delete(str.mustContainChars, s)
	}
	if _, ok := str.mustContainCharsOnce[s]; ok && len(str.mustContainCharsOnce) > 0 {
		delete(str.mustContainCharsOnce, s)
	}
}

func (str *strGo) validateMustContainChars() error {
	for _, v := range str.mustContainChars {
		return errors.New(fmt.Sprintf(ErrMustContainChars, v))
	}
	for _, v := range str.mustContainCharsOnce {
		return errors.New(fmt.Sprintf(ErrMustContainChars, v))
	}
	return nil
}

func (str *strGo) validateMustContainWords() error {
	for _, v := range str.mustContainWords {
		if !strings.Contains(str.str, v) {
			return errors.New(fmt.Sprintf(ErrMustContainWords, v))
		}
	}
	return nil
}

func (str *strGo) validateMustContainCharsOnce(s string) error {
	if v, ok := str.mustContainCharsOnceF[s]; ok && len(str.mustContainCharsOnceF) > 0 {
		if v == "" {
			return errors.New(fmt.Sprintf(ErrMustContainCharsOnce, s))
		}
		str.mustContainCharsOnceF[s] = ""
	}
	return nil
}

func (str *strGo) validateMustContainWordsOnce() error {
	for _, v := range str.mustContainWordsOnce {
		if !strings.Contains(str.str, v) {
			return errors.New(fmt.Sprintf(ErrMustContainWords, v))
		}
		if strings.Count(str.str, v) > 1 {
			return errors.New(fmt.Sprintf(ErrMustContainWordsOnce, v))
		}
	}
	return nil
}

func (str *strGo) validateMustNotContainChars(s string) error {
	if _, ok := str.mustNotContainChars[s]; ok && len(str.mustNotContainChars) > 0 {
		return errors.New(fmt.Sprintf(ErrMustNotContainChars, s))
	}
	return nil
}

func (str *strGo) validateMustNotContainWords() error {
	for _, v := range str.mustNotContainWords {
		if strings.Contains(str.str, v) {
			return errors.New(fmt.Sprintf(ErrMustNotContainWords, v))
		}
	}
	return nil
}

func (str *strGo) validateMustNotContainPrefixChars() error {
	if str.length > 0 && len(str.mustNotContainPrefixChars) > 0 {
		p := string(str.str[0])
		if _, ok := str.mustNotContainPrefixChars[p]; ok {
			return errors.New(fmt.Sprintf(ErrMustNotContainPrefixChars, p))
		}
	}
	return nil
}

func (str *strGo) validateMustNotContainSuffixChars() error {
	if str.length > 0 && len(str.mustNotContainSuffixChars) > 0 {
		s := string(str.str[str.length-1])
		if _, ok := str.mustNotContainSuffixChars[s]; ok {
			return errors.New(fmt.Sprintf(ErrMustNotContainSuffixChars, s))
		}
	}
	return nil
}

func (str *strGo) validateMustBeFollowedByChars(i int, s string) error {
	if _, ok := str.mustBeFollowedByChars[s]; ok && len(str.mustBeFollowedByChars) > 0 && i > 0 {
		b := str.str[i-1]
		a := str.str[i]
		if (i + 1) <= str.length {
			a = str.str[i+1]
		}
		if _, okB := str.mustBeFollowedByCharsF[string(b)]; !okB {
			return errors.New(fmt.Sprintf(ErrMustBeFollowedByChars, s, str.mustBeFollowedByCharsFC))
		}
		if _, okA := str.mustBeFollowedByCharsF[string(a)]; !okA {
			return errors.New(fmt.Sprintf(ErrMustBeFollowedByChars, s, str.mustBeFollowedByCharsFC))
		}
	}
	return nil
}

func (str *strGo) validateMayContainCharsOnce(s string) error {
	if v, ok := str.mayContainCharsOnce[s]; ok && len(str.mayContainCharsOnce) > 0 {
		if v == "" {
			return errors.New(fmt.Sprintf(ErrMayContainCharsOnce, s))
		}
		str.mayContainCharsOnce[s] = ""
	}
	return nil
}

func (str *strGo) validateMayContainWordsOnce() error {
	for _, v := range str.mayContainWordsOnce {
		if strings.Contains(str.str, v) && strings.Count(str.str, v) > 1 {
			return errors.New(fmt.Sprintf(ErrMayContainWordsOnce, v))
		}
	}
	return nil
}

func setChars(m map[string]string, c []string) map[string]string {
	for _, v := range c {
		if len(v) == 1 {
			if _, ok := m[v]; !ok {
				m[v] = v
			}
		}
	}
	return m
}

func setWords(m map[string]string, w []string) map[string]string {
	for _, v := range w {
		if _, ok := m[v]; !ok {
			m[v] = v
		}
	}
	return m
}
