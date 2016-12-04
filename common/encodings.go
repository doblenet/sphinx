package common

import (
	"encoding/base64"
	"encoding/hex"
)


type PassCodec interface {
	Decode(in string) ([]byte, error)
	Encode(in []byte) (string,error)
}

type hexCodec struct {}
type b64Codec struct {}

var (
	str2enc = map[string](func() PassCodec) {
		"b64": b64Codec_New,
		"hex": hexCodec_New,
		"base64": b64Codec_New,	// just in case...
	}
)

func CodecConv(name string) PassCodec {
	if c,ok:= str2enc[name]; ok {
		return c()
	}
	return nil
}


func b64Codec_New() PassCodec { return new(b64Codec) }
func hexCodec_New() PassCodec { return new(hexCodec) }

// **** HEX CODEC ****
func (c *hexCodec) Encode(in []byte) (string, error) {
	return hex.EncodeToString(in), nil
}
func (c *hexCodec) Decode(in string) ([]byte, error) {
	return hex.DecodeString(in)
}


// **** BASE64 CODEC ****
func (c *b64Codec) Encode(in []byte) (string,error) {
	return base64.URLEncoding.EncodeToString(in), nil
}
func (c *b64Codec) Decode(in string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(in)
}
