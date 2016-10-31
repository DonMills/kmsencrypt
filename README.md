# kmsencrypt
A tool designed to do KMS based envelope encryption of files.
___
## Decryption
1. takes a file {filename.kms} that you have encrypted with this program
2. base64 decodes the file and extracts the IV and encryption key from the prefix
2. decrypts the encryption key with the proper EncryptionContext using KMS
3. then takes that key and unencrypts the data
4. saves the data in a local file of the name {filename} 

## Encryption
1. This takes the file {filename},
2. generates a KMS encryption key tied to a supplied EncryptionContext value and KMS Customer Master Key
3. encrypts the file with the encryption key
4. prepends the encrypted key and IV on to the file
5. base64 encodes the new file
6. saves it as {filename.kms}

___

## How to build:

#### git clone into the $GOPATH/src/github.com/DonMills directory
```
mkdir $GOPATH/src/github.com/DonMills
cd $GOPATH/src/github.com/DonMills
git clone https://github.com/DonMills/kmsencrypt.git
```
_or_
```
go get github.com/DonMills/kmsencrypt
```

#### This tool requires the "aws-sdk-go" and the ["urfave/cli"](https://github.com/urfave/cli) packages be installed.
```
go get github.com/aws/aws-sdk-go/
go get github.com/urfave/cli
```
Alternatively, if you have [glide](https://github.com/Masterminds/glide) installed, you can just get the deps like this:
```
glide up
```

#### Then just build or install it...
```
go install
```
_or_
```bash
go build -o kmsencrypt
./kmsencrypt 
```
## Usage Notes
The tool has a full help system, but in general usage is 
```
 kmsencrypt [command] {command specific options}
```
where commands are 
```
kmsencrypt encrypt [localfilename] [context]
OPTIONS:
   -c value  The customer master key id - can set with KMSENCRYPT_CMKID environment variable [$KMSENCRYPT_CMKID]
```
or
```
kmsencrypt decrypt [localfilename] [context]
```
#### Dealing with an "AWS Error: NoCredentialProviders" error or needing ~/.aws/config
In some situations (like needing a STS token to work on an environment) or if you have entries in your ~/.aws/config file that are needed, you may need to set the following environment variable:
```
AWS_SDK_LOAD_CONFIG=1
```
This is a function of the aws sdk for go discussed here: http://docs.aws.amazon.com/sdk-for-go/api/aws/session/
## Mac installation via homebrew
New!  Now you can install on a mac by using homebrew.
```
brew install DonMills/tools/kmsencrypt
```
