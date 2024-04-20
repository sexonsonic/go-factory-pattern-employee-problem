package employee

// Director is a struct for Director
// It embeds Employee
// This means that Director has all the fields and methods of Employee
type Director struct {
	Employee
	// TODO: Add the Employee struct
}

// NewDirector creates a new Director
// It returns a pointer to the Director
// Creational method
func NewDirector() *Director {
	// TODO: Create a new Director
	// Set the name to "Director"
	// Set the salary to 500
	return &Director {
		Employee: Employee {
			Name: "Director",
			Salary: 5000,
			Bonus: 0.3,
		},
	}
}

func (s *Director) GetBonus() float64 {
	return float64(s.Salary * int(s.Bonus))
}