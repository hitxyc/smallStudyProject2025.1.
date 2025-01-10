package config

import "github.com/spf13/viper"

// Config储存用户配置信息
type Config struct {
	AppName  string         `mapstructure:"app_name"`
	Port     string         `mapstructure:"port"`
	Database DatabaseConfig `mapstructure:"database"`
}

// 数据库配置信息
type DatabaseConfig struct {
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	DatabaseName string `mapstructure:"database_name"`
}

// LoadConfig读取配置文件并返回Config对象
func LoadConfig() (*Config, error) {
	viper.SetConfigName("./config/config") //配置文件名(不带拓展名)
	viper.AddConfigPath(".")               //配置文件路径
	viper.SetConfigType("json")            //配置文件类型

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
