package filefuncs

import (
	"bytes"
	"encoding/base64"
)

var sep = []byte{0, 1, 0, 1, 0, 1}

//CreateEncFile takes the key, iv, and encrypted data and concatenates it to a file
func CreateEncFile(ciphertext []byte, iv []byte, cipherdatakey []byte) []byte {
	bufferslice := [][]byte{cipherdatakey, iv, ciphertext}
	concat := bytes.Join(bufferslice, sep)
	encodelen := base64.RawStdEncoding.EncodedLen(len(concat))
	encdata := make([]byte, encodelen)
	base64.RawStdEncoding.Encode(encdata, concat)
	return encdata
}

//SplitEncFile takes the concatenated file and splits out the key, iv, and data
func SplitEncFile(filedata []byte) ([]byte, []byte, []byte) {
	decodelen := base64.RawStdEncoding.DecodedLen(len(filedata))
	decodeddata := make([]byte, decodelen)
	base64.RawStdEncoding.Decode(decodeddata, filedata)
	returnslice := bytes.SplitN(decodeddata, sep, 3)
	key := returnslice[0]
	iv := returnslice[1]
	data := returnslice[2]
	return data, iv, key
}
