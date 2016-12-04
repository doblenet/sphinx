package common;

import (
	"testing"
	"fmt"
)


func TestEncodingHex(t *testing.T) {

 var s = "123456789abcdef"

	fmt.Println("INPUT=",s)

	c1 := CodecConv("hex")
	
	r1,_:=c1.Encode([]byte(s))
	fmt.Println("encoded=",r1)
	s1,_:=c1.Decode(r1)
	fmt.Println("decoded=",string(s1))
}

func TestEncodingBase64(t *testing.T) {

 var s = "123456789abcdef"

	c2 := CodecConv("b64")

	r2,_:=c2.Encode([]byte(s))
	fmt.Println("encoded=",r2)
	s2,_:=c2.Decode(r2)
        fmt.Println("decoded=",string(s2))
}
