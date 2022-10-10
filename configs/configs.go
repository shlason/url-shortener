package configs

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type server struct {
	Host      string
	Port      string
	FeWrokDir string
}

type database struct {
	Dialect  string
	Username string
	Password string
	Protocol string
	Host     string
	Name     string
	Options  string
}

var Server = server{}
var Database = database{}

func init() {
	cfg, err := ini.Load("configs.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	Server = server{
		Host:      cfg.Section("server").Key("host").String(),
		Port:      cfg.Section("server").Key("port").String(),
		FeWrokDir: cfg.Section("server").Key("feWorkDir").String(),
	}
	Database = database{
		Dialect:  cfg.Section("database").Key("dialect").String(),
		Username: cfg.Section("database").Key("username").String(),
		Password: cfg.Section("database").Key("password").String(),
		Protocol: cfg.Section("database").Key("protocol").String(),
		Host:     cfg.Section("database").Key("host").String(),
		Name:     cfg.Section("database").Key("name").String(),
		Options:  cfg.Section("database").Key("options").String(),
	}
}
