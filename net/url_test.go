package net

import (
	"fmt"
)

func urlTest (s string){
	s, err:= RequestURI(s)

	if err==nil{
		fmt.Printf("Test string : %v",s)
	}
}
