package common;

import (
	"testing"
	"fmt"
)


func TestScheme(t *testing.T) {

 var s,h,e string
 var r error


	_,_,_,r = ParseScheme("invalid_scheme")
	if nil==r {
		fmt.Println("Validated 'invalid_scheme'?!!")
		t.Fail()
	}

	s,h,e,r = ParseScheme("{SCRAM-SHA1.base64}")
	if nil!=r {
		t.Fail()
	}
	fmt.Println(s,h,e)

	s,h,e,r = ParseScheme("{SCRAM-SHA256}")
	if nil!=r {
		t.Fail()
	}
	fmt.Println(s,h,e)

	s,h,e,r = ParseScheme("{CRAM-MD5.hex}")
	if nil!=r {
		t.Fail()
	}
	fmt.Println(s,h,e)

	
}
