package godogs

import (
	"github.com/cucumber/godog"
	"testing"
	"goob/app/pkg/domain"
	"context"


)

var streamApp = *domain.NewStreamApp()
type StreamKey string 
var  sk StreamKey = "app"

func iAmLoggedAsAdmin(ctx context.Context) (context.Context, error) {
	streamApp.CreateAdmin()
	return context.WithValue(ctx, sk, streamApp), nil
}


func iCreateNewGenre(ctx context.Context) (context.Context, error) {
	sa := ctx.Value(sk).(domain.StreamApp) 
	sa.CreateGenre("Terror", "Genero para tener miedo")
	return context.WithValue(ctx, sk, streamApp), nil
}

func iShouldSee(arg1 string) error {
	return nil
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
	  ScenarioInitializer: InitializeScenario,
	  Options: &godog.Options{
		Format:   "pretty",
		Paths:    []string{"features"},
		TestingT: t, // Testing instance that will run subtests.
	  },
	}
  
	if suite.Run() != 0 {
	  t.Fatal("non-zero status returned, failed to run feature tests")
	}
  }
  
  
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I am logged as Admin$`, iAmLoggedAsAdmin)
	ctx.Step(`^I create new genre$`, iCreateNewGenre)
	ctx.Step(`^I should see "([^"]*)"$`, iShouldSee)
}
