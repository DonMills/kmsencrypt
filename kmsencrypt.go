package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/DonMills/kmsencrypt/awsfuncs"
	"github.com/DonMills/kmsencrypt/encryption"
	"github.com/DonMills/kmsencrypt/errorhandle"
	"github.com/DonMills/kmsencrypt/filefuncs"

	"github.com/urfave/cli"
)

func decrypt(localfilename string, context string) {
	filedata, err := ioutil.ReadFile(localfilename)
	if err != nil {
		errorhandle.GenError(err)
	}
	encdata, iv, key := filefuncs.SplitEncFile(filedata)
	decryptkey := awsfuncs.DecryptKey(key, context)
	result := encryption.DecryptFile(encdata, iv, decryptkey)
	newfilename := strings.TrimSuffix(localfilename, ".kms")
	err2 := ioutil.WriteFile(newfilename, result, 0644)
	if err2 != nil {
		errorhandle.GenError(err)
	}
}

func encrypt(localfilename string, context string, cmkID string) {
	filedata, err := ioutil.ReadFile(localfilename)
	if err != nil {
		errorhandle.GenError(err)
	}
	cipherenvkey, plainenvkey := awsfuncs.GenerateEnvKey(cmkID, context)
	ciphertext, iv := encryption.EncryptFile(filedata, plainenvkey)
	result := filefuncs.CreateEncFile(ciphertext, iv, cipherenvkey)
	err2 := ioutil.WriteFile(localfilename+".kms", result, 0644)
	if err2 != nil {
		errorhandle.GenError(err)
	}
}

func main() {

	var cmkID string

	app := cli.NewApp()
	app.Name = "kmsencrypt"
	app.Usage = "Encrypt and decrypt files using KMS provided keys"
	app.HelpName = "kmsencrypt"
	app.UsageText = "kmsencrypt [command] {command specific options}"
	app.ArgsUsage = "kmsencrypt [command]"
	app.Version = "1.0rc"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Don Mills",
			Email: "don.mills@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "decrypt",
			Aliases:   []string{"d"},
			Usage:     "Decrypt a file encrypted with kmsencrypt",
			ArgsUsage: "[localfilename] [context]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					fmt.Println("Usage: kmsencrypt decrypt [localfilename] [context]")
					os.Exit(1)
				} else {
					decrypt(c.Args().Get(0), c.Args().Get(1))
				}
				return nil
			},
		},
		{
			Name:      "encrypt",
			Aliases:   []string{"e"},
			Usage:     "Generate a KMS key and encrypt a file",
			ArgsUsage: "[localfilename] [context]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "c",
					Usage:       "The customer master key id - can set with KMSENCRYPT_CMKID environment variable",
					EnvVar:      "KMSENCRYPT_CMKID",
					Destination: &cmkID,
				},
			},
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					fmt.Println("Usage: kmsencrypt encrypt [localfilename] [context] -c [customermasterkey]")
					os.Exit(1)
				} else {
					encrypt(c.Args().Get(0), c.Args().Get(1), cmkID)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
