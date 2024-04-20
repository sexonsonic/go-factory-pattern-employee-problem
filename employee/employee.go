package employee

// EmployeeInterface is an interface for employee
type EmployeeInterface interface {
	SetName(name string)
	SetSalary(salary int)
	GetName() string
	GetSalary() int
	/*
		TODO: Add a new method called GetBonus() that returns a float64
	*/
	GetBonus() float64
}

// Employee is a struct for employee
type Employee struct {
	Name   string
	Salary int
	Bonus  float64
}

// SetName sets the name of the employee
func (e *Employee) SetName(name string) {
	e.Name = name
	// TODO: Set the name of the employee
}

// SetSalary sets the salary of the employee
func (e *Employee) SetSalary(salary int) {
	e.Salary = salary
	// TODO: Set the salary of the employee
}

// GetName gets the name of the employee
func (e *Employee) GetName() string {
	return e.Name
	// TODO: Get the name of the employee
}

// GetSalary gets the salary of the employee
func (e *Employee) GetSalary() int {
	return e.Salary
	// TODO: Get the salary of the employee
}

// GetBonus gets the bonus
func (e *Employee) GetBonus() float64 {
	return 0.0
}
