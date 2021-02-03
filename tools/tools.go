package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var Config map[string]string

func LoadConfig() {
	config_file, err := ioutil.ReadFile("config.json")
	if err != nil {
		Log("Config ERROR!")
		os.Exit(1)
	} else {
		json.Unmarshal(config_file, &Config)
		Log("Config loaded!")
	}
}

func Log(message string) {
	fmt.Println("[" + time.Now().Format("15:04:05") + "] - " + message)
}
