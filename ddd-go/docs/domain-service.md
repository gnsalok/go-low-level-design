In Domain-Driven Design (DDD), **domain services** are used to encapsulate business logic that doesn't naturally belong to a specific **entity** or **value object**. Here's what that means:

---

### **Why Domain Services?**
- Sometimes, certain business logic involves multiple entities or value objects, or it doesn't logically fit within the scope of a single entity or value object.
- In such cases, you use **domain services** to handle this logic while keeping the domain model clean and focused.

---

### **Characteristics of Domain Services**
1. **Stateless**: Domain services are typically stateless. They don't hold any internal state and only operate on the data passed to them.
2. **Business-Oriented**: They encapsulate business rules or operations that are part of the domain logic.
3. **Operate on Entities/Value Objects**: They often work with entities and value objects to perform operations.

---

### **Example**
Imagine you have a system where users can transfer money between accounts. This operation involves two entities: `Account` (source) and `Account` (destination). The logic for transferring money doesn't belong to a single `Account` entity because:
- It involves two accounts.
- The operation is more about the interaction between the accounts than the behavior of a single account.

In this case, you can create a **domain service** like `MoneyTransferService`:

```go
package services

import (
	"ddd-go/internal/domain/entities"
	"errors"
)

type MoneyTransferService struct{}

func NewMoneyTransferService() *MoneyTransferService {
	return &MoneyTransferService{}
}

func (s *MoneyTransferService) TransferMoney(from *entities.Account, to *entities.Account, amount float64) error {
	if from.Balance < amount {
		return errors.New("insufficient funds")
	}

	from.Balance -= amount
	to.Balance += amount

	return nil
}
```

---

### **Why Not Put This in the `Account` Entity?**
- The `Account` entity should focus on its own behavior (e.g., deposit, withdraw).
- The transfer operation involves two accounts, so it doesn't naturally belong to a single `Account` entity.
- By using a domain service, you keep the `Account` entity focused and avoid bloating it with unrelated logic.

---

### **When to Use Domain Services**
- When the logic spans multiple entities or value objects.
- When the logic doesn't belong to a single entity or value object.
- When you want to keep entities and value objects focused on their primary responsibilities.
