package gql_company

import (
	"context"
	"fmt"

	"github.com/Shahid4183/gql-company/models"
	"github.com/jinzhu/gorm"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

// Resolver - generic type for query resolver
type Resolver struct {
	db *gorm.DB
}

// Employee - creates a employee resolver
func (r *Resolver) Employee() EmployeeResolver {
	r.db = GetDBInstance()
	return &employeeResolver{r}
}

// Mutation - creates and returns mutation resolver
func (r *Resolver) Mutation() MutationResolver {
	r.db = GetDBInstance()
	return &mutationResolver{r}
}

// Query - creates and returns query resolver
func (r *Resolver) Query() QueryResolver {
	r.db = GetDBInstance()
	return &queryResolver{r}
}

type employeeResolver struct{ *Resolver }

func (r *employeeResolver) Dept(ctx context.Context, obj *models.Employee) (*models.Department, error) {
	dept := &models.Department{}
	if err := r.db.Where("id = ?", obj.DeptID).Find(dept).Error; err != nil {
		return nil, err
	}
	return dept, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateEmployee(ctx context.Context, input models.NewEmployee) (*models.Employee, error) {
	e := &models.Employee{
		Name:   input.Name,
		Salary: input.Salary,
		Job:    input.Job,
		Mgr:    input.Mgr,
		DeptID: input.DeptID,
	}
	if err := r.db.Create(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *mutationResolver) CreateDepartment(ctx context.Context, input models.NewDepartment) (*models.Department, error) {
	dept := &models.Department{
		Name:     input.Name,
		Location: input.Location,
	}
	if err := r.db.Create(dept).Error; err != nil {
		return nil, err
	}
	return dept, nil
}

func (r *mutationResolver) UpdateEmployee(ctx context.Context, input models.UpdateEmployee) (*models.Employee, error) {
	e := &models.Employee{
		ID:     input.ID,
		Name:   input.Name,
		Salary: input.Salary,
		Job:    input.Job,
		Mgr:    input.Mgr,
		DeptID: input.DeptID,
	}
	if err := r.db.Save(e).Error; err != nil {
		return nil, err
	}
	return e, nil
}

func (r *mutationResolver) DeleteEmployee(ctx context.Context, id int) (string, error) {
	if err := r.db.Where("id = ?", id).Delete(&models.Employee{}).Error; err != nil {
		return "", err
	}
	return fmt.Sprintf("Employee record with id %d has been deleted successfully!", id), nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Employees(ctx context.Context) ([]*models.Employee, error) {
	var list []*models.Employee
	if err := r.db.Preload("Dept").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *queryResolver) EmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	emp := &models.Employee{}
	if err := r.db.Where("id = ?", id).Find(emp).Error; err != nil {
		return nil, err
	}
	return emp, nil
}

func (r *queryResolver) Departments(ctx context.Context) ([]*models.Department, error) {
	var list []*models.Department
	if err := r.db.Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *queryResolver) EmployeeByDeptID(ctx context.Context, deptID int) ([]*models.Employee, error) {
	var list []*models.Employee
	if err := r.db.Where("dept_id = ?", deptID).Preload("Dept").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}
