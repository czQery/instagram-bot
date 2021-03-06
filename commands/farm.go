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
)

func Farm(user map[string]string, sessionid string, csrftoken string) {
	//? PRINT OPTIONS
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
	//? GET USER INPUT
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

	header := req.Header{
		"cookie":      "sessionid=" + sessionid + ";",
		"X-CSRFToken": csrftoken,
	}
	rand.Seed(time.Now().UTC().UnixNano())
	tools.Log(color.FgLightGreen.Render("Starting..."))

	for true {

		//? LOAD DATA
		file_1, _ := ioutil.ReadFile("data.json")

		type profiles_struct struct {
			Id   string `json:"id"`
			Time string `json:"time"`
		}
		var profiles []profiles_struct
		json.Unmarshal(file_1, &profiles)

		//? SELECT RANDOM PROFILE AND FOLLOW RANDOM FOLLOWER
		for i := 0; i < 5; i++ {
			if len(profiles) >= 100 {
				break
			}

			target := community[rand.Intn(len(community))]
			followers, _ := tools.GetFollowers(target, sessionid, csrftoken)
			if len(followers) < 1 {
				tools.GetUser(sessionid, csrftoken)
				if custom {
					tools.Log(color.FgLightRed.Render("Bad custom id!"))
					os.Exit(1)
				} else {
					tools.Log(color.FgLightRed.Render("Bad profile id in community preset: " + input1))
					os.Exit(1)
				}
			}

			random_follower := followers[rand.Intn(len(followers))]

			r1, _ := json.Marshal(random_follower["node"])
			var target_user map[string]string
			json.Unmarshal(r1, &target_user)

			resp_1, err1 := req.Post("https://www.instagram.com/web/friendships/"+target_user["id"]+"/follow/", header)
			if err1 == nil {
				status := resp_1.Response().Status
				if status == "200 OK" {
					tools.Log("Followed: " + color.HEX("FFAA00").Sprint(target_user["username"]) + " Id: " + color.HEX("FFAA00").Sprint(target_user["id"]))

					profiles = append(profiles, profiles_struct{Id: target_user["id"], Time: strconv.FormatInt(time.Now().Unix(), 10)})
				} else {
					tools.Log(color.FgLightRed.Render("Follow failed"))
				}
			} else {
				tools.Log(color.FgLightRed.Render("Follow failed"))
			}
			tools.Sleep(5)
		}

		//? SAVE DATA
		file_2, _ := json.Marshal(profiles)
		_ = ioutil.WriteFile("data.json", file_2, os.ModePerm)

		for i1 := 0; i1 < 6; i1++ {
			followers, _ := tools.GetFollowers(user["id"], sessionid, csrftoken)
			if len(followers) < 1 {
				tools.GetUser(sessionid, csrftoken)
			}
			for _, d1 := range followers {
				//? GET ONE FOLLOWER
				ma1, _ := json.Marshal(d1["node"])
				var followers_user map[string]string
				json.Unmarshal(ma1, &followers_user)

				for i2, d2 := range profiles {
					//? GET ONE FOLLOWING
					ma2, _ := json.Marshal(d2)
					var profiles_user map[string]string
					json.Unmarshal(ma2, &profiles_user)

					//? CHECK ID
					if followers_user["id"] == profiles_user["id"] {
						//! UNFOLLOW
						resp_2, err2 := req.Post("https://www.instagram.com/web/friendships/"+profiles_user["id"]+"/unfollow/", header)
						if err2 == nil {
							status := resp_2.Response().Status
							if status == "200 OK" {
								//! REMOVE USER FROM SLICE
								profiles[i2] = profiles[len(profiles)-1]
								profiles = profiles[:len(profiles)-1]
								tools.Log(color.FgLightGreen.Render("New Follower: ") + color.HEX("FFAA00").Sprint(followers_user["username"]) + color.FgLightGreen.Render(" Id: ") + color.HEX("FFAA00").Sprint(followers_user["id"]))
								_, followers_count := tools.GetFollowers(user["id"], sessionid, csrftoken)
								_, following_count := tools.GetFollowing(user["id"], sessionid, csrftoken)
								tools.Log("Followers: " + color.HEX("FFAA00").Sprint(followers_count))
								tools.Log("Following: " + color.HEX("FFAA00").Sprint(following_count))
								break
							}
						}
					}
					unix_time, _ := strconv.ParseInt(profiles_user["time"], 10, 64)
					diff := time.Now().UTC().Sub(time.Unix(unix_time, 0))

					//? CHECK TIME
					if int64(diff.Hours()) >= 24 {
						//! UNFOLLOW
						resp_3, err3 := req.Post("https://www.instagram.com/web/friendships/"+profiles_user["id"]+"/unfollow/", header)
						if err3 == nil {
							status := resp_3.Response().Status
							if status == "200 OK" {
								//! REMOVE USER FROM SLICE
								profiles[i2] = profiles[len(profiles)-1]
								profiles = profiles[:len(profiles)-1]
								tools.Log("Unfollowed after 24h Id: " + color.HEX("FFAA00").Sprint(profiles_user["id"]))
								break
							}
						}
					}
				}
			}
			//? SAVE DATA
			file_3, _ := json.Marshal(profiles)
			_ = ioutil.WriteFile("data.json", file_3, os.ModePerm)
			tools.Sleep(300)
		}
	}
}
