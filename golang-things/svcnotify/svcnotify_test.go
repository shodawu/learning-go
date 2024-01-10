package svcnotify

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

type Sample struct {
	Svc SvcNotify
}

func (s *Sample) configContentIs(arg1 *messages.PickleStepArgument_PickleDocString) error {
	var err error

	var expected Config
	err = json.Unmarshal([]byte(arg1.Content), &expected)

	if s.Svc.Config != expected {
		return fmt.Errorf("expected return message: %v, but got: %v", arg1, s.Svc.Config)
	}
	return err
}

func (s *Sample) configPathIs(arg1 string) error {
	s.Svc.LoadConfig(arg1)
	return nil
}

func (s *Sample) serviceTriggerBySchduler() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	s := Sample{}
	ctx.Step(`^config content is$`, s.configContentIs)
	ctx.Step(`^config path is "([^"]*)"$`, s.configPathIs)
	ctx.Step(`^service trigger by schduler$`, s.serviceTriggerBySchduler)
}

func TestLoadConfig(t *testing.T) {
	status := godog.TestSuite{
		Name:                "Notify System",
		ScenarioInitializer: InitializeScenario,
	}.Run()

	os.Exit(status)

}
