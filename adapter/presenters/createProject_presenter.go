package adapter

import (
	"marketplace/domain"
	"marketplace/usecases"
)

type CreateProjectPresenterImp struct{}

func NewCreateProjectPresenter() CreateProjectPresenterImp {
	return CreateProjectPresenterImp{}
}

func (cpp CreateProjectPresenterImp) Output(project domain.Project) usecases.ProjectOutput {

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
