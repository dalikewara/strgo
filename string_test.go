package strgo_test

import (
	"github.com/dalikewara/strgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString_MinLength(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MinLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MinLength: 8,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be less than 8")
}

func TestString_MaxLength(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MaxLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MaxLength: 6,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be more than 6")
}

func TestString_OnlyContainsPrefixWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		OnlyContainsPrefixWord: []string{"hn", "joh"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		OnlyContainsPrefixWord: []string{"noh", "koh", "noh2", "koh2"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string prefix doesn't match with the given prefix words")
}

func TestString_OnlyContainsSuffixWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		OnlyContainsSuffixWord: []string{"eo", "oe"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		OnlyContainsSuffixWord: []string{"ne", "ho"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string suffix doesn't match with the given suffix words")
}

func TestString_MustContainsWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MustContainsWord: []string{"john", "hnd"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MustContainsWord: []string{"ohn", "de"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: de")
}

func TestString_MustContainsWordOnce(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MustContainsWordOnce: []string{"john", "hnd"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MustContainsWordOnce: []string{"ohn", "de"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: de, and it must be appeared once in the string")
	err = strgo.String("johndoedoe", &strgo.StringCondition{
		MustContainsWordOnce: []string{"ohn", "doe"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain word: doe, and it must be appeared once in the string")
}

func TestString_MustNotContainsWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsWord: []string{"dor", "johk"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsWord: []string{"hng", "ohn"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain word: ohn")
}

func TestString_MustNotContainsPrefixWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsPrefixWord: []string{"hn", "koh"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsPrefixWord: []string{"noh", "joh"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain prefix word: joh")
}

func TestString_MustNotContainsSuffixWord(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsSuffixWord: []string{"eo", "ode"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoe", &strgo.StringCondition{
		MustNotContainsSuffixWord: []string{"doe", "joh"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain suffix word: doe")
}

func TestString_MayContainsWordOnce(t *testing.T) {
	err := strgo.String("johndoe", &strgo.StringCondition{
		MayContainsWordOnce: []string{"khn", "rhn"},
	})
	assert.Nil(t, err)
	err = strgo.String("johndoedoe", &strgo.StringCondition{
		MayContainsWordOnce: []string{"khn", "doe"},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the word: doe, must be appeared once in the string")
}
