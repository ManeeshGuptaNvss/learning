package pointers

import "fmt"

/*
Pointers are prone to cause errors so use them wisely.
When to use pointers
1. When we want to update the state
2. When we want to optimize the memory for large objects(structs) that are getting called A LOT. 
*/ 
type User struct {
	Name  string
	Age   string
	Email string
	File [] byte // ?? Large ??
}

func (u User) FetchEmail() string {
	return u.Email
}

func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

func PrintEmail() {
	user := User{
		Name:  "Maneesh",
		Email: "nvssmg@gmail.com",
	}
	fmt.Println(user.FetchEmail())
	user.UpdateEmail("nvsmaneesh8@gmail.com")
	fmt.Println(user.FetchEmail())
}
