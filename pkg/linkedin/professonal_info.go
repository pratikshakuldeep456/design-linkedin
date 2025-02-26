package linkedin

type UserProfessionalInfo struct {
	ProfilePic string
	Headline   string
	Summary    string
	Experience []*Experience
	Education  []*Education
	Skills     []*Skills
}

func CreateProfile() *UserProfessionalInfo {
	return &UserProfessionalInfo{}
}
