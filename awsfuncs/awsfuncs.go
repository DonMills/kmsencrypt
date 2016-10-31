// Package awsfuncs contains all the code that deals directly with AWS services
package awsfuncs

import (
	"github.com/DonMills/kmsencrypt/errorhandle"

	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

var kmssvc *kms.KMS

func init() {
	kmssvc = kms.New(session.New())
}

//GenerateEnvKey This function is used to generate KMS encryption keys for
//envelope encryption
func GenerateEnvKey(cmkID string, context string) ([]byte, []byte) {
	genparams := &kms.GenerateDataKeyInput{
		KeyId: aws.String(cmkID),
		EncryptionContext: map[string]*string{
			"Application": aws.String(context),
		},
		KeySpec: aws.String("AES_256"),
	}
	resp, err := kmssvc.GenerateDataKey(genparams)
	if err != nil {
		errorhandle.AWSError(err)
	}
	plainkey := resp.Plaintext
	cipherkey := resp.CiphertextBlob
	return cipherkey, plainkey
}

//DecryptKey does the actual KMS decryption of the stored key
func DecryptKey(output []byte, context string) []byte {
	keyparams := &kms.DecryptInput{
		CiphertextBlob: output, // Required
		EncryptionContext: map[string]*string{
			"Application": aws.String(context),
		},
	}

	plainkey, err := kmssvc.Decrypt(keyparams)
	if err != nil {
		errorhandle.AWSError(err)
	}
	decodelen := base64.StdEncoding.DecodedLen(len(plainkey.Plaintext))
	decodedplainkey := make([]byte, decodelen)
	base64.StdEncoding.Decode(decodedplainkey, plainkey.Plaintext)
	return plainkey.Plaintext
}
