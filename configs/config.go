package configs

import "github.com/spf13/viper"

func initViper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	configErr := v.ReadInConfig()
	if configErr != nil {
		panic("loading config error:" + configErr.Error())
	}

	return v
}

func GetConfig(key string) string {
	config := initViper()

	return config.GetString(key)
}
