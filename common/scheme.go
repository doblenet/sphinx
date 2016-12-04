package common;

import (
	"errors"
	"strings"
//	"hash"
//	"bytes"
)

const (
	schemePrefix = "SCRAM"
)


func ParseScheme(in string) (string,string,string,error) {

 var i,j int;
	l := len(in)-1	// adjusted to 'last valid index' for efficiency reasons
	if in[0]!='{' || in[l]!='}' {
		return "","","",errors.New("Malformed syntax")
	}

	if i=strings.Index(in,"-"); i<1 {
		return "","","",errors.New("Missing hash specifier")
	}
	scheme := in[1:i]

	enc := "plain"
	if j=strings.LastIndex(in,"."); j>1 {
		enc = in[j+1:l]
	} else {
		j = l
	}

	hf := strings.ToLower(in[i+1:j])

	return scheme,hf,enc,nil
}
