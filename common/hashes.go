package common

import (
//	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
)

var (
	str2hash = map[string](func() hash.Hash){
		"md5":    md5.New,
		"sha1":   sha1.New,
		"sha224": sha256.New224,
		"sha256": sha256.New,
		"sha384": sha512.New384,
		"sha512": sha512.New,
	}
)

func HashConv(name string) hash.Hash {

	if hf, ok := str2hash[name]; ok {
		return hf()
	}
	return nil
}
