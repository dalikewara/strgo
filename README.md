# strgo

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/dalikewara/strgo)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/dalikewara/strgo)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/dalikewara/strgo)
![GitHub license](https://img.shields.io/github/license/dalikewara/strgo)

**strgo** is a helper to validate a string based on a format specification that you set before. You can use this package to validate
some common string case like `username`, `name`, `password`, `email`, etc.

## Getting started

### Installation

You can use the `go get` method:

```bash
go get github.com/dalikewara/strgo
```

### Usage

This is an example to validate a `username` string:

- `username` may only contain alphanumeric characters, underscores and periods
- its length must be greater than 2 and not more than 20
- special character must be followed by at least one alphanumeric character
- prefix and suffix cannot be a special character
- allowed special characters must be appeared once in the string

```go
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

validate("johndoe") // valid
validate("john_doe") // valid
validate("john.doe") // valid
validate("johndoe123") // valid
validate("john.doe123") // valid
validate("john_doe123") // valid
validate("john_doe.123") // valid
validate("_johndoe") // not valid
validate("__johndoe") // not valid
validate(".johndoe") // not valid
validate("..johndoe") // not valid
validate("johndoe_") // not valid
validate("johndoe__") // not valid
validate("johndoe.") // not valid
validate("johndoe..") // not valid
validate("john__doe") // not valid
validate("john_.doe") // not valid
validate("john..doe") // not valid
validate("john._doe") // not valid
validate("john@doe") // not valid
validate("@johndoe") // not valid
validate("johndoe@") // not valid
validate("joh_nd_oe") // not valid
validate("joh.nd.oe") // not valid
```

## Release

### Changelog

Read at [CHANGELOG.md](https://github.com/dalikewara/strgo/blob/master/CHANGELOG.md)

### Credits

Copyright &copy; 2022 [Dali Kewara](https://www.dalikewara.com)

### License

[MIT License](https://github.com/dalikewara/strgo/blob/master/LICENSE)
