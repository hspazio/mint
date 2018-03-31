package configurations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

// Dir where all application files are by default
var Dir = filepath.Join(homeDir(), ".mint")

// File containing the configurations
var File = filepath.Join(homeDir(), ".mint.json")

// Configurations are kept in a ".mint.json" file
type Configurations struct {
	Dir    string `json:"dir"`
	Editor string `json:"editor"`
}

func (c *Configurations) String() string {
	confB, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(confB)
}

// Init configuration file with default configs
func Init() error {
	_, err := GetAll()
	if err == nil {
		return fmt.Errorf("configs already initialized")
	}

	conf := &Configurations{Dir: Dir, Editor: os.Getenv("EDITOR")}
	if err := save(conf); err != nil {
		return err
	}

	return os.MkdirAll(conf.Dir, os.ModePerm)
}

// GetAll returns all configurations
func GetAll() (*Configurations, error) {
	conf := &Configurations{}
	b, err := ioutil.ReadFile(File)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(b, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

// Set a configuration by providing the key/value pair
func Set(key, value string) error {
	conf, err := GetAll()
	if err != nil {
		return err
	}
	switch key {
	case "dir":
		conf.Dir = value
	case "editor":
		conf.Editor = value
	default:
		return fmt.Errorf("not a valid configuration")
	}
	if err := save(conf); err != nil {
		return fmt.Errorf("could not save configurations file: %v", err)
	}
	return nil
}

func save(c *Configurations) error {
	conf, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(File, conf, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func homeDir() string {
	usr, err := user.Current()
	if err == nil {
		return usr.HomeDir
	}
	return "."
}
