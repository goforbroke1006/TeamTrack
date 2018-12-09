package config

import (
	"github.com/hashicorp/consul/api"
	"github.com/l-vitaly/consul"
)

type Configuration struct {
	Db struct {
		Host string `consul:"name:host;default:localhost"`
		Port string `consul:"name:port;default:5432"`
		Name string `consul:"name:name;"`
		User string `consul:"name:user;"`
		Pass string `consul:"name:pass;"`
	} `consul:"name:db;"`
}

func ReadFromConsul() (cfg *Configuration, err error) {
	var client consul.Client
	//client, err = consul.NewClientWithDefaultConfig()
	client, err = consul.NewClient(&api.Config{
		Address: "consul:8500",
	})
	err = client.LoadStruct("teamtrack", cfg)
	return
}

//func Read(filename string) (cfg Configuration, err error) {
//	var file *os.File
//	file, err = os.Open(filename)
//	defer file.Close()
//
//	if nil != err {
//		return
//	}
//	ext := path.Ext(filename)
//	switch ext {
//	case "json":
//		decoder := json.NewDecoder(file)
//		err = decoder.Decode(&cfg)
//		break
//	case "yaml":
//	case "yml":
//		d := yaml.NewDecoder(file)
//		err = d.Decode(&cfg)
//		break
//	default:
//		err = errors.New("unsupported configuration type")
//	}
//
//	return
//}
