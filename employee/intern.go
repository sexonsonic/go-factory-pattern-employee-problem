package employee

// Intern is a struct for intern
// It embeds Employee
// This means that Intern has all the fields and methods of Employee
type Intern struct {
	Employee
	// TODO: Add the Employee struct
}

// NewIntern creates a new intern
// It returns a pointer to the intern
// Creational method
func NewIntern() *Intern {
	// TODO: Create a new intern
	// Set the name to "Intern"
	// Set the salary to 100
	return &Intern {
		Employee: Employee {
			Name: "Intern",
			Salary: 100,
		},
	}
}

func (i *Intern) GetBonus() float64 {
	return float64(i.Salary)
}