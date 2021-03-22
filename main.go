package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/czQery/instagram-bot/commands"
	"github.com/czQery/instagram-bot/tools"
)

func main() {
	fmt.Println("[Instagram-bot][v1.5]")
	fmt.Println("[Created by czQery]")

	tools.LoadConfig()

	sessionid, _ := tools.Config["sessionid"]
	csrftoken, _ := tools.Config["csrftoken"]

	if sessionid == "" && csrftoken == "" {
		tools.Log("Config ERROR!")
		os.Exit(1)
	} else {
		//? GET USER
		user := tools.GetUser(sessionid, csrftoken)
		tools.Log("User-name: " + user["username"])
		tools.Log("User-id: " + user["id"])
		fmt.Println("----------")
		//? PRINT COMMANDS
		tools.Log("Please select mode!")
		tools.Log("1. Unfollow all accounts")
		tools.Log("2. Farm followers")
		tools.Log("3. Unfollow all temporary farming accounts")
		//? GET USER INPUT
		var input0 string
		fmt.Scanln(&input0)
		fmt.Println("----------")

		//? SELECT COMMAND
		switch command_id, _ := strconv.Atoi(input0); command_id {
		case 1:
			tools.Log("Selected mode: 1 (Unfollow all accounts)")
			commands.Unfollow(user, sessionid, csrftoken)
		case 2:
			tools.Log("Selected mode: 2 (Farm followers)")
			commands.Farm(user, sessionid, csrftoken)
		case 3:
			tools.Log("Selected mode: 3 (Unfollow all temporary farming accounts)")
			commands.Farm_remove(user, sessionid, csrftoken)
		default:
			tools.Log("Unknown mode!")
			os.Exit(1)
		}
	}
}
