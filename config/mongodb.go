package config

type Mongodb struct {
	Host   string `mapstructure:"host" json:"host" yaml:"host"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	DB     int    `mapstructure:"db" json:"db" yaml:"db"`
	Driver string `mapstructure:"driver" json:"driver" yaml:"driver"`
}
