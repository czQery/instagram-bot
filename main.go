package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/czQery/instagram-bot/tools"
	"github.com/imroc/req"
)

func main() {
	fmt.Println("[Instagram-bot][v1]")
	fmt.Println("[Created by czQery]")

	tools.LoadConfig()

	sessionid, _ := tools.Config["sessionid"]
	csrftoken, _ := tools.Config["csrftoken"]

	if sessionid == "" && csrftoken == "" {
		tools.Log("Config ERROR!")
		os.Exit(1)
	} else {

		//! GET USER
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
			tools.Log("Incorrect login information!")
			os.Exit(1)
		}
		response_data_1 := response_1["data"]
		user := response_data_1["user"]
		tools.Log("User-name: " + user["username"])
		tools.Log("User-id: " + user["id"])

		//! GET FOLLOWING LIST
		query = req.QueryParam{
			"query_hash": "3dec7e2c57367ef3da3d987d89f9dbc8",
			"variables":  `{"id":"` + user["id"] + `","include_reel":false,"fetch_mutual":false,"first":5000}`,
		}
		resp_2, _ := req.Get("https://www.instagram.com/graphql/query/", header, query)
		var response_2 map[string]map[string]map[string]map[string]int
		var response_3 map[string]map[string]map[string]map[string]interface{}
		resp_2.ToJSON(&response_2)
		response_data_2 := response_2["data"]
		response_data_3 := response_data_2["user"]
		response_data_4 := response_data_3["edge_follow"]
		count := response_data_4["count"]
		tools.Log("Following: " + strconv.Itoa(count))

		resp_2.ToJSON(&response_3)
		response_data_5 := response_3["data"]
		response_data_6 := response_data_5["user"]
		response_data_7 := response_data_6["edge_follow"]
		response_data_8 := response_data_7["edges"]

		var target_list []map[string]map[string]string
		kk, _ := json.Marshal(response_data_8)
		json.Unmarshal(kk, &target_list)

		fmt.Println("----------")
		for _, dd := range target_list {
			uu, _ := json.Marshal(dd["node"])
			var target_user map[string]string
			json.Unmarshal(uu, &target_user)
			resp_3, _ := req.Post("https://www.instagram.com/web/friendships/"+target_user["id"]+"/unfollow/", header)
			status := resp_3.Response().Status
			if status == "200 OK" {
				tools.Log("Removed: " + target_user["username"] + " Id: " + target_user["id"])
			} else {
				tools.Log(status + " Error: " + target_user["username"] + " Id: " + target_user["id"])
			}
		}
	}
}
