package config

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

var cnfg Configuration

// Configuration modelo que contendrá un mapa de
// string/interface para leer y almacenar la data
// del archivo de configuración tipo json
type Configuration struct {
	data map[string]interface{}
}

func New(fullpath string) (*Configuration, error) {
	b, err := LoadFile(fullpath)
	if err != nil {
		return nil, err
	}

	err = LoadBytes(b, &cnfg)
	if err != nil {
		return nil, err
	}

	return &cnfg, nil
}

func LoadFile(fullpath string) ([]byte, error) {
	f, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return f, err
	}

	return f, nil
}

func LoadBytes(d []byte, c *Configuration) error {
	err := json.Unmarshal(d, &c.data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configuration) Validate(names ...string) error {
	for _, v := range names {
		_, ok := c.data[v]
		if !ok {
			return errors.New(fmt.Sprintf("no existe el campo %s", v))
		}
	}

	return nil
}

// Get devuelve el valor del campo si existe, tipo string
func (c *Configuration) Get(name string) (string, error) {
	v, ok := c.data[name].(string)
	if !ok {
		return "", errors.New(fmt.Sprintf("no existe el campo %s", name))
	}

	return v, nil
}

// GetInt devuelve el valor del campo si existe, tipo int
func (c *Configuration) GetInt(name string) (int, error) {
	v, ok := c.data[name].(float64)
	if !ok {
		return 0, errors.New(fmt.Sprintf("no existe el campo %s", name))
	}

	return int(v), nil
}
