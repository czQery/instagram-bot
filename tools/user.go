package tools

import (
	"os"

	"github.com/imroc/req"
)

func GetUser(sessionid string, csrftoken string) map[string]string {
	header := req.Header{
		"cookie":      "sessionid=" + sessionid + ";",
		"X-CSRFToken": csrftoken,
	}
	query := req.QueryParam{
		"query_hash": "b1245d9d251dff47d91080fbdd6b274a",
	}
	resp_1, _ := req.Get("https://www.instagram.com/graphql/query/", header, query)
	var response_1 map[string]map[string]map[string]string
	resp_1.ToJSON(&response_1)
	if response_1 == nil {
		Log("Incorrect login information!")
		os.Exit(1)
	}
	response_data_1 := response_1["data"]
	user := response_data_1["user"]
	return user
}
