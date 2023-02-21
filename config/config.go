package config

import "github.com/spf13/viper"

type Config struct {
}

type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	var loaders []Loader
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}

func generateConfigFromViper(v *viper.Viper) Config {
	return Config{}
}
