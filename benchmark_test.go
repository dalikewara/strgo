package strgo_test

import (
	"github.com/dalikewara/strgo"
	"log"
	"regexp"
	"strings"
	"testing"
	"time"
)

func BenchmarkByte(b *testing.B) {
	bt := []byte("akdjfnafjweifwef..,./'91840jsafnkafkabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321awjdbjwfjhabfjwqbfjebfawkhfiuqwuqwqmlksmANXMASNBFIQWHFDIQWDQWIJODFWQHFIWQHEU12Y431U4IU4O12KJEN2JEHIO2UEJSBasbfkjaenfkqnefkehmdqwdiwqbrwqbrjwqkdfwqfjwqnfqehriquhrqwnrwoqrwoqdqwohiwoqjewoqihewqu")
	strgo.Byte("Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliqua.Utenimadminimveniam.quisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequat.Duisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariatur.Excepteursintoccaecatcupidatatnonproident.suntinculpaquiofficiadeseruntmollitanimidestlaborum", &strgo.ByteCondition{
		OnlyContains:                append(bt, strgo.SpecialCharsByte...),
		OnlyContainsPrefix:          bt,
		OnlyContainsSuffix:          bt,
		MustContains:                strgo.AlphanumericByte,
		MustContainsOnce:            []byte{'+'},
		MustNotContains:             strgo.BracketsByte,
		MustNotContainsPrefix:       strgo.SpecialCharsByte,
		MustNotContainsSuffix:       strgo.SpecialCharsByte,
		MayContainsOnce:             []byte{'+'},
		MustBeFollowedBy:            [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
		AtLeastHaveUpperLetterCount: 2,
		AtLeastHaveLowerLetterCount: 2,
		AtLeastHaveNumberCount:      2,
		AtLeastHaveSpecialCharCount: 2,
	})
}

func BenchmarkByteUsername(b *testing.B) {
	strgo.Byte("john_doe.123", &strgo.ByteCondition{
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
		MayContainsOnce:  []byte{'_', '.'},
	})
}

func BenchmarkByteUsernameLongText(b *testing.B) {
	strgo.Byte("Loremipsumdolorsitametconse_ct.eturadipiscingelitabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquaUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatDuisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturExcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum", &strgo.ByteCondition{
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
		MayContainsOnce:  []byte{'_', '.'},
	})
}

func BenchmarkByteEmail(b *testing.B) {
	strgo.Byte("john+doe123@email", &strgo.ByteCondition{
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	})
}

func BenchmarkByteEmailLongText(b *testing.B) {
	strgo.Byte("Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquaUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatDuisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturExcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum", &strgo.ByteCondition{
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	})
}

func BenchmarkBytePassword(b *testing.B) {
	strgo.Byte("john_DOe.123", &strgo.ByteCondition{
		OnlyContains:                strgo.CharsByte,
		AtLeastHaveUpperLetterCount: 2,
		AtLeastHaveLowerLetterCount: 2,
		AtLeastHaveNumberCount:      2,
		AtLeastHaveSpecialCharCount: 2,
	})
}

func BenchmarkBytePasswordLongText(b *testing.B) {
	strgo.Byte("Loremipsumdolorsitametconse_ct.eturadipiscingelitabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquaUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatDuisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturExcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum", &strgo.ByteCondition{
		OnlyContains:                strgo.CharsByte,
		AtLeastHaveUpperLetterCount: 2,
		AtLeastHaveLowerLetterCount: 2,
		AtLeastHaveNumberCount:      2,
		AtLeastHaveSpecialCharCount: 2,
	})
}

func BenchmarkString(b *testing.B) {
	w := strings.Split("Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliqua.Utenimadminimveniam.quisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequat.Duisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariatur.Excepteursintoccaecatcupidatatnonproident.suntinculpaquiofficiadeseruntmollitanimidestlaborum", "@")
	strgo.String("Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliqua.Utenimadminimveniam.quisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequat.Duisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariatur.Excepteursintoccaecatcupidatatnonproident.suntinculpaquiofficiadeseruntmollitanimidestlaborum", &strgo.StringCondition{
		OnlyContainsPrefixWord:    w,
		OnlyContainsSuffixWord:    w,
		MustContainsWord:          w,
		MustContainsWordOnce:      []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsWord:       []string{"+olo2", "met.con2", "lit.abcde2", "4321seddoei2"},
		MustNotContainsPrefixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsSuffixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MayContainsWordOnce:       []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
	})
}

func BenchmarkRegexUsername(b *testing.B) {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-@]+$`)
	_ = regex.MatchString("john_doe.123")
}

func BenchmarkRegexUsernameLongtext(b *testing.B) {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-@]+$`)
	_ = regex.MatchString("Loremipsumdolorsitametconse_ct.eturadipiscingelitabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquaUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatDuisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturExcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum")
}

