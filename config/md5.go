package config

type MD5 struct {
	Salt string `mapstructure:"salt" json:"salt" yaml:"salt"`
}
