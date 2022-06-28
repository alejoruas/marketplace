package repository

import (
	"context"
	"fmt"
	"marketplace/domain"
)

type ImputationUnitySQL struct {
	db SQL
}

func NewImputationUnitySql(db SQL) ImputationUnitySQL {
	return ImputationUnitySQL{db: db}
}

func (p ImputationUnitySQL) Create(ctx context.Context, project domain.Project, imputationUnity domain.ImputationUnity) (domain.ImputationUnity, error) {
	tx, ok := ctx.Value(ProjectTransactionContextKey).(Tx)
	if !ok {
		var err error
		tx, err = p.db.BeginTx(ctx)
		if err != nil {
			return domain.ImputationUnity{}, err
		}
	}

	var query = `INSERT INTO public."DriverImputation" (id_project, ceco, percentage, cia)
					values ($1, $2, $3, $4)`

	err := tx.ExecuteContext(ctx, query, project.GetId(), imputationUnity.GetCeco().Ceco(), imputationUnity.GetPercentage(), imputationUnity.GetCia())

	if err != nil {
		//TODO: Log the error in files o DB
		fmt.Println(err)
		return domain.ImputationUnity{}, err
	}

	return imputationUnity, nil
}

func (p ImputationUnitySQL) FindByProjectId(ctx context.Context, idProject string) ([]domain.ImputationUnity, error) {
	return nil, nil
}
