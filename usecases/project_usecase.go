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
		repositoryProject          domain.ProjectRepository
		repositoryDriverImputation domain.ImputationUnityRepository
		presenter                  CreateProjectPresenter
		ctxTimeOut                 time.Duration
	}
)

func NewProjectInteractor(
	repositoryProject domain.ProjectRepository,
	repositoryDriverImputation domain.ImputationUnityRepository,
	presenter CreateProjectPresenter,
	t time.Duration) CreateProjectUseCase {
	return CreateProjectInteractor{
		repositoryProject:          repositoryProject,
		repositoryDriverImputation: repositoryDriverImputation,
		presenter:                  presenter,
		ctxTimeOut:                 t}
}

func (cpi CreateProjectInteractor) Execute(ctx context.Context, projectInput ProjectInput) (ProjectOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, cpi.ctxTimeOut)
	defer cancel()

	var (
		project domain.Project
		err     error
	)

	err = cpi.repositoryProject.WithTransaction(ctx, func(ctxTx context.Context) error {
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

		project, err = cpi.repositoryProject.Create(ctxTx, project)

		if err != nil {
			fmt.Println(err)
			return err
		}

		for _, di := range project.GetDriverImputacion() {
			_, err = cpi.repositoryDriverImputation.Create(ctxTx, project, di)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println(err)
		return cpi.presenter.Output(domain.Project{}), nil
	}

	return cpi.presenter.Output(project), nil
}
