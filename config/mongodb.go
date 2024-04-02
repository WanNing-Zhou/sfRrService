package config

type Mongodb struct {
	Host                string `mapstructure:"host" json:"host" yaml:"host"`
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`
	DB                  int    `mapstructure:"db" json:"db" yaml:"db"`
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`
	Database            string `mapstructure:"database" json:"database" yaml:"database"`
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"`
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`
}
