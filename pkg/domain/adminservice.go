package domain

type StreamApp struct {
	adminUser *UserData
}


var StreamAppData StreamApp

func (sta *StreamApp) CreateAdmin() (UserData, error) {
	var err error
	 sta.adminUser, err = NewUser("Admin6789101", "admin@example.com", "12345678901")
	 sta.adminUser.SetAdmin(true)
	 if err != nil {
		return *sta.adminUser, err
	 }
	return *sta.adminUser, nil
}



