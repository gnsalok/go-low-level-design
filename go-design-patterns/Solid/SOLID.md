#system-design/lld #programming/go

See  [go-design pattern](https://github.com/gnsalok/design-patterns-go)  repo for Implementation 

In software Engineering, it’s important to have a set of design principles to guide you in creating high-quality, maintainable, and scalable code. One such set of principles is the SOLID.

[Medium Link](https://akhileshmj.medium.com/solid-principles-go-design-pattern-6af77d665b8e)
### Single Responsibility Principle (SRP)

This principle states that a struct should have only one reason to change, meaning that a struct should have only one responsibility. This helps to keep the code clean and maintainable, as changes to the struct only need to be made in one place.

Let’s say I have a struct `Employee` that keeps track of an employee's name, salary, and address:

```go
type Employee struct {  
Name string  
Salary float64  
Address string  
}

```

According to the SRP, each struct should have only one responsibility, so in this case, it would be better to split the responsibilities of the `Employee` struct into two separate structs: `EmployeeInfo` and `EmployeeAddress`.

```go
type EmployeeInfo struct {  
Name string  
Salary float64  
}  
  
type EmployeeAddress struct {  
Address string  
}
```

Now, we can have separate functions that handle the different responsibilities of each struct:

```go 
type EmployeeInfo struct {  
Name string  
Salary float64  
}  
  
type EmployeeAddress struct {  
Address string  
}

```

By following the SRP, the code has become more maintainable and easier to understand, as each structure now has a clear and specific responsibility. If we need to make changes to the salary calculation or address handling, we know exactly where to look, without having to wade through a lot of unrelated code.

### Open/Closed Principle (OCP)

This principle states that a struct should be open for extension but closed for modification, meaning that the behavior of a struct can be extended without changing its code. This helps to keep the code flexible and adaptable to changing requirements.

Let’s say We have a task to build a payment system that will be able to process credit card payments. It should also be flexible enough to accept different types of payment methods in the future.

  ```go

package main  
  
import "fmt"  
  
type PaymentMethod interface {  
  Pay()  
}  
  
type Payment struct{}  
  
func (p Payment) Process(pm PaymentMethod) {  
  pm.Pay()  
}  
  
type CreditCard struct {  
  amount float64  
}  
  
func (cc CreditCard) Pay() {  
  fmt.Printf("Paid %.2f using CreditCard", cc.amount)  
}  
  
func main() {  
  p := Payment{}  
  cc := CreditCard{12.23}  
  p.Process(cc)  
}
```

As per OCP, `Payment` struct is open for extension and closed for modification. Since we are using `PaymentMethod` the interface, we don’t have to edit `Payment` behavior when adding new payment methods. Adding something like PayPal would look like this:

```go

type PayPal struct {  
  amount float64  
}  
  
func (pp PayPal) Pay() {  
  fmt.Printf("Paid %.2f using PayPal", pp.amount)  
}  
  
// then in main()  
pp := PayPal{22.33}  
p.Process(pp)
```




### Liskov Substitution Principle (LSP)

Consider the struct:

```go
type Animal struct {  
  Name string  
}  
  
func (a Animal) MakeSound() {  
  fmt.Println("Animal sound")  
}

```
Now, let’s say we want to create a new struct `Bird` that represents a specific type of animal:

```go 
type Bird struct {  
  Animal  
}  
  
func (b Bird) MakeSound() {  
  fmt.Println("Chirp chirp")  
}
```

This principle states that objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program. This helps to ensure that the relationships between classes are well-defined and maintainable.

```go
type AnimalBehavior interface {  
  MakeSound()  
}  
  
// MakeSound represent a program that works with animals and is expected  
// to work with base class (Animal) or any subclass (Bird in this case)  
func MakeSound(ab AnimalBehavior) {  
  ab.MakeSound()  
}  
  
a := Animal{}  
b := Bird{}  
MakeSound(a)  
MakeSound(b)

```

This demonstrates inheritance in Go, as well as the Liskov Substitution Principle, as objects of a subtype `Bird` can be used wherever objects of the base type `Animal` are expected, without affecting the correctness of the program.


### Interface Segregation Principle (ISP)

ISP states that clients should not be forced to depend on interfaces they do not use, meaning that the interfaces should be designed to be as small and specific as possible. This helps to keep the code flexible and avoids unnecessary coupling between classes.

I'd be glad to provide an example of Interface Segregation Principle (ISP) in Golang, combining the best aspects of previous responses and addressing potential issues:

**Scenario: User Management**

Imagine a system for managing users, where you might need to perform various operations:

- Creating new users
- Retrieving user information
- Updating user details (e.g., name, email)
- Deleting users
- Assigning roles/permissions to users (optional)

**Without ISP:**

```go
package user

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
  Role  string `json:"role,omitempty"` // Optional: If roles are managed
}
```

This single `UserManager` interface encompasses all functionalities. However, this approach violates ISP:

- Clients that only need user creation or retrieval might be forced to implement unused methods like `UpdateUser` and `DeleteUser`.
- If role management is not required, including the `AssignRole` method adds unnecessary complexity.

**With ISP:**

```go
package user

// User represents user data
type User struct {
  ID    int    `json:"id"`
  Name  string `json:"name"`
  Email string `json:"email"`
  // ... other user details (optional)
  Role  string `json:"role,omitempty"` // Optional: If roles are managed
}

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

// RoleAssigner defines an interface for assigning roles (optional)
type RoleAssigner interface {
  AssignRole(userID int, role string) error
}

// UserManager implements a composite structure for user management
type UserManager struct {
  creator UserCreator
  getter  UserGetter
  updater UserUpdater
  deleter UserDeleter
  assigner RoleAssigner // Optional: If needed
}

// NewUserManager creates a new UserManager instance
func NewUserManager(
  creator UserCreator,
  getter UserGetter,
  updater UserUpdater,
  deleter UserDeleter,
  assigner RoleAssigner, // Optional: If needed
) *UserManager {
  return &UserManager{
    creator: creator,
    getter:  getter,
    updater: updater,
    deleter: deleter,
    assigner: assigner, // Optional: If needed
  }
}
```

**Benefits:**

- **Improved Maintainability:** Smaller interfaces are easier to understand, test, and modify.
- **Reduced Coupling:** Clients depend only on the interfaces they need, promoting loose coupling and flexibility.
- **Code Reusability:** Interfaces can be reused in different contexts, potentially across packages.
- **Optional Role Management:** The `RoleAssigner` interface can be omitted if role management is not required.

### Dependency Inversion Principle (DIP)

This principle states that high-level modules should, but rather both should depend on abstractions. This helps to reduce the coupling between components and make the code more flexible and maintainable.

Suppose we have a struct `Worker` that represents a worker in a company and a struct `Supervisor` that represents a supervisor:

```go
type Worker struct {  
  ID int  
  Name string  
}  
  
func (w Worker) GetID() int {  
  return w.ID  
}  
  
func (w Worker) GetName() string {  
  return w.Name  
}  
  
type Supervisor struct {  
  ID int  
  Name string  
}  
  
func (s Supervisor) GetID() int {  
  return s.ID  
}  
  
func (s Supervisor) GetName() string {  
  return s.Name  
}
```

Now, for the anti-pattern, let’s say we have a high-level module `Department` that represents a department in a company, and needs to store information about the workers and supervisors, which are considered a low-level module:

```go
type Department struct {  
  Workers []Worker  
  Supervisors []Supervisor  
}
```

According to the Dependency Inversion Principle, high-level modules should not depend on low-level modules. Instead, both should depend on abstractions. To fix the anti-pattern example, We can create an interface `Employee` that represents both, worker and supervisor:

```go
type Employee interface {  
  GetID() int  
  GetName() string  
}
```

Now We can update the `Department` struct so it no longer depends on low-level modules:

```go
type Department struct {  
  Employees []Employee  
}
```

Considering SOLID principles while code requires a shift in the way of thinking, but the benefits are well worth the effort. These principles can be complex and challenging while implementing the code, but they provide robust, scalable software that will continue to meet the needs for years to come.
