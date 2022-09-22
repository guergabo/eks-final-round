package utils

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadJSONFile(filename string) ([]byte, error) {
	jsonFileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFileReader.Close()

	byteValue, err := ioutil.ReadAll(jsonFileReader)
	if err != nil {
		return nil, err
	}

	return byteValue, nil
}

func MustRead(v []byte, err error) []byte {
	if err != nil {
		log.Fatalln(err)
	}
	return v
}
