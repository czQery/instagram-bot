package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/czQery/instagram-bot/tools"
	"github.com/gookit/color"
	"github.com/imroc/req"
	"github.com/tidwall/gjson"
)

func Farm() {
	// PRINT OPTIONS
	tools.Log("Please select which community you are targeting!")
	tools.Log("1. Custom")
	tools.Log("2. Games")
	tools.Log("3. Memes")
	tools.Log("4. Thicc girls")
	tools.Log("5. Thicc boys")
	tools.Log("6. Cars")
	tools.Log("7. Tech")
	tools.Log("8. DIY")
	tools.Log("9. Food")
	tools.Log("10. Actually good photos")
	// GET USER INPUT
	var input1 string
	fmt.Scanln(&input1)
	fmt.Println("----------")

	custom := false
	var community []string
	switch community_id, _ := strconv.Atoi(input1); community_id {
	case 1:
		tools.Log(color.FgLightGreen.Render("Selected community: 1 (Custom)"))
		custom = true
		tools.Log("Please enter custom id!")
		var input2 string
		fmt.Scanln(&input2)
		fmt.Println("----------")
		community = append(community, input2)

	case 2:
		tools.Log(color.FgLightGreen.Render("Selected community: 2 (Games)"))
		community = append(community, "1192899491")  //Fortnite
		community = append(community, "46218147")    //IGN
		community = append(community, "2077685663")  //Ninja
		community = append(community, "1725401510")  //Shroud
		community = append(community, "7650489800")  //Dream
		community = append(community, "12130726473") //Tommy
		community = append(community, "529464688")   //S1mple
	case 3:
		tools.Log(color.FgLightGreen.Render("Selected community: 3 (Memes)"))
		community = append(community, "18065136551") //Daily Dose Of Internet
		community = append(community, "32039297384") //Phailure
		community = append(community, "9978699610")  //Weebshit
		community = append(community, "32345649705") //Homer.explains
		community = append(community, "7793657210")  //Unsquidable
		community = append(community, "25566185512") //Comics.collector
		community = append(community, "13506898")    //Pewdiepie
		community = append(community, "10033548262") //Pewmemes
	case 4:
		tools.Log(color.FgLightGreen.Render("Selected community: 4 (Thicc girls)"))
		community = append(community, "11389317073") //Belle.delphiny
		community = append(community, "938811795")   //Amouranth
		community = append(community, "336942508")   //Poki
		community = append(community, "182541893")   //Angievarona
		community = append(community, "614810319")   //Lily.adrianne
		community = append(community, "3055748987")  //Jennatwitch
		community = append(community, "188331362")   //Evamenta
		community = append(community, "5152138052")  //Pandorakaaki
	case 5:
		tools.Log(color.FgLightGreen.Render("Selected community: 5 (Thicc boys)"))
		community = append(community, "1929749275") //Rice
		community = append(community, "2278169415") //Mrbeast
		community = append(community, "6860189")    //Justinbieber
		community = append(community, "1414857648") //Tom Ellis
		community = append(community, "1660092007") //Radoslav_raychev
	case 6:
		tools.Log(color.FgLightGreen.Render("Selected community: 6 (Cars)"))
		community = append(community, "4086860751") //Mdc.media
		community = append(community, "2272114078") //Supercarsbuzz
		community = append(community, "297604134")  //Tesla
		community = append(community, "2074583971") //Carcrazy.india
		community = append(community, "25749975")   //Mercedesbenz
		community = append(community, "43109246")   //Bmw
		community = append(community, "2465409402") //Supercar
	case 7:
		tools.Log(color.FgLightGreen.Render("Selected community: 7 (Tech)"))
		community = append(community, "297604134")   //Tesla
		community = append(community, "20311520")    //SpaceX
		community = append(community, "4236812322")  //Elon musk
		community = append(community, "5821462185")  //Apple
		community = append(community, "30047490566") //Samsung
		community = append(community, "360157628")   //Nvidiageforce
		community = append(community, "14673726")    //Amd
	case 8:
		tools.Log(color.FgLightGreen.Render("Selected community: 8 (DIY)"))
		community = append(community, "18328422") //5.min.crafts
	case 9:
		tools.Log(color.FgLightGreen.Render("Selected community: 9 (Food)"))
		community = append(community, "1573550968")  //Chefincamicia
		community = append(community, "214024091")   //Soniaperonac
		community = append(community, "1584854974")  //Dabizdiverxo
		community = append(community, "39629390339") //Ketosmart_
		community = append(community, "175473620")   //Breakfastnbowls
	case 10:
		tools.Log(color.FgLightGreen.Render("Selected community: 10 (Actually good photos)"))
		community = append(community, "1987493425") //Airpixels
		community = append(community, "9868480")    //Thiswildidea
		community = append(community, "36045182")   //Chrisburkard
		community = append(community, "174143945")  //Hannes_becker
	default:
		tools.Log(color.FgLightRed.Render("Unknown community!"))
		os.Exit(1)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	tools.Log(color.FgLightGreen.Render("Starting..."))

	type users_struct struct {
		Id   string `json:"id"`
		Time int64  `json:"time"`
	}

	var (
		resp      *req.Resp
		err       error
		file_load []byte
		file_save []byte
		// Saved users in data.json
		users          []users_struct
		users_count    string
		users_user     users_struct
		users_user_num int
		// Target profile for grabbing users
		target_profile          string
		target_followers        []gjson.Result
		target_user             gjson.Result
		target_user_id          string
		target_user_username    string
		i1                      int
		d1                      gjson.Result
		followers               []gjson.Result
		followers_user          gjson.Result
		followers_user_id       string
		followers_user_username string
		followers_count         string
	)

	for true {

		// Load data
		file_load, _ = ioutil.ReadFile("data.json")
		json.Unmarshal(file_load, &users)

		// Follow random 5 users from random profile
		for i1 = 0; i1 < 5; i1++ {

			// Check if you are following less than 100 users
			if len(users) >= 100 {
				break
			}

			// Select random profile
			target_profile = community[rand.Intn(len(community))]
			target_followers, _ = tools.GetFollowers(target_profile)

			// Check if you have custom list of profiles
			if len(target_followers) < 1 {
				if custom {
					tools.Log(color.FgLightRed.Render("Bad custom id!"))
					os.Exit(1)
				} else {
					tools.Log(color.FgLightRed.Render("Bad profile id in community preset: " + input1))
					os.Exit(1)
				}
			}

			// Declare target user
			target_user = target_followers[rand.Intn(len(target_followers))].Get("node")
			target_user_id = target_user.Get("id").Str
			target_user_username = target_user.Get("username").Str

			// Send POST follow request
			resp, err = req.Post("https://www.instagram.com/web/friendships/"+target_user.Get("id").Str+"/follow/", tools.Header)
			if err == nil {
				if resp.Response().StatusCode == 200 {
					tools.Log("Followed: " + color.HEX("FFAA00").Sprint(target_user_username) + " Id: " + color.HEX("FFAA00").Sprint(target_user_id))

					// Add user to list of followed users
					users = append(users, users_struct{Id: target_user_id, Time: time.Now().Unix()})
				} else {
					tools.Log(color.FgLightRed.Render("Follow failed"))
				}
			} else {
				tools.Log(color.FgLightRed.Render("Follow failed"))
			}
			tools.Sleep(5)
		}

		// Save data
		file_save, _ = json.Marshal(users)
		ioutil.WriteFile("data.json", file_save, os.ModePerm)

		// Check 6 users if they have already started following me back
		for i1 = 0; i1 < 6; i1++ {

			// Get my followers
			followers, _ = tools.GetFollowers(tools.User.Get("id").Str)
			if len(followers) < 1 {
				tools.GetUser()
			}

			// Loop through my followers
			for _, d1 = range followers {
				// GET ONE FOLLOWER
				followers_user = d1.Get("node")

				// Loop through followed users
				for users_user_num, users_user = range users {

					// Declare follower
					followers_user_id = followers_user.Get("id").Str
					followers_user_username = followers_user.Get("username").Str

					// Check id of follower and following user
					if followers_user_id == users_user.Id {

						// Send POST unfollow request
						resp, err = req.Post("https://www.instagram.com/web/friendships/"+users_user.Id+"/unfollow/", tools.Header)
						if err == nil {
							if resp.Response().StatusCode == 200 {

								// Remove user from followed users list
								users[users_user_num] = users[len(users)-1]
								users = users[:len(users)-1]
								tools.Log(color.FgLightGreen.Render("New Follower: ") + color.HEX("FFAA00").Sprint(followers_user_username) + color.FgLightGreen.Render(" Id: ") + color.HEX("FFAA00").Sprint(followers_user_id))
								_, followers_count = tools.GetFollowers(tools.User.Get("id").Str)
								_, users_count = tools.GetFollowing(tools.User.Get("id").Str)
								tools.Log("Followers: " + color.HEX("FFAA00").Sprint(followers_count))
								tools.Log("Following: " + color.HEX("FFAA00").Sprint(users_count))
								break
							}
						}
					}

					// Check users followed for more than 24 hours
					if time.Now().UTC().Sub(time.Unix(users_user.Time, 0)).Hours() >= 24 {

						// Send POST unfollow request
						resp, err = req.Post("https://www.instagram.com/web/friendships/"+users_user.Id+"/unfollow/", tools.Header)
						if err == nil {
							if resp.Response().StatusCode == 200 {

								// Remove user from followed users list
								users[users_user_num] = users[len(users)-1]
								users = users[:len(users)-1]
								tools.Log("Unfollowed after 24h Id: " + color.HEX("FFAA00").Sprint(users_user.Id))
								break
							}
						}
					}
				}
			}

			// Save data
			file_save, _ = json.Marshal(users)
			_ = ioutil.WriteFile("data.json", file_save, os.ModePerm)
			tools.Sleep(300)
		}
	}
}
