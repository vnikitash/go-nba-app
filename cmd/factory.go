package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GenerateTeams() []Team {
	//30 teams
	tl := 30
	teams := make([]Team, 0, tl)
	querySlice := make([]string, 0, tl)

	queryTmpl := "INSERT INTO `teams` (`id`, `name`) VALUES %s"

	for i := 1; i < tl+1; i++ {
		nt := Team{getUUIDString(), fmt.Sprintf("Team %d", i), make([]Player, 0,5), 0, 0, 0, 0,0,0}
		nt.Players = GenerateTeamPlayers(&nt)
		querySlice = append(querySlice, fmt.Sprintf("('%s', '%s')", nt.ID, nt.Name))
		teams = append(teams, nt)
	}

	ins, _ := db.Query(fmt.Sprintf(queryTmpl, strings.Join(querySlice, ",")))
	ins.Close()

	return teams
}

func GenerateTeamPlayers(t *Team) []Player {
	players := make([]Player, 0, 5)

	querySlice := make([]string, 0, 5)

	queryTmpl := "INSERT INTO `players` (`id`, `name`, `accuracy`, `team_id`) VALUES %s"

	for i := 1; i < 6; i++ {
		np := Player{getUUIDString(), fmt.Sprintf("Player %d", i), Rnd(60,92), t.ID}
		players = append(players, np)
		querySlice = append(querySlice, fmt.Sprintf("('%s', '%s', '%d', '%s')", np.ID, np.Name, np.Accuracy, t.ID))
	}


	ins, _ := db.Query(fmt.Sprintf(queryTmpl, strings.Join(querySlice, ",")))
	ins.Close()

	return players
}

func GeneratePlays(tms []Team) []Play {
	amountOfPlays := int(len(tms) / 2)

	querySlice := make([]string, 0, amountOfPlays)

	queryTmpl := "INSERT INTO `plays` (`id`, `team1_id`, `team2_id`) VALUES %s"

	//fmt.Println("Teams :" + strconv.Itoa(amountOfPlays))
	plays := make([]Play, 0, amountOfPlays)
	for i := 0; i < len(tms); i += 2 {
		np := Play{getUUIDString(), tms[i], tms[i + 1], 0}
		querySlice = append(querySlice, fmt.Sprintf("('%s', '%s', '%s')", np.ID, np.TeamA.ID, np.TeamB.ID))
		plays = append(plays, np)
	}

	ins, _ := db.Query(fmt.Sprintf(queryTmpl, strings.Join(querySlice, ",")))
	ins.Close()

	return plays
}

func getUUIDString() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		panic("Error with uuid generator")
	}

	return uuid.String()
}