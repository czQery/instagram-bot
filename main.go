package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/czQery/instagram-bot/commands"
	"github.com/czQery/instagram-bot/tools"
	"github.com/gookit/color"
)

func main() {
	fmt.Println(color.FgDarkGray.Render("[") + color.FgLightRed.Render("Instagram-bot") + color.FgDarkGray.Render("][") + color.HEX("FFAA00").Sprint("v1.7") + color.FgDarkGray.Render("]"))
	fmt.Println(color.FgDarkGray.Render("[") + color.FgLightRed.Render("Created by czQery") + color.FgDarkGray.Render("]"))

	tools.LoadConfig()

	if tools.Config["sessionid"] == "" && tools.Config["csrftoken"] == "" {
		tools.Log("Config ERROR!")
		os.Exit(1)
	} else {
		// GET USER
		tools.GetUser()
		tools.Log("User-name: " + color.HEX("FFAA00").Sprint(tools.User.Get("username").Str))
		tools.Log("User-id: " + color.HEX("FFAA00").Sprint(tools.User.Get("id").Str))
		fmt.Println("----------")
		// PRINT COMMANDS
		tools.Log("Please select mode!")
		tools.Log("1. Unfollow all accounts")
		tools.Log("2. Farm followers")
		tools.Log("3. Unfollow all temporary farming accounts")
		// GET USER INPUT
		var input0 string
		fmt.Scanln(&input0)
		fmt.Println("----------")

		// SELECT COMMAND
		switch command_id, _ := strconv.Atoi(input0); command_id {
		case 1:
			tools.Log(color.FgLightGreen.Render("Selected mode: 1 (Unfollow all accounts)"))
			commands.Unfollow()
		case 2:
			tools.Log(color.FgLightGreen.Render("Selected mode: 2 (Farm followers)"))
			commands.Farm()
		case 3:
			tools.Log(color.FgLightGreen.Render("Selected mode: 3 (Unfollow all temporary farming accounts)"))
			commands.Farm_remove()
		default:
			tools.Log(color.FgLightRed.Render("Unknown mode!"))
			os.Exit(1)
		}
	}
}
