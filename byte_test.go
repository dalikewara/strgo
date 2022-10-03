package strgo_test

import (
	"github.com/dalikewara/strgo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByte_MinLength(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MinLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MinLength: 8,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be less than 8")
}

func TestByte_MaxLength(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MaxLength: 7,
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MaxLength: 6,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string length cannot be more than 6")
}

func TestByte_OnlyContains(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContains: []byte{'j', 'o', 'h', 'n', 'd', 'e'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContains: []byte{'j', 'o', 'h', 'n', 'd'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain char: e")
}

func TestByte_OnlyContainsPrefix(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContainsPrefix: []byte{'o', 'j'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContainsPrefix: []byte{'o', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain prefix char: j")
}

func TestByte_OnlyContainsSuffix(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContainsSuffix: []byte{'o', 'e'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		OnlyContainsSuffix: []byte{'o', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string cannot contain suffix char: e")
}

func TestByte_MustContains(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustContains: []byte{'n', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustContains: []byte{'n', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain char: k")
}

func TestByte_MustContainsOnce(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustContainsOnce: []byte{'n', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustContainsOnce: []byte{'n', 'd', 'h', 'k'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must contain char: k")
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustContainsOnce: []byte{'n', 'o'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be appeared once in the string")
}

func TestByte_MustNotContains(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContains: []byte{'k', 'r'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContains: []byte{'k', 'e'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain char: e")
}

func TestByte_MustNotContainsPrefix(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContainsPrefix: []byte{'o', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContainsPrefix: []byte{'o', 'j'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain prefix: j")
}

func TestByte_MustNotContainsSuffix(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContainsSuffix: []byte{'o', 'h'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustNotContainsSuffix: []byte{'o', 'e'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must not contain suffix: e")
}

func TestByte_MustBeFollowedBy(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MustBeFollowedBy: [2][]byte{{'h', 'd'}, {'m', 'n', 'o'}},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustBeFollowedBy: [2][]byte{{'o', 'd'}, {'e', 'o', 'd', 'j', 'h'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: d, must be followed with at least one of these characters: eodjh")
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustBeFollowedBy: [2][]byte{{'h', 'o'}, {'d', 'k', 'l'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be followed with at least one of these characters: dkl")
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MustBeFollowedBy: [2][]byte{{'h', 'o'}, {'d', 'k', 'j'}},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be followed with at least one of these characters: dkj")
}

func TestByte_AtLeastHaveUpperLetterCount(t *testing.T) {
	err := strgo.Byte("JoHndoe", &strgo.ByteCondition{
		AtLeastHaveUpperLetterCount: 1,
	})
	assert.Nil(t, err)
	err = strgo.Byte("joHndoe", &strgo.ByteCondition{
		AtLeastHaveUpperLetterCount: 2,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must have at least 2 upper case letter(s)")
}

func TestByte_AtLeastHaveLowerLetterCount(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		AtLeastHaveLowerLetterCount: 1,
	})
	assert.Nil(t, err)
	err = strgo.Byte("JOHNDoE", &strgo.ByteCondition{
		AtLeastHaveLowerLetterCount: 2,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must have at least 2 lower case letter(s)")
}

func TestByte_AtLeastHaveNumberCount(t *testing.T) {
	err := strgo.Byte("johndoe123", &strgo.ByteCondition{
		AtLeastHaveNumberCount: 1,
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe456", &strgo.ByteCondition{
		AtLeastHaveNumberCount: 4,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must have at least 4 number(s)")
}

func TestByte_AtLeastHaveSpecialCharCount(t *testing.T) {
	err := strgo.Byte("john_doe", &strgo.ByteCondition{
		AtLeastHaveSpecialCharCount: 1,
	})
	assert.Nil(t, err)
	err = strgo.Byte("john_doe", &strgo.ByteCondition{
		AtLeastHaveSpecialCharCount: 2,
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the string must have at least 2 special char(s)")
}

func TestByte_MayContainsOnce(t *testing.T) {
	err := strgo.Byte("johndoe", &strgo.ByteCondition{
		MayContainsOnce: []byte{'k', 'r', 'h', 'n', 'j', 'd'},
	})
	assert.Nil(t, err)
	err = strgo.Byte("johndoe", &strgo.ByteCondition{
		MayContainsOnce: []byte{'k', 'o'},
	})
	assert.NotNil(t, err)
	assert.EqualError(t, err, "the char: o, must be appeared once in the string")
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
		return strgo.Byte(username, &strgo.ByteCondition{
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
		return strgo.Byte(email, &strgo.ByteCondition{
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

func TestPassword(t *testing.T) {
	/*
		- `password` can only contain alphanumeric characters and special characters: !"#$% &'()*+,-./:;<=>?@[\]^_`{|}~
		- its length must be greater than 5 and not more than 32
		- at least contain one lower and upper case letter, one number and one special character
	*/
	var validate = func(password string) error {
		return strgo.Byte(password, &strgo.ByteCondition{
			MinLength:                   6,
			MaxLength:                   32,
			OnlyContains:                strgo.CharsByte,
			AtLeastHaveUpperLetterCount: 1,
			AtLeastHaveLowerLetterCount: 1,
			AtLeastHaveNumberCount:      1,
			AtLeastHaveSpecialCharCount: 1,
		})
	}
	err := validate("J()hndoe123")
	assert.Nil(t, err)
	err = validate("John_doe123")
	assert.Nil(t, err)
	err = validate("johndoe")
	assert.NotNil(t, err)
	err = validate("johndoe123")
	assert.NotNil(t, err)
	err = validate("Johndoe123")
	assert.NotNil(t, err)
}
