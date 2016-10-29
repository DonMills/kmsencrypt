package filefuncs

import (
	"bytes"
	"encoding/base64"
	//"github.com/DonMills/s3encrypt/errorhandle"
)

var sep = []byte{0, 1, 0, 1, 0, 1}

func CreateEncFile(ciphertext []byte, iv []byte, cipherdatakey []byte) []byte {
	//fmt.Printf("Key: %v\n, IV: %v\n, Data: %v\n", cipherdatakey, iv, ciphertext)
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
	//fmt.Printf("Data raw: %v\n", filedata)
	//fmt.Printf("Data decoded: %v\n", decodeddata)
	returnslice := bytes.SplitN(decodeddata, sep, 3)
	key := returnslice[0]
	iv := returnslice[1]
	suffix := []byte{0}
	//lendata := len(returnslice[2]) - 1
	data := bytes.TrimSuffix(returnslice[2], suffix)
	//fmt.Printf("Key: %v\n iv: %v", key, iv)
	//fmt.Printf("Data: %v\n", data)
	return data, iv, key
}
