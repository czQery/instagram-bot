package tools

import (
	"os"

	"github.com/gookit/color"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

func GetUser() {
	query := req.QueryParam{
		"query_hash": "b1245d9d251dff47d91080fbdd6b274a",
	}
	resp, err := req.Get("https://www.instagram.com/graphql/query/", Header, query)
	if err != nil || resp.Response().StatusCode != 200 {
		Log(color.FgLightRed.Render("Login request failed!"))
		os.Exit(1)
	}

	data := gjson.Parse(resp.String())
	if data.Get("status").Str != "ok" || data.Get("data.user.id").Str == "" {
		Log(color.FgLightRed.Render("Incorrect login information!"))
		os.Exit(1)
	}

	User = data.Get("data.user")
}