func BenchmarkRegexEmail(b *testing.B) {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]+$`)
	_ = regex.MatchString("john+doe123@email")
}

func BenchmarkRegexEmailLongText(b *testing.B) {
	regex, _ := regexp.Compile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]+$`)
	_ = regex.MatchString("Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliquaUtenimadminimveniamquisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequatDuisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariaturExcepteursintoccaecatcupidatatnonproidentsuntinculpaquiofficiadeseruntmollitanimidestlaborum")
}

func TestElapsedTime(t *testing.T) {
	text := "Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliqua.Utenimadminimveniam.quisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequat.Duisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariatur.Excepteursintoccaecatcupidatatnonproident.suntinculpaquiofficiadeseruntmollitanimidestlaborum"
	text2 := "akdjfnafjweifwef..,./'91840jsafnkafkabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321awjdbjwfjhabfjwqbfjebfawkhfiuqwuqwqmlksmANXMASNBFIQWHFDIQWDQWIJODFWQHFIWQHEU12Y431U4IU4O12KJEN2JEHIO2UEJSBasbfkjaenfkqnefkehmdqwdiwqbrwqbrjwqkdfwqfjwqnfqehriquhrqwnrwoqrwoqdqwohiwoqjewoqihewqu"
	text2Byte := []byte(text2)
	words := strings.Split(text, "a")
	r1 := `^[a-z0-9._%+\-@]+$`
	r2 := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]+$`

	log.Printf("					text len 	= %v", len(text))
	log.Println("					elapsed time	= ns/nanoseconds < µs/microseconds")

	// bytes

	start := time.Now()
	err := strgo.Byte(text, &strgo.ByteCondition{
		OnlyContains:                append(text2Byte, strgo.SpecialCharsByte...),
		OnlyContainsPrefix:          text2Byte,
		OnlyContainsSuffix:          text2Byte,
		MustContains:                strgo.AlphanumericByte,
		MustContainsOnce:            []byte{'+'},
		MustNotContains:             strgo.BracketsByte,
		MustNotContainsPrefix:       strgo.SpecialCharsByte,
		MustNotContainsSuffix:       strgo.SpecialCharsByte,
		MayContainsOnce:             []byte{'+'},
		MustBeFollowedBy:            [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
		AtLeastHaveUpperLetterCount: 2,
		AtLeastHaveLowerLetterCount: 2,
		AtLeastHaveNumberCount:      2,
		AtLeastHaveSpecialCharCount: 2,
	})
	elapsed := time.Since(start)
	log.Printf("strgo bytes 					%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Byte("john_doe.123", &strgo.ByteCondition{
		MinLength:        3,
		MaxLength:        20,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
		MayContainsOnce:  []byte{'_', '.'},
	})
	elapsed = time.Since(start)
	log.Printf("strgo bytes username (john_doe.123) 		%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Byte("john+doe123@email", &strgo.ByteCondition{
		MinLength:        4,
		MaxLength:        255,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	})
	elapsed = time.Since(start)
	log.Printf("strgo bytes email (john+doe123@email) 		%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Byte("John_Doe.123", &strgo.ByteCondition{
		OnlyContains:                strgo.CharsByte,
		AtLeastHaveUpperLetterCount: 2,
		AtLeastHaveLowerLetterCount: 2,
		AtLeastHaveNumberCount:      2,
		AtLeastHaveSpecialCharCount: 2,
	})
	elapsed = time.Since(start)
	log.Printf("strgo bytes password (John_Doe.123) 		%s %v", elapsed, err)

	// string

	start = time.Now()
	err = strgo.String(text, &strgo.StringCondition{
		OnlyContainsPrefixWord:    words,
		OnlyContainsSuffixWord:    words,
		MustContainsWord:          words,
		MustContainsWordOnce:      []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsWord:       []string{"+olo2", "met.con2", "lit.abcde2", "4321seddoei2"},
		MustNotContainsPrefixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsSuffixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MayContainsWordOnce:       []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
	})
	elapsed = time.Since(start)
	log.Printf("strgo string 					%s %v", elapsed, err)

	// regex

	start = time.Now()
	regex, _ := regexp.Compile(r1)
	_ = regex.MatchString(text)
	elapsed = time.Since(start)
	log.Printf("regex %s				%s", r1, elapsed)

	start = time.Now()
	regex, _ = regexp.Compile(r2)
	_ = regex.MatchString(text)
	elapsed = time.Since(start)
	log.Printf("regex %s 	%s", r2, elapsed)

	start = time.Now()
	regex, _ = regexp.Compile(r1)
	_ = regex.MatchString("john_doe.123")
	elapsed = time.Since(start)
	log.Printf("regex username (john_doe.123)			%s", elapsed)

	start = time.Now()
	regex, _ = regexp.Compile(r2)
	_ = regex.MatchString("john+doe123@email")
	elapsed = time.Since(start)
	log.Printf("regex email (john+doe123@email)			%s", elapsed)
}
