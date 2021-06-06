package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/satori/go.uuid"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

var speedUp int32
var db *sql.DB

func main() {

	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/nba")
	db = conn

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//1 second read life is 5 virtual seconds seconds
	speedUp = 5
	//For randomizer
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", serveFiles)
	http.HandleFunc("/start", startGames)
	http.HandleFunc("/reset", resetData)
	http.HandleFunc("/data", getData)
	err = http.ListenAndServe(":9000", nil)

	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
	}
}

func getData(rw http.ResponseWriter, req *http.Request) {
	query := "SELECT " +
		"`plays`.`id`, " +
		"`t1`.`name` as `t1_name`, " +
		"`plays`.`team1_score`, " +
		"`plays`.`team1_ast`, " +
		"`t2`.`name` as `t2_name`, " +
		"`plays`.`team2_score`,`plays`.`team2_ast`, " +

		"`plays`.`team1_scored2`, " +
		"`plays`.`team1_scored2_att`, " +
		"`plays`.`team1_scored3`, " +
		"`plays`.`team1_scored3_att`, " +

		"`plays`.`team2_scored2`, " +
		"`plays`.`team2_scored2_att`, " +
		"`plays`.`team2_scored3`, " +
		"`plays`.`team2_scored3_att` " +

		"FROM `plays`" +
		"LEFT JOIN `teams` as `t1` ON `t1`.`id` = `plays`.`team1_id` " +
		"LEFT JOIN `teams` as `t2` ON `t2`.`id` = `plays`.`team2_id`"

	rows, err := db.Query(query)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	res := make([]Dashboard, 0, 15)
	for rows.Next() {
		d := Dashboard{}

		rows.Scan(
			&d.ID,
			&d.TeamA,
			&d.ScoreA,
			&d.ASTA,
			&d.TeamB,
			&d.ScoreB,
			&d.ASTB,
			&d.TeamAScored2,
			&d.TeamAScored2Att,
			&d.TeamAScored3,
			&d.TeamAScored3Att,
			&d.TeamBScored2,
			&d.TeamBScored2Att,
			&d.TeamBScored3,
			&d.TeamBScored3Att,
			)

		res = append(res, d)
	}

	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(res)
}

func resetData(rw http.ResponseWriter, req *http.Request) {
	tr1, _ := db.Query("TRUNCATE `teams`")
	tr1.Close()
	tr2, _ := db.Query("TRUNCATE `players`")
	tr2.Close()
	tr3, _ := db.Query("TRUNCATE `plays`")
	tr3.Close()
}

func startGames(rw http.ResponseWriter, req *http.Request) {
	resetData(rw, req)
	wg := sync.WaitGroup{}
	wg.Add(1)
	tms := GenerateTeams()

	rand.Seed(time.Now().UnixNano())
	//Shuffle teams
	rand.Shuffle(len(tms), func(i, j int) { tms[i], tms[j] = tms[j], tms[i] })

	plays := GeneratePlays(tms)

	for i := 0; i < len(plays); i++ {
		go plays[i].Start(&wg)
	}

	wg.Done()

	//Here is a bug, when we call this function one more time without server restart, it produces one more routine
	//Need to add dependency on context
	go func () {
		for {
			syncGameScores(plays)
			time.Sleep(time.Second)
		}
	}()

}

func syncGameScores(plays []Play) {

	amountOfPlays := len(plays)

	querySlice := make([]string, 0, amountOfPlays)
	querySlice2 := make([]string, 0, amountOfPlays)

	queryTmpl := "DELETE FROM plays WHERE `id` IN (%s)"
	queryTmpl2 := "INSERT INTO `plays` (`id`, `team1_id`, `team2_id`, `team1_score`, `team2_score`, `team1_ast`, `team2_ast`, `team1_scored2`, `team1_scored2_att`, `team1_scored3`, `team1_scored3_att`, `team2_scored2`, `team2_scored2_att`, `team2_scored3`, `team2_scored3_att`) VALUES %s"

	for _, play := range plays {
		querySlice = append(querySlice, fmt.Sprintf("'%s'", play.ID))
		querySlice2 = append(querySlice2,
			fmt.Sprintf("('%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d')",
				play.ID,
				play.TeamA.ID,
				play.TeamB.ID,
				play.TeamA.Score,
				play.TeamB.Score,
				play.TeamA.Assists,
				play.TeamB.Assists,
				play.TeamA.scored2,
				play.TeamA.scored2attempts,
				play.TeamA.scored3,
				play.TeamA.scored3attempts,
				play.TeamB.scored2,
				play.TeamB.scored2attempts,
				play.TeamB.scored3,
				play.TeamB.scored3attempts,

			))
	}

	del, _ := db.Query(fmt.Sprintf(queryTmpl, strings.Join(querySlice, ",")))

	ins, err := db.Query(fmt.Sprintf(queryTmpl2, strings.Join(querySlice2, ",")))
	if err != nil {
		panic(err)
	}

	defer del.Close()
	defer ins.Close()
}

func serveFiles(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html"
	}
	http.ServeFile(w, r, p)
}