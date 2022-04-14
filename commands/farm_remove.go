package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/czQery/instagram-bot/tools"
	"github.com/gookit/color"
	"github.com/imroc/req"
)

func Farm_remove() {
	// Load data
	file_load, _ := ioutil.ReadFile("data.json")
	type users_struct struct {
		Id   string `json:"id"`
		Time string `json:"time"`
	}
	var users []users_struct
	json.Unmarshal(file_load, &users)

	var (
		users_user     users_struct
		users_user_num int
		resp           *req.Resp
	)

	for range users {
		for users_user_num, users_user = range users {
			// Send POST unfollow request
			resp, _ = req.Post("https://www.instagram.com/web/friendships/"+users_user.Id+"/unfollow/", tools.Header)
			if resp.Response().StatusCode == 200 {
				tools.Log("Removed Id: " + color.HEX("FFAA00").Sprint(users_user.Id))

				// Remove user from followed users list
				users[users_user_num] = users[len(users)-1]
				users = users[:len(users)-1]
				break
			} else {
				tools.Log("Waiting...")
				for {
					// Try again to send POST unfollow request
					resp, _ = req.Post("https://www.instagram.com/web/friendships/"+users_user.Id+"/unfollow/", tools.Header)
					if resp.Response().StatusCode == 200 {
						tools.Log("Removed Id: " + color.HEX("FFAA00").Sprint(users_user.Id))

						// Remove user from followed users list
						users[users_user_num] = users[len(users)-1]
						users = users[:len(users)-1]
						break
					}
					tools.Sleep(300)
				}
				break
			}
		}
		tools.Sleep(3)
	}

	file_save, _ := json.Marshal(users)
	_ = ioutil.WriteFile("data.json", file_save, os.ModePerm)
	tools.Log("Done!")
	os.Exit(1)
}
