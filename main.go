/*
	nikocount
	- This is a custom POC script that shows how you can monitor a players stattrak kills over time.
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"nikocount/responses"
	"strconv"
	"strings"
	"time"
)

func getPlayerKills() error{

	var steamResponse = &responses.SteamResponse{}
	var client = &http.Client{Timeout: 10 * time.Second}

	r, err := client.Get("https://steamcommunity.com/inventory/76561198041683378/730/2?l=english&count=2000&start_assetid=17072530641")

	if err != nil {
		return err
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(steamResponse)


	for _, t := range steamResponse.Descriptions {
		for _, s := range t.Descriptions {

			if strings.Contains(s.Value, "StatTrak™ Confirmed Kills:"){

				weapon := strings.Replace(t.Name, "StatTrak™ ", "", -1)
				weapon = strings.TrimSpace(strings.Split( weapon, "|")[0])
				confirmedKills, err := strconv.Atoi(strings.Replace(s.Value, "StatTrak™ Confirmed Kills: ", "", -1))

				if err != nil {
					return err
				}

				fmt.Println("Weapon:", weapon, confirmedKills)
			}
		}
	}

	return nil
}


func main() {
	getPlayerKills()
}
