package commands

import (
	"fmt"

	"github.com/czQery/instagram-bot/tools"
	"github.com/gookit/color"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

func Unfollow() {
	target_list, count := tools.GetFollowing(tools.User.Get("id").Str)

	tools.Log("Following: " + color.HEX("FFAA00").Sprint(count))

	var (
		resp                 *req.Resp
		target_user          gjson.Result
		target_user_id       string
		target_user_username string
	)

	// Loop through following users
	for _, target_user = range target_list {
		target_user_id = target_user.Get("node.id").Str
		target_user_username = target_user.Get("node.username").Str

		// Send POST unfollow request
		resp, _ = req.Post("https://www.instagram.com/web/friendships/"+target_user_id+"/unfollow/", tools.Header)
		fmt.Println(resp)
		if resp.Response().StatusCode == 200 {
			tools.Log("Removed: " + color.HEX("FFAA00").Sprint(target_user_username) + " Id: " + color.HEX("FFAA00").Sprint(target_user_id))
		} else {
			tools.Log("Waiting...")
			for {
				// Try again to send POST unfollow request
				resp, _ = req.Post("https://www.instagram.com/web/friendships/"+target_user_id+"/unfollow/", tools.Header)
				if resp.Response().StatusCode == 200 {
					tools.Log("Removed: " + color.HEX("FFAA00").Sprint(target_user_username) + " Id: " + color.HEX("FFAA00").Sprint(target_user_id))
					break
				}
				tools.Sleep(300)
			}
		}
		tools.Sleep(3)
	}
}
