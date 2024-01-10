package game

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
)

func (gs *Sample) gameATargetIs(arg1 int) error {
	gs.Game.Answer = arg1
	return nil
}

func (gs *Sample) gameReturnMessage(arg1 string) error {
	if arg1 != gs.Game.GuessMsg {
		return fmt.Errorf("expected return message: %v, but got: %v", arg1, gs.Game.GuessMsg)
	}
	return nil
}

func (gs *Sample) hasMoreTimeToGuess(arg1 int) error {
	gs.Game.Opptunity = arg1
	return nil
}

func (gs *Sample) playerGuess(arg1 int) error {
	gs.Game.GuessMsg = gs.Game.Guess(arg1)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	g := Sample{}
	ctx.Step(`^game a target is (\d+)$`, g.gameATargetIs)
	ctx.Step(`^game return message "([^"]*)"$`, g.gameReturnMessage)
	ctx.Step(`^has (\d+) more time to guess$`, g.hasMoreTimeToGuess)
	ctx.Step(`^player guess (\d+)$`, g.playerGuess)
}

func TestGuessBDD(t *testing.T) {
	status := godog.TestSuite{
		Name:                "猜數字",
		ScenarioInitializer: InitializeScenario,
	}.Run()

	os.Exit(status)

}
