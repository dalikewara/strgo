package strgo_test

import (
	"github.com/dalikewara/strgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrGo_MinLength(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MinLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MinLength: 8,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be less than 8")
}

func TestStrGo_MaxLength(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MaxLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MaxLength: 6,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be more than 6")
}

func TestStrGo_OnlyContains(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		OnlyContains: []byte{'j', 'o', 'h', 'n', 'd', 'e'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		OnlyContains: []byte{'j', 'o', 'h', 'n', 'd'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain char: e")
}

func TestStrGo_OnlyContainsPrefix(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsPrefix: []byte{'o', 'j'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsPrefix: []byte{'o', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain prefix char: j")
}

func TestStrGo_OnlyContainsSuffix(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsSuffix: []byte{'o', 'e'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsSuffix: []byte{'o', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain suffix char: e")
}

func TestStrGo_OnlyContainsPrefixWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsPrefixWord: []string{"hn", "joh"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsPrefixWord: []string{"noh", "koh", "noh2", "koh2"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string prefix doesn't match with the given prefix words")
}

func TestStrGo_OnlyContainsSuffixWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsSuffixWord: []string{"eo", "oe"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		OnlyContainsSuffixWord: []string{"ne", "ho"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string suffix doesn't match with the given suffix words")
}

func TestStrGo_MustContains(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustContains: []byte{'n', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustContains: []byte{'n', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain char: k")
}

func TestStrGo_MustContainsWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustContainsWord: []string{"john", "hnd"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustContainsWord: []string{"ohn", "de"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: de")
}

func TestStrGo_MustContainsOnce(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustContainsOnce: []byte{'n', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustContainsOnce: []byte{'n', 'd', 'h', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain char: k")
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustContainsOnce: []byte{'n', 'o'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be appeared once in the string")
}

func TestStrGo_MustContainsWordOnce(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustContainsWordOnce: []string{"john", "hnd"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustContainsWordOnce: []string{"ohn", "de"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: de, and it must be appeared once in the string")
	err = strgo.Validate("johndoedoe", &strgo.Condition{
		MustContainsWordOnce: []string{"ohn", "doe"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: doe, and it must be appeared once in the string")
}

func TestStrGo_MustNotContains(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContains: []byte{'k', 'r'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContains: []byte{'k', 'e'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain char: e")
}

func TestStrGo_MustNotContainsWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsWord: []string{"dor", "johk"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsWord: []string{"hng", "ohn"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain word: ohn")
}

func TestStrGo_MustNotContainsPrefix(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsPrefix: []byte{'o', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsPrefix: []byte{'o', 'j'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain prefix: j")
}

func TestStrGo_MustNotContainsSuffix(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsSuffix: []byte{'o', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsSuffix: []byte{'o', 'e'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain suffix: e")
}

func TestStrGo_MustNotContainsPrefixWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsPrefixWord: []string{"hn", "koh"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsPrefixWord: []string{"noh", "joh"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain prefix word: joh")
}

func TestStrGo_MustNotContainsSuffixWord(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsSuffixWord: []string{"eo", "ode"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustNotContainsSuffixWord: []string{"doe", "joh"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain suffix word: doe")
}

func TestStrGo_MayContainsOnce(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MayContainsOnce: []byte{'k', 'r', 'h', 'n', 'j', 'd'},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MayContainsOnce: []byte{'k', 'o'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be appeared once in the string")
}

func TestStrGo_MayContainsWordOnce(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MayContainsWordOnce: []string{"khn", "rhn"},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoedoe", &strgo.Condition{
		MayContainsWordOnce: []string{"khn", "doe"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the word: doe, must be appeared once in the string")
}

func TestStrGo_MustBeFollowedBy(t *testing.T) {
	err := strgo.Validate("johndoe", &strgo.Condition{
		MustBeFollowedBy: [2][]byte{{'h', 'd'}, {'m', 'n', 'o'}},
	})
	assert.Nil(t, err)
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustBeFollowedBy: [2][]byte{{'o', 'd'}, {'e', 'o', 'd', 'j', 'h'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: d, must be followed with at least one of these characters: eodjh")
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustBeFollowedBy: [2][]byte{{'h', 'o'}, {'d', 'k', 'l'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be followed with at least one of these characters: dkl")
	err = strgo.Validate("johndoe", &strgo.Condition{
		MustBeFollowedBy: [2][]byte{{'h', 'o'}, {'d', 'k', 'j'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be followed with at least one of these characters: dkj")
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
		return strgo.Validate(username, &strgo.Condition{
			MinLength:        3,
			MaxLength:        20,
			OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
			MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
			MayContainsOnce:  []byte{'_', '.'},
		})
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
		return strgo.Validate(email, &strgo.Condition{
			MinLength:        4,
			MaxLength:        255,
			OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
			MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
			MustContainsOnce: []byte{'@'},
		})
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
