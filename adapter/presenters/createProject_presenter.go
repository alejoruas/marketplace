package adapter

import (
	"marketplace/domain"
	"marketplace/usecases"
)

type createProjectPresenter struct{}

func NewCreCreateProjectPresenter() usecases.CreateProjectPresenter {
	return createProjectPresenter{}
}

func (cpp createProjectPresenter) Output(project domain.Project) usecases.ProjectOutput {

	var driverImputacion []usecases.ImputationUnityInOutput

	for _, di := range project.GetDriverImputacion() {
		driverImputacion = append(driverImputacion, usecases.ImputationUnityInOutput{Ceco: di.GetCeco().Ceco(), Cia: di.GetCia(), Percentage: di.GetPercentage()})
	}

	return usecases.ProjectOutput{
		Id:               project.GetId(),
		Name:             project.GetName(),
		Budget:           project.GetBudget(),
		DriverImputation: driverImputacion,
	}
}
