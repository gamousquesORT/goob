package domain_test

import (
	domain "goob/domain/cmd/tlmec"
	"reflect"
	"testing"
)

func createValidFilm(t *testing.T) domain.FilmData {
	t.Helper()

	want := new(domain.FilmData)
	g := domain.GenresData{"Terror", "da miedo"}
	want.Name = "Matrix"
	want.Descr = "pelicula sci-fi buena parte 1"
	want.Rate = domain.G
	want.Sponsored = true
	want.Genres = make([]domain.FilmGenre, 1)
	fg := domain.FilmGenre{}
	fg.Gener = g
	fg.GnType = domain.MainGenre
	want.Genres[0] = fg
	return *want
}

func TestValidFilm(t *testing.T) {
	t.Run("Should return no error given a valid Film data", func(t *testing.T) {
		g := domain.GenresData{"Terror", "da miedo"}
		got, err := domain.NewFilmData("Matrix", g, domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)

		want := createValidFilm(t)

		if !reflect.DeepEqual(got, &want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)

	})

	t.Run("Should return an error given a Film without name", func(t *testing.T) {
		g := domain.GenresData{}
		_, err := domain.NewFilmData("", g, domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)

		assertError(t, err, domain.ErrFilmNameMissing)
	})

	t.Run("Should return an error given a secondary genres at film creation", func(t *testing.T) {
		g := domain.GenresData{"Terror", "da miedo"}
		_, err := domain.NewFilmData("Matrix", g, domain.SecondaryGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)

		assertError(t, err, domain.ErrExpectedPrimaryGenre)

	})

	t.Run("Should return two Genres given two valid Genres", func(t *testing.T) {
		g1 := domain.GenresData{"Terror", "da miedo"}
		f1, _ := domain.NewFilmData("Matrix", g1, domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)

		g2 := domain.GenresData{"Acción", "entretenido"}

		err1 := f1.AddGenre(g2, domain.SecondaryGenre)

		got := f1.GetGenres()

		want := make([]domain.FilmGenre, 2)
		want[0] = domain.FilmGenre{domain.GenresData{Name: "Terror", Description: "da miedo"}, domain.MainGenre}
		want[1] = domain.FilmGenre{domain.GenresData{Name: "Acción", Description: "entretenido"}, domain.SecondaryGenre}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err1)

	})

	t.Run("Should return an error given more than four genres", func(t *testing.T) {
		g := domain.GenresData{"Terror", "da miedo"}
		f1, _ := domain.NewFilmData("Matrix", g, domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)

		f1.AddGenre(g, domain.SecondaryGenre)
		f1.AddGenre(g, domain.SecondaryGenre)
		f1.AddGenre(g, domain.SecondaryGenre)

		err := f1.AddGenre(g, domain.SecondaryGenre)
		assertError(t, err, domain.ErrTooManyGenres)
	})

}
