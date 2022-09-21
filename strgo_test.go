package strgo_test

import (
	"fmt"
	"github.com/dalikewara/strgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	str := strgo.New("johndoe")
	assert.Implements(t, (*strgo.StrGo)(nil), str)
}

func TestStrGo_Validate(t *testing.T) {
	err := strgo.New("johndoe").Validate()
	assert.Nil(t, err)
	err = strgo.New("").Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, strgo.ErrEmpty)
}

func TestStrGo_MinLength(t *testing.T) {
	err := strgo.New("johndoe").MinLength(7).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MinLength(8).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMinLength, 8))
}

func TestStrGo_MaxLength(t *testing.T) {
	err := strgo.New("johndoe").MaxLength(7).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MaxLength(6).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMaxLength, 6))
}

func TestStrGo_OnlyContainChars(t *testing.T) {
	err := strgo.New("johndoe").OnlyContainChars([]string{"j", "o", "h", "n", "d", "e"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").OnlyContainChars([]string{"j", "o", "h", "n", "d"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrOnlyContainChars, "e"))
}

func TestStrGo_OnlyContainPrefixChars(t *testing.T) {
	err := strgo.New("johndoe").OnlyContainPrefixChars([]string{"o", "j"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").OnlyContainPrefixChars([]string{"o", "k"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrOnlyContainPrefixChars, "j"))
}

func TestStrGo_OnlyContainSuffixChars(t *testing.T) {
	err := strgo.New("johndoe").OnlyContainSuffixChars([]string{"o", "e"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").OnlyContainSuffixChars([]string{"o", "k"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrOnlyContainSuffixChars, "e"))
}

func TestStrGo_OnlyContainPrefixWords(t *testing.T) {
	err := strgo.New("johndoe").OnlyContainPrefixWords([]string{"hn", "joh"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").OnlyContainPrefixWords([]string{"noh", "koh"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrOnlyContainPrefixWords, "[noh koh]"))
}

func TestStrGo_OnlyContainSuffixWords(t *testing.T) {
	err := strgo.New("johndoe").OnlyContainSuffixWords([]string{"eo", "oe"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").OnlyContainSuffixWords([]string{"ne", "ho"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrOnlyContainSuffixWords, "[ne ho]"))
}

func TestStrGo_MustContainChars(t *testing.T) {
	err := strgo.New("johndoe").MustContainChars([]string{"n", "h"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustContainChars([]string{"n", "k"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustContainChars, "k"))
}

func TestStrGo_MustContainWords(t *testing.T) {
	err := strgo.New("johndoe").MustContainWords([]string{"john", "hnd"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustContainWords([]string{"ohn", "de"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustContainWords, "de"))
}

func TestStrGo_MustContainCharsOnce(t *testing.T) {
	err := strgo.New("johndoe").MustContainCharsOnce([]string{"n", "h"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustContainCharsOnce([]string{"n", "k"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustContainChars, "k"))
	err = strgo.New("johndoe").MustContainCharsOnce([]string{"n", "o"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMayContainCharsOnce, "o"))
}

func TestStrGo_MustContainWordsOnce(t *testing.T) {
	err := strgo.New("johndoe").MustContainWordsOnce([]string{"john", "hnd"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustContainWordsOnce([]string{"ohn", "de"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustContainWords, "de"))
	err = strgo.New("johndoedoe").MustContainWordsOnce([]string{"ohn", "doe"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMayContainWordsOnce, "doe"))
}

func TestStrGo_MustNotContainChars(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainChars([]string{"k", "r"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainChars([]string{"k", "e"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainChars, "e"))
}

func TestStrGo_MustNotContainWords(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainWords([]string{"dor", "johk"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainWords([]string{"hng", "ohn"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainWords, "ohn"))
}

func TestStrGo_MustNotContainPrefixChars(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainPrefixChars([]string{"o", "h"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainPrefixChars([]string{"o", "j"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainPrefixChars, "j"))
}

func TestStrGo_MustNotContainSuffixChars(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainSuffixChars([]string{"o", "h"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainSuffixChars([]string{"o", "e"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainSuffixChars, "e"))
}

func TestStrGo_MustNotContainPrefixWords(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainPrefixWords([]string{"hn", "koh"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainPrefixWords([]string{"noh", "joh"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainPrefixWords, "joh"))
}

func TestStrGo_MustNotContainSuffixWords(t *testing.T) {
	err := strgo.New("johndoe").MustNotContainSuffixWords([]string{"eo", "ode"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustNotContainSuffixWords([]string{"doe", "joh"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustNotContainSuffixWords, "doe"))
}

func TestStrGo_MustBeFollowedByChars(t *testing.T) {
	err := strgo.New("johndoe").MustBeFollowedByChars([]string{"h", "d"}, []string{"m", "n", "o"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MustBeFollowedByChars([]string{"h", "o"}, []string{"d", "k", "l"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustBeFollowedByChars, "o", "[d k l]"))
	err = strgo.New("johndoe").MustBeFollowedByChars([]string{"h", "o"}, []string{"d", "k", "j"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMustBeFollowedByChars, "o", "[d k j]"))
}

func TestStrGo_MayContainCharsOnce(t *testing.T) {
	err := strgo.New("johndoe").MayContainCharsOnce([]string{"k", "r"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoe").MayContainCharsOnce([]string{"k", "o"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMayContainCharsOnce, "o"))
}

func TestStrGo_MayContainWordsOnce(t *testing.T) {
	err := strgo.New("johndoe").MayContainWordsOnce([]string{"khn", "rhn"}).Validate()
	assert.Nil(t, err)
	err = strgo.New("johndoedoe").MayContainWordsOnce([]string{"khn", "doe"}).Validate()
	assert.NotNil(t, err)
	assert.EqualError(t, err, fmt.Sprintf(strgo.ErrMayContainWordsOnce, "doe"))
}

func TestUsername(t *testing.T) {
	/*
		- `username` can only contain alphanumeric characters, underscores and periods
		- its length must be greater than 2 and not more than 20
		- special character must be followed by at least one alphanumeric character
		- prefix and suffix cannot be a special character
		- allowed special characters must be appeared once in the string
	*/
	var validate = func(username string) error {
		return strgo.New(username).
			OnlyContainChars(strgo.ALPHANUMERIC).
			MinLength(3).
			MaxLength(20).
			OnlyContainChars([]string{"_", "."}).
			MustBeFollowedByChars([]string{"_", "."}, strgo.ALPHANUMERIC).
			MustNotContainPrefixChars([]string{"_", "."}).
			MustNotContainSuffixChars([]string{"_", "."}).
			MayContainCharsOnce([]string{"_", "."}).
			Validate()
	}
	err := validate("johndoe")
	assert.Nil(t, err)
	err = validate("john_doe")
	assert.Nil(t, err)
	err = validate("john.doe")
	assert.Nil(t, err)
	err = validate("johndoe123")
	assert.Nil(t, err)
	err = validate("john.doe123")
	assert.Nil(t, err)
	err = validate("john_doe123")
	assert.Nil(t, err)
	err = validate("john_doe.123")
	assert.Nil(t, err)
	err = validate("_johndoe")
	assert.NotNil(t, err)
	err = validate("__johndoe")
	assert.NotNil(t, err)
	err = validate(".johndoe")
	assert.NotNil(t, err)
	err = validate("..johndoe")
	assert.NotNil(t, err)
	err = validate("johndoe_")
	assert.NotNil(t, err)
	err = validate("johndoe__")
	assert.NotNil(t, err)
	err = validate("johndoe.")
	assert.NotNil(t, err)
	err = validate("johndoe..")
	assert.NotNil(t, err)
	err = validate("john__doe")
	assert.NotNil(t, err)
	err = validate("john_.doe")
	assert.NotNil(t, err)
	err = validate("john..doe")
	assert.NotNil(t, err)
	err = validate("john._doe")
	assert.NotNil(t, err)
	err = validate("john@doe")
	assert.NotNil(t, err)
	err = validate("@johndoe")
	assert.NotNil(t, err)
	err = validate("johndoe@")
	assert.NotNil(t, err)
	err = validate("joh_nd_oe")
	assert.NotNil(t, err)
	err = validate("joh.nd.oe")
	assert.NotNil(t, err)
}

func TestEmail(t *testing.T) {
	/*
		- `email` can only contain alphanumeric characters and these special characters: _.-@+
		- its length must be greater than 3 and not more than 255
		- special character must be followed by at least one alphanumeric character
		- prefix and suffix cannot be a special character
		- must contain char @ and must be appeared once in the string
	*/
	var validate = func(email string) error {
		return strgo.New(email).
			MinLength(4).
			MaxLength(255).
			OnlyContainChars(strgo.ALPHANUMERIC).
			OnlyContainChars([]string{"_", ".", "@", "-", "+"}).
			MustBeFollowedByChars([]string{"_", ".", "@", "-", "+"}, strgo.ALPHANUMERIC).
			MustNotContainPrefixChars([]string{"_", ".", "@", "-", "+"}).
			MustNotContainSuffixChars([]string{"_", ".", "@", "-", "+"}).
			MustContainCharsOnce([]string{"@"}).
			Validate()
	}
	err := validate("johndoe@email.com")
	assert.Nil(t, err)
	err = validate("john_doe@email.com")
	assert.Nil(t, err)
	err = validate("john_do.e@email.com")
	assert.Nil(t, err)
	err = validate("john-doe@email.com")
	assert.Nil(t, err)
	err = validate("johndoe@email")
	assert.Nil(t, err)
	err = validate("johndoe123@email")
	assert.Nil(t, err)
	err = validate("johndoe123@email")
	assert.Nil(t, err)
	err = validate("john+doe123@email")
	assert.Nil(t, err)
	err = validate("johndoe123email")
	assert.NotNil(t, err)
	err = validate("johndoe123.email")
	assert.NotNil(t, err)
	err = validate("john@doe123@email")
	assert.NotNil(t, err)
	err = validate(".johndoe123@email")
	assert.NotNil(t, err)
	err = validate("johndoe123@email.")
	assert.NotNil(t, err)
	err = validate("johndoe123@")
	assert.NotNil(t, err)
	err = validate("john_.doe123@email")
	assert.NotNil(t, err)
	err = validate("johndoe123.@email")
	assert.NotNil(t, err)
}
