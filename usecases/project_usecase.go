package usecases

import (
	"context"
	"fmt"
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

	var (
		project domain.Project
		err     error
	)

	err = cpi.repository.WithTransaction(ctx, func(ctxTx context.Context) error {
		project, err = domain.NewProject(projectInput.Name, projectInput.Budget)

		if err != nil {
			return err
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

		project, err = cpi.repository.Create(ctxTx, project)

		//TODO: create driverImputation

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return cpi.presenter.Output(domain.Project{}), nil
	}

	return cpi.presenter.Output(project), nil
}
