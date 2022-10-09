package godogs

import (
	"github.com/cucumber/godog"
)

func iAmLoggedAsAdmin() error {
	return godog.ErrPending
}

func iCreateNewGenre() error {
	return godog.ErrPending
}

func iShouldSee(arg1 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am logged as Admin$`, iAmLoggedAsAdmin)
	ctx.Step(`^I create new genre$`, iCreateNewGenre)
	ctx.Step(`^I should see "([^"]*)"$`, iShouldSee)
}
