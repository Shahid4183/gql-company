package models

// Department - schema for department table
type Department struct {
	ID       int    `json:"id"` // auto increamented
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Employee - schema for employee table
type Employee struct {
	ID     int        `json:"id"` // auto increamented
	Name   string     `json:"name"`
	Job    string     `json:"job"`
	Mgr    int        `json:"mgr"`
	Salary float64    `json:"salary"`
	DeptID int        `json:"deptId"`
	Dept   Department `json:"dept" gorm:"foreignkey:DeptID;association_foreignkey:ID"`
}

// NewEmployee - input structure for new employee
type NewEmployee struct {
	Name   string  `json:"name"`
	Job    string  `json:"job"`
	Mgr    int     `json:"mgr"`
	Salary float64 `json:"salary"`
	DeptID int     `json:"deptId"`
}

// NewDepartment - input structure for new employee
type NewDepartment struct {
	Name     string `json:"name"`
	Location string `json:"location"`
}

// UpdateEmployee - input structure for updating existing employee
type UpdateEmployee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Job    string  `json:"job"`
	Mgr    int     `json:"mgr"`
	Salary float64 `json:"salary"`
	DeptID int     `json:"deptId"`
}
