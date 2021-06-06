package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Team struct {
	ID		string
	Name 	string
	Players []Player
	Score 	int
	Assists int
	scored3 int
	scored2 int
	scored3attempts int
	scored2attempts int
}

func (t *Team) Run(pctx context.Context) int {

	ctx, cancelFunc := context.WithTimeout(pctx, time.Second * time.Duration(rand.Int31n(24 / speedUp)))

	playerId := Rnd(0,4)
	attacker := t.Players[playerId]

	//Time for attack attempt, if time more than 24 seconds => team losts the ball
	tfaa := Rnd(2,25)

	for {
		select {
		case <-ctx.Done():
			//Time is over, kill the 24 second context
			cancelFunc()
			return 0
		default:
			tfaa--
			if tfaa > 0 {
				//Player is running with the ball
				time.Sleep(time.Second)
			} else {
				//Time for attach
				score := 0
				//Try to find open player with better Accuracy for assist, 50% chance for this decision
				if Rnd(1,10) > 5 {
					//Here is not 100% sure that assist will be applied, in case current player is the most accurate => no ball pass
					//OR more accurate player is blocked => no ball passed
					//In this case newPlayer is the current one and assist should not be counted
					newPlayer := t.assistTry(attacker)

					//Some chance for shooting 3 score
					shot3 := Rnd(0,4) == 2
					score = t.Attack(&newPlayer, shot3)

					if score > 0 {
						//Increase score, if game did not finished
						//In case newPlayer is not the same as attacking one => add assist counter
						t.scoreUp(ctx, score, newPlayer.ID != attacker.ID)

						return score
					}
				} else {
					//Direct shot without assist
					//Some chance for shooting 3 score
					shot3 := Rnd(0,4) == 2
					score = t.Attack(&attacker, shot3)
					if score > 0 {
						t.scoreUp(ctx, score, false)

						return score
					}
				}

				//If player missed, other teammate can try to catch and shot
				chanceTeammateCatchTheBall := Rnd(0, 3) == 1
				if score == 0 && chanceTeammateCatchTheBall {
					//Ideally to change player, but it won't be done
					score = t.Attack(&attacker, Rnd(0,4) == 2)
				}

				//Increase score, if game did not finished
				if score > 0 && t.scoreUp(ctx, score, false) {
					return score
				}

				//In case game finished return 0 or all tries have been failed
				return 0
			}
		}
	}

}

/** Attack attempt
We need to know is it 3 points or 2 points, as 3 points is more difficult to shot
In case ball won't reach the target we scored 0, but someone probably will pick up the ball
but this is out of this func responsibility
*/
func (t *Team) Attack(attacker *Player, score3 bool) int {
	score := 2
	if score3 {
		score = 3
		t.scored3attempts++
	} else {
		t.scored2attempts++
	}

	if attacker.Accuracy >= ComputeComplexity(score3) {
		fmt.Println(fmt.Sprintf("%s scored %d!", t.Name, score))
		return score
	}

	return 0
}

func (t *Team) assistTry(currentPlayer Player) Player {
	//Getting list of players, check their acc and emulate situation that more accurate player could be blocked and
	//assist is not possible, if possible give pass
	for _, player := range t.Players {
		chanceTeammateIsOpened := Rnd(0,3) > 1
		if player.Accuracy > currentPlayer.Accuracy && chanceTeammateIsOpened {
			//Success pass
			return player
		}
	}
	//No pass, all the more accurate players are weaker or blocked
	return currentPlayer
}

func (t *Team) scoreUp(ctx context.Context, score int, assist bool) bool {

	scored := false

	select {
		case <-ctx.Done():
			//Sorry but time is over
			break
		default:

			scored = true

			t.Score += score
			if assist {
				t.Assists++
			}

			if score == 2 {
				fmt.Println("Remembered score 2")
				t.scored2++
			} else {
				fmt.Println("Remembered score 3")
				t.scored3++
			}
	}

	return scored
}
