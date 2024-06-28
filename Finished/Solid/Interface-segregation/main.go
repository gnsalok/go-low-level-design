package main

/*

ISP (Interface Segregation) states that clients should not be forced to depend on interfaces they do not use,
meaning that the interfaces should be designed to be as small and specific as possible.
This helps to keep the code flexible and avoids unnecessary coupling between classes.

*/

type UserManager interface {
	CreateUser(name, email string) (User, error)
	GetUser(id int) (User, error)
	UpdateUser(id int, updates map[string]interface{}) error
	DeleteUser(id int) error
	// Optional: AssignRole(userID int, role string) error (if needed)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// ... other user details (optional)
	Role string `json:"role,omitempty"` // Optional: If roles are managed
}

/*
// ----------- User manager interface encompasses all functionalities.
- Clients that only need user creation or retrieval might be forced to implement unused methods like `UpdateUser` and `DeleteUser`.
- A big interface can be broken into various interfaces  like below
- Note that, If any specific functionality is needed by any interface, that can be embedded
*/

//

// UserCreator defines an interface for creating users
type UserCreator interface {
	CreateUser(name, email string) (User, error)
}

// UserGetter defines an interface for retrieving user information
type UserGetter interface {
	GetUser(id int) (User, error)
}

// UserUpdater defines an interface for updating user details
type UserUpdater interface {
	UpdateUser(id int, updates map[string]interface{}) error
}

// UserDeleter defines an interface for deleting users
type UserDeleter interface {
	DeleteUser(id int) error
}

func main() {

}
