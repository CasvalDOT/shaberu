package formatter

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

type formatter struct {
	name string
}

type formatConf struct {
	Binary string   `json:"binary"`
	Args   []string `json:"args"`
}

type IFormatter interface {
	Format(*string)
	load() formatConf
}

func (f *formatter) Format(filePath *string) {
	conf := f.load()

	commandArgs := []string{}
	commandArgs = append(commandArgs, conf.Args...)
	commandArgs = append(commandArgs, *filePath)

	cmd := exec.Command(conf.Binary, commandArgs...)

	cmd.Run()
}

func (f *formatter) load() formatConf {

	homePath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Cannot get user home directory:", err.Error())
		os.Exit(2)
	}

	confPath := homePath + "/.config/translater/formatters.json"

	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		os.Create(confPath)
	}

	content, err := ioutil.ReadFile(confPath)
	if err != nil {
		fmt.Println("Cannot get formatter configuration:", err.Error())
		os.Exit(2)
	}

	var configuration map[string]formatConf
	json.Unmarshal(content, &configuration)

	return configuration[f.name]
}

func New(name string) IFormatter {
	return &formatter{
		name: name,
	}
}
