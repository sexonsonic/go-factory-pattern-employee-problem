package employee

// Manager is a struct for manager
// It embeds Employee
// This means that Manager has all the fields and methods of Employee
type Manager struct {
	Employee
	// TODO: Add the Employee struct
}

// NewManager creates a new manager
// It returns a pointer to the manager
// Creational method
func NewManager() *Manager {
	// TODO: Create a new manager
	// Set the name to "Manager"
	// Set the salary to 1000
	return &Manager {
		Employee: Employee {
			Name: "Manager",
			Salary: 1000,
			Bonus: 0.2,
		},
	}
}

func (m *Manager) GetBonus() float64 {
	return float64(m.Salary * int(m.Bonus))
}