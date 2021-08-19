package main

var Config ConfigModel

type ConfigModel struct {
	User string `yaml:"user"`
	Secret string `yaml:"secret"`
	Dir string `yaml:"dir"`
}
