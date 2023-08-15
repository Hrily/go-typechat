package typechat

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

// TypeDefinitions are the GoLang type definitions used to translate natural
// language into JSON
type TypeDefinitions struct {
	definitions string
}

// NewTypeDefinitions from given definitions string
func NewTypeDefinitions(definitions string) *TypeDefinitions {
	return &TypeDefinitions{
		definitions: definitions,
	}
}

// NewTypeDefinitionsFromFile reads type definitions from given files
func NewTypeDefinitionsFromFile(files ...string) (*TypeDefinitions, error) {
	buffer := bytes.NewBuffer([]byte{})
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			return nil, errors.Wrap(err, "failed to open type definitions file")
		}

		defer file.Close()
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read type definitions file")
		}
		buffer.Write(bytes)
	}

	return NewTypeDefinitions(buffer.String()), nil
}
