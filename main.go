package main

import (
	"5.1/mycript"
	"flag"
	"log"
)

func main() {
	var fileSource, hashFile, outFile string

	flag.StringVar(&fileSource, "source-file", "", "File source")
	flag.StringVar(&hashFile, "hash-file", "", "File hash")
	flag.StringVar(&outFile, "out-file", "sign.txt", "File output")

	flag.Parse()
	action := flag.Args()[0]

	switch action {
	case "enc":

		encoder, err := mycript.NewEncoder(fileSource, hashFile)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = encoder.CreateSign()
		if err != nil {
			return
		}
		err = encoder.SaveToFile(outFile)
		if err != nil {
			return
		}

	case "dec":
		decoder, err := mycript.NewDecoder(hashFile, fileSource, outFile)
		if err != nil {
			return
		}

		err = decoder.DecryptSha256()
		if err != nil {
			return
		}
		decoder.Equality()
	default:
		log.Fatal("Use enc or dec param")
	}
}
