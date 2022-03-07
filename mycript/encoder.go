package mycript

import (
	"5.1/signature"
	"5.1/signature/contract"
	"io/ioutil"
	"os"
)

type Encoder struct {
	hashSign   string
	fileSource string
	signature  contract.Signature
}

func NewEncoder(fileSource string, fileHashSign string) (enc *Encoder, err error) {
	hashSting, err := ioutil.ReadFile(fileHashSign)
	if err != nil {
		return
	}
	enc = &Encoder{fileSource: fileSource, hashSign: string(hashSting)}
	return
}

func (enc *Encoder) EncryptSha256() (err error) {
	file, err := os.Open(enc.fileSource)
	defer file.Close()
	if err != nil {
		return
	}

	sign, err := signature.NewSignatureSha256FromFileSource(file, enc.hashSign)
	if err != nil {
		return
	}

	enc.signature = sign
	return
}

func (enc Encoder) SaveToFile(path string) (err error) {
	err = ioutil.WriteFile(path, enc.signature.SignatureByte(), 0644)
	return
}

func (enc *Encoder) CreateSign() (err error) {
	file, err := os.Open(enc.fileSource)
	if err != nil {
		return
	}
	defer file.Close()
	sig, err := signature.NewSignatureSha256FromFileSource(file, enc.hashSign)
	if err != nil {
		return
	}
	enc.signature = sig
	return
}
