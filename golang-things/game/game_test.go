package game

import (
	"testing"
)

// 單元測試重點
// 1. 程式碼覆蓋率
// 2. 邊界值

func TestGuess(t *testing.T) {

	tGame := Sample{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Guesses: []Guess{
			{
				Guess:   8,
				WantMsg: "過小",
			},
			{
				Guess:   99,
				WantMsg: "過大",
			},
			{
				Guess:   99,
				WantMsg: "失敗",
			},
		},
	}

	for _, guess := range tGame.Guesses {
		got := tGame.Game.Guess(guess.Guess)
		if got != guess.WantMsg {
			t.Errorf("Guess() = %v; want %v", got, guess.WantMsg)
		}
	}

}

// TestGuessCodeFirst 先有code 再有testing case
func TestGuessCodeFirst(t *testing.T) {

	mGameCases := map[string]Sample{}
	mGameCases["PassAt1st"] = Sample{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Guesses: []Guess{
			{
				Guess:   90,
				WantMsg: "恭喜過關",
			},
		},
	}

	mGameCases["PassAt4st"] = Sample{
		Game: Game{
			Answer:    90,
			Opptunity: 3,
		},
		Guesses: []Guess{
			{
				Guess:   91,
				WantMsg: "過大",
			},
			{
				Guess:   89,
				WantMsg: "過小",
			},
			{
				Guess:   1,
				WantMsg: "失敗",
			},
			{
				Guess:   90,
				WantMsg: "遊戲已結束",
			},
		},
	}

	for k, gameCase := range mGameCases {
		for _, guess := range gameCase.Guesses {
			got := gameCase.Game.Guess(guess.Guess)
			if got != guess.WantMsg {
				t.Errorf("Test for %v , Guess() = %v; want %v", k, got, guess.WantMsg)
			}
		}
	}

}
