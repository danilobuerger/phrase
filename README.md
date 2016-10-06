[![Build Status](https://travis-ci.org/danilobuerger/phrase.svg?branch=master)](https://travis-ci.org/danilobuerger/phrase) [![Coverage Status](https://coveralls.io/repos/github/danilobuerger/phrase/badge.svg?branch=master)](https://coveralls.io/github/danilobuerger/phrase?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/danilobuerger/phrase)](https://goreportcard.com/report/github.com/danilobuerger/phrase)

# phrase

Converts a phrase to a different case (e.g. CamelCase).

## CamelCase

`phrase.ToCamel` converts a lower case string with separators to a camel case string.

```go
phrase.ToCamel("abc def", " ")
// Returns: "AbcDef" (without quotes)
```
