package main

import (
	"fmt"
	"strings"
)

type Avenger struct {
	Name       string
	PersonName string
	Assigned   int
	Completed  int
	Abilities  string
}

type Mission struct {
	Team           []string
	MissionName    string
	MissionStatus  string
	MissionDetails string
}

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
func checkMissions(missions []Mission) {
	for _, v := range missions {
		fmt.Println(v)
	}
}

func printAvengerDetails(avengers []Avenger, name string) {
	for _, v := range avengers {
		if name == v.Name {
			fmt.Println(v)
			return
		}
	}
	fmt.Println("avenger does not exist")
	return
}

func printMissionDetails(missions []Mission, name string) {
	for _, v := range missions {
		if name == v.MissionName {
			fmt.Println(v)
			return
		}
	}
	fmt.Println("mission not found")
	return

}

func updateMissionStatus(missions []Mission, name, status string) []Mission {
	for _, v := range missions {
		if v.MissionDetails == name {
			v.MissionStatus = status
		}
	}
	return missions
}

func getMissionDetails(mission []Mission, name string) (Mission, int) {
	for _, mi := range mission {
		if mi.MissionName == name {
			return mi, 1
		}
	}
	return mission[0], 0
}
func getStatus(avengers []Avenger, name string) string {
	var a int
	for i, _ := range avengers {
		if avengers[i].Name == name {
			a = avengers[i].Assigned
		}
		if a == 0 {
			return "Available"
		} else {
			return "On Mission"
		}

	}
	return ""
}
func getMission(mission []Mission, name string) []string {
	task := []string{}
	for i, _ := range mission {
		for j, _ := range mission[i].Team {
			if mission[i].Team[j] == name {
				task = append(task, mission[i].MissionName)
			}
		}
	}
	return task
}
func listAvengers(avengers []Avenger, mission []Mission) {
	fmt.Println("Avenger Name | Status | Assigned Mission")
	var status1 string
	for i, _ := range avengers {
		status1 = getStatus(avengers, avengers[i].Name)
		task := getMission(mission, avengers[i].Name)
		if len(task) == 0 {
			fmt.Println(avengers[i].Name, status1)

		} else {
			fmt.Println(avengers[i].Name, status1, task)
		}
	}
}

func checkCount(avenger []Avenger, name string) (int, int) {
	for i, _ := range avenger {
		if avenger[i].Name == name {
			return avenger[i].Assigned, 1
		}
	}
	return 999, 999
}

func updateCount(avengers []Avenger, name string) []Avenger {
	for i, _ := range avengers {
		if avengers[i].Name == name {
			avengers[i].Assigned += 1
		}
	}
	return avengers

}
func assignMission(mission []Mission, team []string, name, details string) []Mission {
	newMission := &Mission{
		Team:           team,
		MissionName:    name,
		MissionStatus:  "Assigned",
		MissionDetails: details,
	}
	mission = append(mission, *newMission)
	return mission
}
func main() {
	var avengers = []Avenger{
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
	var missions = []Mission{}
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
			if len(missions) == 0 {
				fmt.Println("No missions are assigned as of now")
			} else {
				checkMissions(missions)
			}

		case 2:
			fmt.Println("Assign mission to avengers")
			var name string
			flag1 := true
			var details string
			fmt.Println("Enter Avengers")
			fmt.Scan(&name)
			team := strings.Split(name, ",")
			for i, _ := range team {
				cnt, _ := checkCount(avengers, team[i])
				if cnt >= 2 {
					fmt.Println("alredy workingon two mission ", team[i])
					flag1 = false
				} else {
					avengers = updateCount(avengers, team[i])
				}

			}
			if flag1 == true {
				fmt.Println("enter mission name")
				fmt.Scan(&name)
				fmt.Println("Enter the mission details")
				fmt.Scan(&details)
				missions = assignMission(missions, team, name, details)
				fmt.Println("Mission has been assigened")
				fmt.Println("notificatin sent")
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
			printMissionDetails(missions, name)

		case 4:
			fmt.Println("check avengers details")
			var name string
			fmt.Println("enter the avenger name")
			_, err := fmt.Scan(&name)
			if err != nil {
				fmt.Println("inter the valid name")
			} else {
				printAvengerDetails(avengers, name)
			}

		case 5:
			fmt.Println("update mission status")
			var name string
			var status string
			fmt.Println("Enter the mission name")
			fmt.Scan(&name)
			fmt.Println("Enter the missin status")
			fmt.Scan(&status)
			mission := updateMissionStatus(missions, name, status)
			team, _ := getMissionDetails(mission, name)
			fmt.Println("Email and sms sent to team ", team.Team)
		case 6:
			fmt.Println("list avengers")
			listAvengers(avengers, missions)
		}
	}

}
