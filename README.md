[![Build Status](https://travis-ci.org/qumonintelligence/go-common.svg?branch=master)](https://travis-ci.org/qumonintelligence/go-common)

# go-common
Common library for Golang

# Common strings utils


## lang.StringOf

Return "" if a string pointer is nil
```
var value *string

lang.StringOf(value) // return ""
```

```
var value *string
x := "abc"
value = &x

lang.StringOf(value) // return "abc"
```

```
var value []byte
lang.StringOf(value) // return string(value)
```