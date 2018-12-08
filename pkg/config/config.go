package config

import (
	"encoding/json"
	"github.com/l-vitaly/consul"
	"os"
	"path"

	"github.com/hashicorp/consul/api"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Db struct {
		Host string `consul:"default:localhost"`
		Port uint64 `consul:"default:5432"`
		Name string `consul:""`
		User string `consul:""`
		Pass string `consul:""`
	}
}

func ReadFromConsul() (cfg *Configuration, err error) {
	var client consul.Client
	client, err = consul.NewClientWithDefaultConfig()
	err = client.LoadStruct("teamtrack", cfg)
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
