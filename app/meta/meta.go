package meta

import "github.com/spf13/viper"

type Meta struct {
	App         string `json:"app"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

func Init() *Meta {
	var (
		env     = viper.GetString("APP_ENV")
		version = viper.GetString("APP_VERSION")
	)

	if env != "production" {
		version += "-" + env
	}

	return &Meta{
		App:         "datz",
		Version:     version,
		Environment: env,
	}
}

func (meta *Meta) GetApp() string {
	return meta.App
}

func (meta *Meta) GetVersion() string {
	return meta.Version
}

func (meta *Meta) GetAppVersion() string {
	return meta.GetApp() + "-" + meta.GetVersion()
}
func (meta *Meta) GetEnvironment() string {
	return meta.Environment
}
