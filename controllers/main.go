package controllers

import "github.com/Yoonnyy/GoMxu/configuration"

var config configuration.Config

func Init(cfg configuration.Config) {
	config = cfg
}
