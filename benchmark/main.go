package main

import (
	"github.com/dalikewara/strgo"
	"log"
	"regexp"
	"strings"
	"time"
)

func main() {
	text := "Loremipsumd+olorsitamet.consectetur@adipiscingelit.abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321seddoeiusmodtemporincididuntutlaboreetdoloremagnaaliqua.Utenimadminimveniam.quisnostrudexercitationullamcolaborisnisiutaliquipexeacommodoconsequat.Duisauteiruredolorinreprehenderitinvoluptatevelitessecillumdoloreeufugiatnullapariatur.Excepteursintoccaecatcupidatatnonproident.suntinculpaquiofficiadeseruntmollitanimidestlaborum"
	text2 := "akdjfnafjweifwef..,./'91840jsafnkafkabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321awjdbjwfjhabfjwqbfjebfawkhfiuqwuqwqmlksmANXMASNBFIQWHFDIQWDQWIJODFWQHFIWQHEU12Y431U4IU4O12KJEN2JEHIO2UEJSBasbfkjaenfkqnefkehmdqwdiwqbrwqbrjwqkdfwqfjwqnfqehriquhrqwnrwoqrwoqdqwohiwoqjewoqihewqu"
	text2Byte := []byte(text2)
	words := strings.Split(text, "a")
	r1 := `^[a-z0-9._%+\-@]+$`
	r2 := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]+$`

	log.Printf("					text len 	= %v", len(text))
	log.Println("					elapsed time	= ns/nanoseconds < µs/microseconds")

	start := time.Now()
	err := strgo.Validate(text, &strgo.Condition{
		OnlyContains:              append(text2Byte, strgo.SpecialCharsByte...),
		OnlyContainsPrefix:        text2Byte,
		OnlyContainsSuffix:        text2Byte,
		OnlyContainsPrefixWord:    words,
		OnlyContainsSuffixWord:    words,
		MustContains:              strgo.AlphanumericByte,
		MustContainsWord:          words,
		MustContainsOnce:          []byte{'+'},
		MustContainsWordOnce:      []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContains:           strgo.BracketsByte,
		MustNotContainsWord:       []string{"+olo2", "met.con2", "lit.abcde2", "4321seddoei2"},
		MustNotContainsPrefix:     strgo.SpecialCharsByte,
		MustNotContainsSuffix:     strgo.SpecialCharsByte,
		MustNotContainsPrefixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsSuffixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MayContainsOnce:           []byte{'+'},
		MayContainsWordOnce:       []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustBeFollowedBy:          [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
	})
	elapsed := time.Since(start)
	log.Printf("all conditions set 					%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Validate(text, &strgo.Condition{
		OnlyContains:          append(text2Byte, strgo.SpecialCharsByte...),
		OnlyContainsPrefix:    text2Byte,
		OnlyContainsSuffix:    text2Byte,
		MustContains:          strgo.AlphanumericByte,
		MustContainsOnce:      []byte{'+'},
		MustNotContains:       strgo.BracketsByte,
		MustNotContainsPrefix: strgo.SpecialCharsByte,
		MustNotContainsSuffix: strgo.SpecialCharsByte,
		MayContainsOnce:       []byte{'+'},
		MustBeFollowedBy:      [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
	})
	elapsed = time.Since(start)
	log.Printf("all bytes/chars conditions set 			%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Validate(text, &strgo.Condition{
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
	log.Printf("all words conditions set 				%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Validate("john_doe.123", &strgo.Condition{
		MinLength:        3,
		MaxLength:        20,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
		MayContainsOnce:  []byte{'_', '.'},
	})
	elapsed = time.Since(start)
	log.Printf("username (john_doe.123) conditions set 		%s %v", elapsed, err)

	start = time.Now()
	err = strgo.Validate("john+doe123@email", &strgo.Condition{
		MinLength:        4,
		MaxLength:        255,
		OnlyContains:     append(strgo.AlphanumericByte, []byte{'_', '.', '@', '-', '+'}...),
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	})
	elapsed = time.Since(start)
	log.Printf("email (john+doe123@email) conditions set 		%s %v", elapsed, err)

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
	log.Printf("username (john_doe.123) regex			%s", elapsed)

	start = time.Now()
	regex, _ = regexp.Compile(r2)
	_ = regex.MatchString("john+doe123@email")
	elapsed = time.Since(start)
	log.Printf("email (john+doe123@email) regex			%s", elapsed)
}
