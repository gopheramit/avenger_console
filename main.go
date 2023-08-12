package main

import (
	"fmt"
	"strings"
)

func PrintConsole() {
	fmt.Println("------------------shield---------------")
	fmt.Println("Welcome to captain fury")
	fmt.Println("1.Check the missions")
	fmt.Println("2.Assign mission to avengers")
	fmt.Println("3.Check mission details")
	fmt.Println("4.Check avengers details")
	fmt.Println("5.Update mission status")
	fmt.Println("6.List avengers")
}

type Avenger struct {
	Name       string
	PersonName string
	Assigned   int
	Completed  int
	Abilities  string
	Missions   []Mission
}

type Mission struct {
	Name    string
	Status  string
	Details string
}

type Manger struct {
	avengers []Avenger
	missions []Mission
}

func main() {
	var avengersData = []Avenger{
		Avenger{
			Name:       "Iron Man",
			PersonName: "Tony Stark",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Genius,Rich,Tech armored suit",
		},
		Avenger{
			Name:       "Captain America",
			PersonName: "Steve Rogers",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Leadrship,ethics,super soldier",
		},
		Avenger{
			Name:       "Hulk",
			PersonName: "Bruce Banner",
			Assigned:   0,
			Completed:  0,
			Abilities:  "smart,strongest avenger",
		},
		Avenger{
			Name:       "Thor",
			PersonName: "Thor",
			Assigned:   0,
			Completed:  0,
			Abilities:  "God of Thunder",
		},
		Avenger{
			Name:       "Black Widow",
			PersonName: "Natasha Rommanof",
			Assigned:   0,
			Completed:  0,
			Abilities:  "smart,spy",
		},
		Avenger{
			Name:       "Hawkeye",
			PersonName: "Clint Barton",
			Assigned:   0,
			Completed:  0,
			Abilities:  "Guy who can aim very well and see evrything",
		},
	}
	manger := Manger{
		avengers: avengersData,
		missions: []Mission{},
	}
	//var missions = []Mission{}
	for {
		PrintConsole()
		var choice int
		fmt.Println("enter the choice")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("enter valid choice")
		}
		switch choice {
		case 1:
			fmt.Println("check the missions")
			if len(manger.missions) == 0 {
				fmt.Println("No missions are assigned as of now")
			} else {
				manger.GetMissions()

			}

		case 2:
			fmt.Println("Assign mission to avengers")
			var names string
			var missionDetails string
			var missionName string
			fmt.Println("Enter avengers names")
			fmt.Scan(&names)
			avengerNames := strings.Split(names, ",")
			fmt.Println("Enter missionDetails")
			fmt.Scan(&missionDetails)
			fmt.Println("Enter missionName")
			fmt.Scan(&missionName)
			mission := &Mission{
				Name:    missionName,
				Details: missionDetails,
			}
			for _, name := range avengerNames {
				//avenger here is pointer
				avenger := manger.findAvengerByName(name)
				if avenger != nil {
					manger.AddMission(mission, avenger)
				}
			}

		case 3:
			fmt.Println("check mission details")
			var name string
			fmt.Println("enter mission name")
			_, err := fmt.Scan(&name)
			if err != nil {
				fmt.Println("enter valid mission name")
				return
			}
			mission := manger.FindMissionByName(name)
			if mission != nil {
				fmt.Println(mission)
			} else {
				fmt.Println("Mission not found")
			}

		case 4:
			fmt.Println("check avengers details")
			var name string
			fmt.Println("enter the avenger name")
			_, err := fmt.Scan(&name)
			if err != nil {
				fmt.Println("inter the valid name")
			} else {
				avenger := manger.findAvengerByName(name)
				if avenger != nil {
					fmt.Println(avenger)
				} else {
					fmt.Println("Avenger deatils not found")
				}
			}

		case 5:
			fmt.Println("update mission status")
			var name string
			var status string
			fmt.Println("Enter the mission name")
			fmt.Scan(&name)
			fmt.Println("Enter the missin status")
			fmt.Scan(&status)
			manger.UpdateMissionStatus(name, status)
			// mission := manger.FindMissionByName(name)
			// if mission != nil {
			// 	mission.Status = status
			// }
			// mission := updateMissionStatus(missions, name, status)
			// team, _ := getMissionDetails(mission, name)
			// fmt.Println("Email and sms sent to team ", team.Team)
		case 6:
			fmt.Println("list avengers")
			manger.ListAvengers()
			// listAvengers(avengers, missions)
		}
	}

}
