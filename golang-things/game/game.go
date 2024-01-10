package game

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"number-game/interfaces"
)

// Game ..
type Game struct {
	Answer    int
	Opptunity int
	GuessMsg  string //for godog
}

// Guess ...
func (g *Game) Guess(num int) string {
	var s string = "遊戲已結束"
	if g.Opptunity <= 0 {
		return s
	}

	g.Opptunity--

	if g.Answer == num {
		s = "恭喜過關" //
	} else {
		if g.Opptunity <= 0 {
			s = "失敗" //
		} else {
			if g.Answer > num {
				s = "過小" //
			} else {
				s = "過大" //
			}
		}
	}

	return s
}

// Sample ...
type Sample struct {
	Desc    string
	Game    Game
	Guesses []Guess
	Command interfaces.ICommand
}

// Guess ...
type Guess struct {
	Guess     int
	WantMsg   string
	WatnTimes int
}

// Exec ...
func (gs *Sample) Exec() bool {

	log.Printf("game.Guess(), test for %v\n", gs.Desc)
	for iguess, guess := range gs.Guesses {
		got := gs.Game.Guess(guess.Guess)
		log.Printf("step %v\n", iguess)
		if got != guess.WantMsg || gs.Game.Opptunity != guess.WatnTimes {
			log.Fatalf("Got: %v, %v; but expect:  %v, %v",
				got, gs.Game.Opptunity, guess.WantMsg, guess.WatnTimes)
		}
	}

	return true
}

//猜數字遊戲， 可以設定1..100的數字，跟最多猜幾次。
// 猜對了變回傳「恭喜過關」並且結束遊戲
// 如果猜錯就回傳「過大」、「過小」來提示玩家
// 如果猜到最後一次都猜錯，就回傳「失敗」，並且結束遊戲。
// 已結束的遊戲，總是回傳「遊戲已結束」

// LoadSamples ...
func LoadSamples() (res []interfaces.ICommand) {
	var mGameCases []*Sample
	fByte, _ := ioutil.ReadFile("./game/game-guess.json")
	json.Unmarshal(fByte, &mGameCases)

	for _, gs := range mGameCases {
		res = append(res, gs)
	}
	return
}
