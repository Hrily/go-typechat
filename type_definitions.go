package typechat

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type TypeDefinitions struct {
	definitions string
}

func NewTypeDefinitions(definitions string) *TypeDefinitions {
	return &TypeDefinitions{
		definitions: definitions,
	}
}

func NewTypeDefinitionsFromFile(filenames ...string) (*TypeDefinitions, error) {
	buffer := bytes.NewBuffer([]byte{})
	for _, filename := range filenames {
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
