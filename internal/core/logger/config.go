package core_logger

type Config struct {
	Level string `envconfig:"LOGGER_LEVEL" required:"true"`
	Folder string `envconfig:"LOGGER_FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}
	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get logger config: %w", err)
		panic(err)
	}
	return config
}
