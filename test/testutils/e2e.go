package testutils

import "io/ioutil"

// GetBytesFromFile ...
func GetBytesFromFile(path string) []byte {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return b
}
