package commands

import (
	"encoding/json"

	"github.com/czQery/instagram-bot/tools"
	"github.com/imroc/req"
)

func Unfollow(user map[string]string, sessionid string, csrftoken string) {
	target_list, count := tools.GetFollowing(user["id"], sessionid, csrftoken)

	tools.Log("Following: " + count)

	header := req.Header{
		"cookie":      "sessionid=" + sessionid + ";",
		"X-CSRFToken": csrftoken,
	}

	for _, dd := range target_list {
		uu, _ := json.Marshal(dd["node"])
		var target_user map[string]string
		json.Unmarshal(uu, &target_user)
		resp_3, _ := req.Post("https://www.instagram.com/web/friendships/"+target_user["id"]+"/unfollow/", header)
		status := resp_3.Response().Status
		if status == "200 OK" {
			tools.Log("Removed: " + target_user["username"] + " Id: " + target_user["id"])
		} else {
			tools.Log("Waiting...")
			for {
				resp_4, _ := req.Post("https://www.instagram.com/web/friendships/"+target_user["id"]+"/unfollow/", header)
				status := resp_4.Response().Status
				if status == "200 OK" {
					tools.Log("Removed: " + target_user["username"] + " Id: " + target_user["id"])
					break
				}
				tools.Sleep(300)
			}
		}
		tools.Sleep(3)
	}
}
