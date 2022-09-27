package domain_test

import ("testing"
"streamapp.com/domain"
"reflect"
)

func createValidFilm(t *testing.T) domain.FilmData {
	t.Helper()

	want := new(domain.FilmData)
	g := domain.GenresData{"Terror", "da miedo"}
	want.Name ="Matrix"
	want.Descr = "pelicula sci-fi buena parte 1"
	want.Rate = domain.G
	want.Sponsored = true
	want.Genres = make([]domain.FilmGenre, 4)
	fg := domain.FilmGenre{}
	fg.Gener = g
	fg.GnType = domain.MainGenre
	want.Genres[0] = fg
	return *want
}

func TestValidFilm(t *testing.T) {
	t.Run("Should return no error given a valid Film data", func(t *testing.T) {
		g := domain.GenresData{"Terror", "da miedo"}
		got, err := domain.NewFilmData("Matrix", g ,domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)
		
		want := createValidFilm(t)

		if !reflect.DeepEqual(got, &want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})
	

	t.Run("Should return an error given a Film without name", func(t *testing.T) {
		g := domain.GenresData{}
		_, err := domain.NewFilmData("", g,domain.MainGenre,  "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)
	
		assertError(t, err, domain.ErrFilmNameMissing)
	})


	t.Run("Should return an error given a secondary genres without a primary genre", func(t *testing.T) {
		g := domain.GenresData{"Terror", "da miedo"}
		got, err := domain.NewFilmData("Matrix", g ,domain.MainGenre, "path to picture storage", "pelicula sci-fi buena parte 1", domain.G, true)
		
		want := new(domain.FilmData)
		want.Name ="Matrix"
		want.Descr = "pelicula sci-fi buena parte 1"
		want.Rate = domain.G
		want.Sponsored = true
		want.Genres = make([]domain.FilmGenre, 4)
		fg := domain.FilmGenre{}
		fg.Gener = g
		fg.GnType = domain.MainGenre
		want.Genres[0] = fg

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v , want %v", got, want)
		}

		assertNoError(t, err)
	
	})

	t.Run("Shoud return given a -1 rating it should return -1", func(t *testing.T) {
		
	})
}