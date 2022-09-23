package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadJSONFile(filename string) ([]byte, error) {
	jsonFileReader, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFileReader.Close()

	bs, err := ioutil.ReadAll(jsonFileReader)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func WriteJSONFile(filename string, data interface{}) error {
	fileReader, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	if err != nil {
		return err
	}
	defer fileReader.Close()

	bs, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	_, err = fileReader.Write(bs)
	if err != nil {
		return err
	}

	return nil
}

func GetCLIArgs() []string {
	return os.Args[1:]
}
