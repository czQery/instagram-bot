package tools

import (
	"github.com/gookit/color"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

func GetFollowing(userid string) ([]gjson.Result, string) {
	query := req.QueryParam{
		"query_hash": "3dec7e2c57367ef3da3d987d89f9dbc8",
		"variables":  `{"id":"` + userid + `","include_reel":false,"fetch_mutual":false,"first":5000}`,
	}
	resp, _ := req.Get("https://www.instagram.com/graphql/query/", Header, query)

	data := gjson.Parse(resp.String())
	if data.Get("status").Str != "ok" {
		Log(color.FgLightRed.Render("Obtaining followers has failed!"))
		return nil, ""
	}

	return data.Get("data.user.edge_follow.edges").Array(), data.Get("data.user.edge_follow.count").String()
}
