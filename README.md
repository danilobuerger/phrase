# phrase

Converts a phrase to a different case (e.g. CamelCase).

## CamelCase

`phrase.ToCamel` converts a lower case string with separators to a camel case string.

```go
phrase.ToCamel("abc def", " ")
// Returns: "AbcDef" (without quotes)
```
