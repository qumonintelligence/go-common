[![Build Status](https://travis-ci.org/qumonintelligence/go-common.svg?branch=master)](https://travis-ci.org/qumonintelligence/go-common)

# go-common
Common library for Golang

# Common strings utils

```
var value *string

lang.StringOf(value) // return ""

x := "abc"
value = &x

lang.StringOf(value) // return "abc"
```