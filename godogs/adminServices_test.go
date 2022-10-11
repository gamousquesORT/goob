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
type GenKey string
var genreName GenKey = "genName"

func iAmLoggedAsAdmin(ctx context.Context) (context.Context, error) {
	streamApp.CreateAdmin()
	return context.WithValue(ctx, sk, &streamApp), nil
}


func iCreateNewGenre(ctx context.Context, gName, gDescription string) (context.Context, error)  {
	sa := ctx.Value(sk).(*domain.StreamApp) 
	sa.CreateGenre(gName, gDescription)
	ctx = context.WithValue(ctx, genreName, gName)
	return context.WithValue(ctx, sk, &streamApp), nil
}

func iShouldBeAbleToRetrieveItGetting(ctx context.Context, gName, gDescription string) error {
	sa := ctx.Value(sk).(*domain.StreamApp)
	gd, err := sa.GetGenres(ctx.Value(genreName).(string))
	if err == nil && gd.Name == gName && gd.Description == gDescription {
		return nil
	}
	return err
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
	ctx.Step(`^I create new genre "([^"]*)", "([^"]*)"$`, iCreateNewGenre)
	ctx.Step(`^I should be able to retrieve it getting "([^"]*)", "([^"]*)"$`, iShouldBeAbleToRetrieveItGetting)
}


