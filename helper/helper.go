package helper

import (
	"bytes"
	"encoding/gob"
)

// Serialize ...
func Serialize(s interface{}) ([]byte, error) {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(s)
	if err != nil {
		return nil, err
	}

	return res.Bytes(), nil
}

// Deserialize ...
func Deserialize(data []byte) (interface{}, error) {
	var d interface{}
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
