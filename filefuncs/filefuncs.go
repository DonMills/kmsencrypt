package filefuncs

import (
	"bytes"
	"encoding/base64"
)

var sep = []byte{0, 1, 0, 1, 0, 1}

func CreateEncFile(ciphertext []byte, iv []byte, cipherdatakey []byte) []byte {
	bufferslice := [][]byte{cipherdatakey, iv, ciphertext}
	concat := bytes.Join(bufferslice, sep)
	encodelen := base64.StdEncoding.EncodedLen(len(concat))
	encdata := make([]byte, encodelen)
	base64.StdEncoding.Encode(encdata, concat)
	return encdata
}

func SplitEncFile(filedata []byte) ([]byte, []byte, []byte) {
	decodelen := base64.StdEncoding.DecodedLen(len(filedata))
	decodeddata := make([]byte, decodelen)
	base64.StdEncoding.Decode(decodeddata, filedata)
	returnslice := bytes.SplitN(decodeddata, sep, 3)
	key := returnslice[0]
	iv := returnslice[1]
	suffix := []byte{0}
	data := bytes.TrimSuffix(returnslice[2], suffix)
	return data, iv, key
}
