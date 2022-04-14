package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

var Config map[string]string

var User gjson.Result

var Header req.Header

func LoadConfig() {
	config_file, err := ioutil.ReadFile("config.json")
	if err != nil {
		Log(color.FgLightRed.Render("Config ERROR!"))
		os.Exit(1)
	} else {
		json.Unmarshal(config_file, &Config)
		Log(color.FgLightGreen.Render("Config loaded!"))
		Header = req.Header{
			"cookie":      "sessionid=" + Config["sessionid"] + ";csrftoken=" + Config["csrftoken"] + ";",
			"user-agent":  "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36",
			"x-csrftoken": Config["csrftoken"],
		}
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
