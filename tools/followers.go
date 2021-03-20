package tools

import (
	"encoding/json"
	"strconv"

	"github.com/imroc/req"
)

func GetFollowers(userid string, sessionid string, csrftoken string) ([]map[string]map[string]string, string) {
	header := req.Header{
		"cookie":      "sessionid=" + sessionid + ";",
		"X-CSRFToken": csrftoken,
	}
	query := req.QueryParam{
		"query_hash": "5aefa9893005572d237da5068082d8d5",
		"variables":  `{"id":"` + userid + `","include_reel":false,"fetch_mutual":false,"first":5000}`,
	}
	resp_1, _ := req.Get("https://www.instagram.com/graphql/query/", header, query)
	var response_2 map[string]map[string]map[string]map[string]int
	var response_3 map[string]map[string]map[string]map[string]interface{}
	resp_1.ToJSON(&response_2)
	response_data_2 := response_2["data"]
	response_data_3 := response_data_2["user"]
	response_data_4 := response_data_3["edge_followed_by"]
	count := response_data_4["count"]

	resp_1.ToJSON(&response_3)
	response_data_5 := response_3["data"]
	response_data_6 := response_data_5["user"]
	response_data_7 := response_data_6["edge_followed_by"]
	response_data_8 := response_data_7["edges"]

	var target_list []map[string]map[string]string
	kk, _ := json.Marshal(response_data_8)
	json.Unmarshal(kk, &target_list)

	return target_list, strconv.Itoa(count)
}
