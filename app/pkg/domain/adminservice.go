package domain



type StreamApp struct {
	adminUser *UserData
	genres []GenresData
}


func NewStreamApp() *StreamApp {
	return new(StreamApp)
}


func (sta *StreamApp) CreateAdmin() (UserData, error) {
	var err error
	 sta.adminUser, err = NewUser("Admin6789101", "admin@example.com", "12345678901")
	 sta.adminUser.SetAdmin(true)
	 if err != nil {
		return *sta.adminUser, err
	 }
	return *sta.adminUser, nil
}


func (sta *StreamApp) CreateGenre(name, desc string) (error) {
	gen, err := NewGenresData(name, desc)
	sta.genres = append(sta.genres, *gen)

	return err
}

func (sta StreamApp) GetGenres(name string) (GenresData, error) {
	for _,g := range sta.genres {
		if g.Name == name {
			return g, nil
		}
	}
	return GenresData{}, ErrInvalidGenresData
}


