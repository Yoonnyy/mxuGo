package configuration

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	FileUploadDeactive    bool     `yaml:"file_upload_deactive"`
	UrlShorteningDeactive bool     `yaml:"url_shortening_deactive"`
	UrlRedirectDeactive   bool     `yaml:"url_redirect_deactive"`
	FileDownloadDeactive  bool     `yaml:"file_download_deactive"`
	DatabaseUrl           string   `yaml:"database_url"`
	SlugLength            int      `yaml:"slug_length"`
	Address               string   `taml:"address"`
	Port                  int      `yaml:"port"`
	UploadsFolder         string   `yaml:"uploads_folder"`
	MaxFileSize           int      `yaml:"max_file_size"`
	ForbiddenFileMimes    []string `yaml:"forbidden_file_mimes"`
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

	// create uploads directiory
	err = os.MkdirAll(config.UploadsFolder, os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf(
			`either the user does not have permission to create directiory 
or the uploads folder you provided is invalid
folder you provided: %v`,
			config.UploadsFolder))
	}
	if config.UploadsFolder[len(config.UploadsFolder)-1] == '/' {
		config.UploadsFolder = config.UploadsFolder[:len(config.UploadsFolder)-1]
	}
	return &config
}

func defaultIfNil(value, defaultValue interface{}) interface{} {
	if value == nil || value == "" || value == 0 {
		return defaultValue
	}
	return value
}
