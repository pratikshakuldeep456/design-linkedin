package linkedin

import (
	"fmt"
)

type User struct {
	ID              int
	Name            string
	Email           string
	Password        string
	IsLoggedIn      bool
	Profile         *UserProfessionalInfo
	Connections     map[int]*User
	PendingRequests map[int]*User
}

// type Connections struct {
// }

func CreateUser(id int, name, email, password string) *User {
	return &User{
		ID:              id,
		Name:            name,
		Email:           email,
		Password:        password,
		IsLoggedIn:      true,
		Profile:         CreateProfile(),
		Connections:     make(map[int]*User),
		PendingRequests: make(map[int]*User),
	}
}

func (u *User) Login() error {
	if u.IsLoggedIn {
		return fmt.Errorf("user is already logged in")
	}
	u.IsLoggedIn = true

	return nil
}

func (u *User) LogOut() error {
	if !u.IsLoggedIn {
		return fmt.Errorf("user is not logged in")
	}
	u.IsLoggedIn = false

	return nil
}

func (u *User) UpdateInfo(Headline, summary string, experience *Experience, education *Education, skills *Skills) error {
	u.Profile.Headline = Headline
	u.Profile.Summary = summary

	// Append only if non-nil
	if experience != nil {
		u.Profile.Experience = append(u.Profile.Experience, experience)
	}
	if education != nil {
		u.Profile.Education = append(u.Profile.Education, education)
	}
	if skills != nil {
		u.Profile.Skills = append(u.Profile.Skills, skills)
	}

	return nil
}

func (u *User) UserExistInConnection(id int) error {
	_, k := u.Connections[id]
	if k {
		return fmt.Errorf("selected user is a connection")
	}
	return nil
}

func (u *User) UserExistInPendingRequest(id int) (*User, error) {
	data, exists := u.PendingRequests[id]

	if !exists {
		return nil, fmt.Errorf("selected user havent sent connection request")
	}

	return data, nil
}
func (u *User) AcceptRequest(id int) error {

	err := u.UserExistInConnection(id)
	if err != nil {
		return err
	}

	data, errr := u.UserExistInPendingRequest(id)
	if errr != nil {
		return errr
	}
	u.Connections[id] = data
	return nil
}

func (u *User) DeclineRequest(id int) error {
	err := u.UserExistInConnection(id)
	if err != nil {
		return err
	}

	_, errr := u.UserExistInPendingRequest(id)
	if errr != nil {
		return errr
	}

	delete(u.PendingRequests, id)
	return nil
}

func (u *User) SendRequest(to *User) error {
	if u.ID == to.ID {
		return fmt.Errorf("cant send request to self")
	}
	err := u.UserExistInConnection(to.ID)
	if err != nil {
		return fmt.Errorf("user already present in connection")
	}

	to.PendingRequests[u.ID] = u

	return nil
}

func (u *User) FetchconnectionList() []*User {
	var users []*User

	for _, user := range u.Connections {
		fmt.Println("user is", user)
		users = append(users, user)

	}

	return users

}
