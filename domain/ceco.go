package domain

import (
	"errors"
	"strconv"
)

type Ceco struct {
	idCeco string
}

var (
	ErrCecoFormat    = errors.New("The CECO has an invalid format")
	ErrCecoNoCompany = errors.New("The CECO does not belong to a company")
)

func NewCeco(idCeco string) (Ceco, error) {
	var ceco Ceco

	if len(idCeco) != 9 {
		return ceco, ErrCecoFormat
	}

	_, err1 := strconv.Atoi(idCeco)
	if err1 != nil {
		return ceco, ErrCecoFormat
	}

	if BelongAnyCompany(Ceco{idCeco: idCeco}) == false {
		return ceco, ErrCecoNoCompany
	}

	ceco = Ceco{idCeco: idCeco}

	return ceco, nil
}

func (c Ceco) Ceco() string {
	return c.idCeco
}
