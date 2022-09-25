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
	specialCharByte := []byte{
		' ', '!', '"', '#', '$', '%', '&', '\'', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@',
		'[', '\\', ']', '^', '_', '`', '{', '|', '}', '~',
	}
	alphanumericByte := []byte{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
	bracketByte := []byte{
		'(', ')', '[', ']', '{', '}',
	}
	words := strings.Split(text, "a")
	r1 := `^[a-z0-9._%+\-@]+$`
	r2 := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]+$`

	start := time.Now()
	strgo.Validate(text, &strgo.Condition{
		OnlyContains:              text2Byte,
		OnlyContainsPrefix:        text2Byte,
		OnlyContainsSuffix:        text2Byte,
		OnlyContainsPrefixWord:    words,
		OnlyContainsSuffixWord:    words,
		MustContains:              alphanumericByte,
		MustContainsWord:          words,
		MustContainsOnce:          []byte{'+'},
		MustContainsWordOnce:      []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContains:           bracketByte,
		MustNotContainsWord:       []string{"+olo2", "met.con2", "lit.abcde2", "4321seddoei2"},
		MustNotContainsPrefix:     specialCharByte,
		MustNotContainsSuffix:     specialCharByte,
		MustNotContainsPrefixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustNotContainsSuffixWord: []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MayContainsOnce:           []byte{'+'},
		MayContainsWordOnce:       []string{"+olo", "met.con", "lit.abcde", "4321seddoei"},
		MustBeFollowedBy:          [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
	}, &strgo.Condition{
		OnlyContains: specialCharByte,
	})
	elapsed := time.Since(start)
	log.Printf("all conditions set 					%s", elapsed)

	start = time.Now()
	strgo.Validate(text, &strgo.Condition{
		OnlyContains:          text2Byte,
		OnlyContainsPrefix:    text2Byte,
		OnlyContainsSuffix:    text2Byte,
		MustContains:          alphanumericByte,
		MustContainsOnce:      []byte{'+'},
		MustNotContains:       bracketByte,
		MustNotContainsPrefix: specialCharByte,
		MustNotContainsSuffix: specialCharByte,
		MayContainsOnce:       []byte{'+'},
		MustBeFollowedBy:      [2][]byte{strgo.SpecialCharsByte, strgo.CharsByte},
	}, &strgo.Condition{
		OnlyContains: specialCharByte,
	})
	elapsed = time.Since(start)
	log.Printf("all bytes/chars conditions set 			%s", elapsed)

	start = time.Now()
	strgo.Validate(text, &strgo.Condition{
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
	log.Printf("all words conditions set 				%s", elapsed)

	start = time.Now()
	strgo.Validate("john_doe.123", &strgo.Condition{
		MinLength:        3,
		MaxLength:        20,
		OnlyContains:     strgo.AlphanumericByte,
		MustBeFollowedBy: [2][]byte{{'_', '.'}, strgo.AlphanumericByte},
		MayContainsOnce:  []byte{'_', '.'},
	}, &strgo.Condition{
		OnlyContains: []byte{'_', '.'},
	})
	elapsed = time.Since(start)
	log.Printf("username conditions set 				%s", elapsed)

	start = time.Now()
	strgo.Validate("john+doe123@email", &strgo.Condition{
		MinLength:        4,
		MaxLength:        255,
		OnlyContains:     strgo.AlphanumericByte,
		MustBeFollowedBy: [2][]byte{{'_', '.', '@', '-', '+'}, strgo.AlphanumericByte},
		MustContainsOnce: []byte{'@'},
	}, &strgo.Condition{
		OnlyContains: []byte{'_', '.', '@', '-', '+'},
	})
	elapsed = time.Since(start)
	log.Printf("email conditions set 				%s", elapsed)

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
	log.Printf("username regex					%s", elapsed)

	start = time.Now()
	regex, _ = regexp.Compile(r2)
	_ = regex.MatchString("john+doe123@email")
	elapsed = time.Since(start)
	log.Printf("email regex						%s", elapsed)
}
