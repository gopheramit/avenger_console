package main

import "fmt"

func (a *Avenger) IsAvialable() bool {
	if len(a.Missions) < 2 {
		return true
	}
	return false
}

func (m *Manger) GetMissions() {
	for _, v := range m.missions {
		fmt.Println(v)
	}
}

func (m *Manger) findAvengerByName(name string) *Avenger {
	for i := range m.avengers {
		if m.avengers[i].Name == name {
			return &m.avengers[i]
		}
	}
	return nil
}

func (m *Manger) AddMission(mission *Mission, avenger *Avenger) {
	if avenger.IsAvialable() {
		m.missions = append(m.missions, *mission)
		avenger.Missions = append(avenger.Missions, *mission)
	}

}

func (m *Manger) FindMissionByName(name string) *Mission {
	for i := range m.missions {
		if m.missions[i].Name == name {
			return &m.missions[i]
		}
	}
	return nil
}

func (m *Manger) UpdateMissionStatus(name, status string) {
	mission := m.FindMissionByName(name)
	if mission != nil {
		mission.Status = status
	}

}
func (m *Manger) ListAvengers() {
	for i := range m.avengers {
		fmt.Println(m.avengers[i])
	}
}

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
