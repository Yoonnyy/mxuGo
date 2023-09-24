package configuration

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	FileUploadActive    bool     `yaml:"file_upload_active"`
	UrlShorteningActive bool     `yaml:"url_shortening_active"`
	DatabaseUrl         string   `yaml:"database_url"`
	SlugLength          int      `yaml:"slug_length"`
	Address             string   `taml:"address"`
	Port                int      `yaml:"port"`
	UploadsFolder       string   `yaml:"uploads_folder"`
	MaxFileSize         int      `yaml:"max_file_size"`
	ForbiddenFileMimes  []string `yaml:"forbidden_file_mimes"`
}

func ParseConfig() *Config {
	// read file
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		panic("Can't read config.yaml")
	}

	// unmarshal yaml
	config := Config{}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		panic("Can't decode config.yaml")
	}

	// validate
	if config.DatabaseUrl == "" {
		panic("empty database url")
	}
	config.MaxFileSize = defaultIfNil(config.MaxFileSize, 104857600).(int)
	config.Address = defaultIfNil(config.Address, "127.0.0.1").(string)
	config.Port = defaultIfNil(config.Port, 1315).(int)
	config.SlugLength = defaultIfNil(config.SlugLength, 6).(int)
	config.UploadsFolder = defaultIfNil(config.UploadsFolder, "uploads").(string)
	config.UrlShorteningActive = defaultIfNil(config.UrlShorteningActive, true).(bool)

	return &config
}

func defaultIfNil(value, defaultValue interface{}) interface{} {
	if value == nil || value == "" {
		return defaultValue
	}
	return defaultValue
}
