package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/czQery/instagram-bot/tools"
	"github.com/gookit/color"
	"github.com/imroc/req"
)

func Farm_remove(user map[string]string, sessionid string, csrftoken string) {
	header := req.Header{
		"cookie":      "sessionid=" + sessionid + ";",
		"X-CSRFToken": csrftoken,
	}

	//? LOAD DATA
	file_1, _ := ioutil.ReadFile("data.json")
	type profiles_struct struct {
		Id   string `json:"id"`
		Time string `json:"time"`
	}
	var profiles []profiles_struct
	json.Unmarshal(file_1, &profiles)

	for range profiles {
		for i1, d1 := range profiles {
			//? GET ONE FOLLOWING
			ma1, _ := json.Marshal(d1)
			var profiles_user map[string]string
			json.Unmarshal(ma1, &profiles_user)

			resp_3, _ := req.Post("https://www.instagram.com/web/friendships/"+profiles_user["id"]+"/unfollow/", header)
			status := resp_3.Response().Status
			if status == "200 OK" {
				tools.Log("Removed Id: " + color.HEX("FFAA00").Sprint(profiles_user["id"]))
				profiles[i1] = profiles[len(profiles)-1]
				profiles = profiles[:len(profiles)-1]
				break
			} else {
				tools.Log("Waiting...")
				for {
					resp_4, _ := req.Post("https://www.instagram.com/web/friendships/"+profiles_user["id"]+"/unfollow/", header)
					status := resp_4.Response().Status
					if status == "200 OK" {
						tools.Log("Removed Id: " + color.HEX("FFAA00").Sprint(profiles_user["id"]))
						profiles[i1] = profiles[len(profiles)-1]
						profiles = profiles[:len(profiles)-1]
						break
					}
					tools.Sleep(300)
				}
				break
			}
		}
		tools.Sleep(3)
	}
	file_2, _ := json.Marshal(profiles)
	_ = ioutil.WriteFile("data.json", file_2, os.ModePerm)
	tools.Log("Done!")
	os.Exit(1)
}
