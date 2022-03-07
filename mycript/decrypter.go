package mycript

import (
	"5.1/signature"
	"fmt"
	"io/ioutil"
	"os"
)

type Decryptor struct {
	fileHash   string
	hashString string
	fileSource string
	fileSigned string
	signature  signature.SignatureSha256
}

func NewDecoder(fileHash string, fileSource string, fileSigned string) (dec *Decryptor, err error) {
	hashString, err := ioutil.ReadFile(fileHash)
	if err != nil {
		return
	}

	dec = &Decryptor{fileHash: fileHash, hashString: string(hashString), fileSource: fileSource, fileSigned: fileSigned}
	return
}

//add validate()

//gen signature

//build signature from signFile.txt
func (dec *Decryptor) DecryptSha256() (err error) {
	file, err := os.Open(dec.fileSigned)
	defer file.Close()
	if err != nil {
		return
	}

	sign, err := signature.NewFileSourceFromSignatureSha256(file)
	if err != nil {
		return
	}

	fmt.Printf("sign %x \n", sign)

	dec.signature = sign
	return
}

//signature.Equals(signature)
func (dec Decryptor) Equality() {
	file, err := os.Open(dec.fileSource)
	if err != nil {
		return
	}
	defer file.Close()
	sig, err := signature.NewSignatureSha256FromFileSource(file, dec.hashString)
	if err != nil {
		return
	}
	sig.Equals(dec.signature)
}
