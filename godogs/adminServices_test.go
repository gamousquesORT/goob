package godogs

import (
	"context"

	"goob/app/pkg/domain"
	"testing"
	"strconv"

	"github.com/cucumber/godog"
)

var streamApp = *domain.NewStreamApp()

type StreamKey string

var sk StreamKey = "app"

type GenKey string

var genreName GenKey = "genName"

func iAmLoggedAsAdmin(ctx context.Context) (context.Context, error) {
	streamApp.CreateAdmin()
	return context.WithValue(ctx, sk, &streamApp), nil
}

func iCreateNewGenre(ctx context.Context, gName, gDescription string) (context.Context, error) {
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

func iAddTheFollowingGenres(ctx context.Context, genreTable *godog.Table) error {
	sa := ctx.Value(sk).(*domain.StreamApp)
	for _, r := range genreTable.Rows {
		var g domain.GenresData
		g.Name = r.Cells[0].Value
		g.Description = r.Cells[1].Value

		sa.CreateGenre(g.Name, g.Description)

	}
	return nil
}

func iShouldHaveGenresInTheApp(ctx context.Context, numberOfgenres int) error {
	sa := ctx.Value(sk).(*domain.StreamApp)
	if sa.GetNumberOfGenres() == numberOfgenres {
		return nil
	}

	return godog.ErrUndefined
}

func iAddTheFollowingMovies(ctx context.Context, movies *godog.Table) error {

	sa := ctx.Value(sk).(*domain.StreamApp)
	for _, r := range movies.Rows {
		var m domain.FilmData
		m.Name = r.Cells[0].Value
		m.Description = r.Cells[1].Value
		rat, _ := strconv.Atoi(r.Cells[2].Value)
		if rat == 0 {
			m.Rate = domain.G
		} else {
			m.Rate = domain.PG
		}
		boolValue, _ := strconv.ParseBool(r.Cells[3].Value)
		m.Sponsored = boolValue
		gd, err := sa.GetGenres(r.Cells[4].Value)
		if err != nil {
			return godog.ErrPending
		}
		
		err = sa.CreateFilm(m, gd)

		if err != nil {
			return godog.ErrPending
		}
	}
	return nil

}

func iShouldHaveMovieInTheApp(ctx context.Context, numberOfMovies int) error {
	sa := ctx.Value(sk).(*domain.StreamApp)
	if sa.GetNumberOfMovies() == numberOfMovies {
		return nil
	}
	return godog.ErrPending
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
	ctx.Step(`^I add the following Genres$`, iAddTheFollowingGenres)
	ctx.Step(`^I should have (\d+) Genres in the app$`, iShouldHaveGenresInTheApp)
	ctx.Step(`^I add the following Movies$`, iAddTheFollowingMovies)
	ctx.Step(`^I added the followin genres$`, iAddTheFollowingGenres)
	ctx.Step(`^I should have  (\d+) movie in the app$`, iShouldHaveMovieInTheApp)

}
