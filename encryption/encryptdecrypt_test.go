package encryption

import (
	"fmt"
	"testing"
)

func TestEncryptDecryptCBCOne(t *testing.T) {
	data := []byte("Test String one")
	fmt.Printf("Original Data: %s\n", string(data))
	key := GenerateDataKey()
	ctext, iv := EncryptFile(data, key)
	result := DecryptFile(ctext, iv, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in Decryptfile/Encryptfile functions")
	}
}

func TestEncryptDecryptCBCTwo(t *testing.T) {
	data := []byte("abcd1234efgh5678")
	fmt.Printf("Original Data: %s\n", string(data))
	key := GenerateDataKey()
	ctext, iv := EncryptFile(data, key)
	result := DecryptFile(ctext, iv, key)
	fmt.Printf("Decrypted Data: %s\n", string(result))
	if string(data) != string(result) {
		t.Error("Decryption failed!  Error in Decryptfile/Encryptfile functions")
	}
}
