package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Play struct {
	ID string
	TeamA Team
	TeamB Team
	BallHolder int
}

func (p *Play) GetPlayName() string {

	tmpl := "Play: %s vs %s"
	return fmt.Sprintf(tmpl, p.TeamA.Name, p.TeamB.Name)
}



func (p *Play) Start(wg *sync.WaitGroup) {
	//Waiting other Plays to be ready to be started
	wg.Wait()

	//Define context with timeout for 48 minutes
	ctx, gameFinished := context.WithTimeout(context.Background(), time.Second * time.Duration(rand.Int31n(60 * 48 / speedUp)))
	//define which team holds the ball
	attackingTeam := &p.TeamA
	if Rnd(0,1) == 0 {
		attackingTeam = &p.TeamB
	}

	for {
		select {
		case <-ctx.Done():
			gameFinished()
			fmt.Println("Game Finished")
			fmt.Println(fmt.Sprintf("%s => SCRORE: %d, AST: %d", p.TeamA.Name, p.TeamA.Score, p.TeamA.Assists))
			fmt.Println(fmt.Sprintf("%s => SCRORE: %d, AST: %d", p.TeamB.Name, p.TeamB.Score, p.TeamB.Assists))
			return
		default:
			points := attackingTeam.Run(ctx)

			if points == 0 {
				//fmt.Println(fmt.Sprintf("%s lost the ball...", attackingTeam.Name))
			}

			//After 24 seconds or succeeded goal change ball holder team
			if attackingTeam.Name == p.TeamB.Name {
				attackingTeam = &p.TeamA
			} else {
				attackingTeam = &p.TeamB
			}

			break
		}
	}
}