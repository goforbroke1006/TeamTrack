package config

import (
	"encoding/json"
	"os"
	"path"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Db struct {
		Host string
		Port uint64
		Name string
		User string
		Pass string
	}
}

func Read(filename string) (cfg Configuration, err error) {
	var file *os.File
	file, err = os.Open(filename)
	defer file.Close()

	if nil != err {
		return
	}
	ext := path.Ext(filename)
	switch ext {
	case "json":
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&cfg)
		break
	case "yaml":
	case "yml":
		d := yaml.NewDecoder(file)
		err = d.Decode(&cfg)
		break
	default:
		err = errors.New("unsupported configuration type")
	}

	return
}
