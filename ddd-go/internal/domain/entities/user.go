package entities

type User struct {
    ID    string `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// NewUser creates a new User instance
func NewUser(id, name, email string) *User {
    return &User{
        ID:    id,
        Name:  name,
        Email: email,
    }
}

// UpdateEmail updates the email of the User
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}