package main

var Config ConfigModel

type ConfigModel struct {
	User string `yaml:"user"`
	Secret string `yaml:"secret"`
	Dir string `yaml:"dir"`
	Slim SlimConfig `yaml:"Slim"`
	Watermark WatermarkConfig `yaml:"Watermark"`
}

type SlimConfig struct {
	IsEnable bool `yaml:"enable"`
	Pixels int `yaml:"compression threshold"`
}

type WatermarkConfig struct {
	Word struct {
		IsEnable bool `yaml:"enable"`
		Content string `yaml:"content"`
	}`yaml:"word"`

	Image struct {
		IsEnable bool `yaml:"enable"`
		Path string `yaml:"path"`
	}`yaml:"image"`
}
