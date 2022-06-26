package repository

import (
	"context"
	"fmt"
	"marketplace/domain"
)

type ProjectSQL struct {
	db SQL
}

type contextKey int

const (
	ProjectTransactionContextKey contextKey = iota
)

func NewProjectSql(db SQL) ProjectSQL {
	return ProjectSQL{db: db}
}

func (p ProjectSQL) Create(ctx context.Context, project domain.Project) (domain.Project, error) {
	tx, ok := ctx.Value(ProjectTransactionContextKey).(Tx)
	if !ok {
		var err error
		tx, err = p.db.BeginTx(ctx)
		if err != nil {
			return domain.Project{}, err
		}
	}

	var query = `INSERT INTO public."Project" (Id, Name, Budget)
					values ($1, $2, $3)`

	err := tx.ExecuteContext(ctx, query, project.GetId(), project.GetName(), project.GetBudget())

	if err != nil {
		//TODO: Log the error in files o DB
		fmt.Println(err)
		return domain.Project{}, err
	}

	return project, nil
}

func (p ProjectSQL) FindById(context.Context, string) (domain.Project, error) {
	return domain.Project{}, nil
}
func (p ProjectSQL) FindAll(context.Context) ([]domain.Project, error) {
	return nil, nil
}

func (p ProjectSQL) WithTransaction(ctx context.Context, fn func(ctxTx context.Context) error) error {
	tx, err := p.db.BeginTx(ctx)
	if err != nil {
		return err
	}

	ctxTx := context.WithValue(ctx, ProjectTransactionContextKey, tx)
	err = fn(ctxTx)

	if err != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
