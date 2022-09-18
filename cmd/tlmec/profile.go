package domain





type ProfileData struct {
	Alias string
	Pin int
}



func NewProfile(alias string, pin int) (*ProfileData, error) {
	val := new(ProfileData)
	val.Alias = alias
	val.Pin = pin

	return  val,nil
}
