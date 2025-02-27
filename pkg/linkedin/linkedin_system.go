package linkedin

import (
	"fmt"
	"sync"
)

type LinkedinSystem struct {
	Users map[int]*User
	Jobs  map[int]*Job
}

var (
	instance *LinkedinSystem
	once     sync.Once
)

func GetInstance() *LinkedinSystem {
	once.Do(func() {
		instance = &LinkedinSystem{
			Users: map[int]*User{},
			Jobs:  map[int]*Job{},
		}
	})

	return instance
}

func (ls *LinkedinSystem) RegisterUser(User User) {
	fmt.Println("user is registeered", User.Name)
	ls.Users[User.ID] = &User
}

func (ls *LinkedinSystem) Login(id int) error {

	err := ls.Users[id].Login()
	if err != nil {
		return err
	}

	fmt.Println("user is logging in")
	return nil
}

func (ls *LinkedinSystem) LogOut(id int) error {

	err := ls.Users[id].LogOut()
	if err != nil {
		return err
	}

	fmt.Println("user is loggout in")
	return nil
}

func (ls *LinkedinSystem) UpdateInfo(id int, Headline string, summary string, experience *Experience, education *Education, skills *Skills) error {
	ls.Users[id].UpdateInfo(Headline, summary, experience, education, skills)
	return nil
}
