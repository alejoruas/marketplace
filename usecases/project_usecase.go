package usecases

import (
	"context"
	"marketplace/domain"
	"time"
)

type (

	//input output data
	ImputationUnityInOutput struct {
		Ceco       string  `json:"ceco"`
		Cia        string  `json:"cia"`
		Percentage float64 `json:"percentage"`
	}

	//Input data
	ProjectInput struct {
		Name             string                    `json:"name"`
		Budget           float64                   `json:"budget"`
		DriverImputation []ImputationUnityInOutput `json:"driverImputation"`
	}

	//Output data
	ProjectOutput struct {
		Id               string                    `json:"id"`
		Name             string                    `json:"name"`
		Budget           float64                   `json:"budget"`
		DriverImputation []ImputationUnityInOutput `json:"driverImputation"`
	}

	//Input port
	CreateProjectUseCase interface {
		Execute(context.Context, ProjectInput) (ProjectOutput, error)
	}

	//Output port
	CreateProjectPresenter interface {
		Output(domain.Project) ProjectOutput
	}

	CreateProjectInteractor struct {
		repository domain.ProjectRepository
		presenter  CreateProjectPresenter
		ctxTimeOut time.Duration
	}
)

func NewProjectInteractor(repository domain.ProjectRepository, presenter CreateProjectPresenter, t time.Duration) CreateProjectUseCase {
	return CreateProjectInteractor{repository: repository, presenter: presenter, ctxTimeOut: t}
}

func (cpi CreateProjectInteractor) Execute(ctx context.Context, projectInput ProjectInput) (ProjectOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, cpi.ctxTimeOut)
	defer cancel()

	project, err1 := domain.NewProject(projectInput.Name, projectInput.Budget)

	if err1 != nil {
		return cpi.presenter.Output(domain.Project{}), err1
	}

	var driverImputacion []domain.ImputationUnity

	for i := 0; i < len(projectInput.DriverImputation); i++ {
		di, _ := domain.NewImputationUnity(
			projectInput.DriverImputation[i].Ceco,
			projectInput.DriverImputation[i].Cia,
			projectInput.DriverImputation[i].Percentage)

		driverImputacion = append(driverImputacion, di)
	}

	project.SetDriverImputation(driverImputacion)

	//TODO: call the Repositry

	project.SetId("4343434")

	return cpi.presenter.Output(project), nil
}
