package domain

import (
	"context"
	"errors"
)

var (
	ErrProductStockNegative = errors.New("Negative inventory")
)

type (
	Project struct {
		id               string
		name             string
		budget           float64
		driverImputacion []ImputationUnity
	}

	ProjectRepository interface {
		Create(context.Context, Project) (Project, error)
		FindById(context.Context, string) (Project, error)
		FindAll(context.Context) ([]Project, error)
	}
)

func NewProject(name string, budget float64) (Project, error) {
	var newProject Project = Project{name: name, budget: budget}
	return newProject, nil
}

func (p Project) GetId() string {
	return p.id
}

func (p Project) GetName() string {
	return p.name
}

func (p Project) GetBudget() float64 {
	return p.budget
}

func (p Project) GetDriverImputacion() []ImputationUnity {
	return p.driverImputacion
}

func (p *Project) SetId(id string) {
	p.id = id
}

func (p *Project) SetDriverImputation(driver []ImputationUnity) {
	p.driverImputacion = driver
}
