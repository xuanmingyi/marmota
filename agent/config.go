package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type conf_server struct {
	Uri string `yaml:"uri"`
}

type conf_agent struct {
	Interval int64             `yaml:"interval"`
	Uuid     string            `yaml:"uuid"`
	Name     string            `yaml:"name"`
	Desc     string            `yaml:"desc"`
	Metadata map[string]string `yaml:"metadata"`
}

type conf struct {
	Server conf_server `yaml:"server"`
	Agent  conf_agent  `yaml:"agent"`
}

func ReadConfig(path string) (c *conf, err error) {
	c = new(conf)

	content, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(content, c)
	if err != nil {
		panic(err)
	}

	return c, nil
}
