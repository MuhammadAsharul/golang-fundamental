package management

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

// methdod
func (user User) Display() string {
	return fmt.Sprintf("Name : %s %s, dengan email : %s", user.FirstName, user.LastName, user.Email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

// embedded struct
func (group Group) DisplayGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Printf("\nMember Count : %d \n", len(group.Users))
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
