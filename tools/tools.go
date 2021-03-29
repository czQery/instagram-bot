package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gookit/color"
)

var Config map[string]string

func LoadConfig() {
	config_file, err := ioutil.ReadFile("config.json")
	if err != nil {
		Log(color.FgLightRed.Render("Config ERROR!"))
		os.Exit(1)
	} else {
		json.Unmarshal(config_file, &Config)
		Log(color.FgLightGreen.Render("Config loaded!"))
	}
}

func Log(message string) {
	b := color.FgDarkGray.Render
	n := color.FgLightRed.Render
	//h := color.HEX("FFAA00").Sprint
	fmt.Println(b("[") + n(time.Now().Format("15:04:05")) + b("] ") + message)
}

func Sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}
