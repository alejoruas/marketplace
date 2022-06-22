package domain

import (
	"errors"
	"marketplace/infrastructure/utils"
)

var Companies = []string{"20", "21", "22", "26", "28"}

var (
	ErrCompanyFormat   = errors.New("Id company is invalid")
	ErrCompanyNotExist = errors.New("Id company does not exist")
)

type Company struct {
	idCompany string
}

func NewCompany(idCompany string) (Company, error) {
	var company Company
	if len(idCompany) != 2 {
		return company, ErrCompanyFormat
	}

	if !utils.Contains(Companies, idCompany) {
		return company, ErrCompanyNotExist
	}

	return Company{idCompany: idCompany}, nil
}

func BelongAnyCompany(ceco Ceco) bool {
	return utils.Contains(Companies, ceco.idCeco[:2])
}
