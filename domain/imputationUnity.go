package domain

type ImputationUnity struct {
	ceco       Ceco
	cia        string
	percentage float64
}

func NewImputationUnity(ceco string, cia string, percentage float64) (ImputationUnity, error) {
	var newImputationUnity ImputationUnity

	newceco, errCeco := NewCeco(ceco)

	if errCeco != nil {
		return newImputationUnity, errCeco
	}

	return ImputationUnity{ceco: newceco, cia: cia, percentage: percentage}, nil
}

func (i ImputationUnity) GetCeco() Ceco {
	return i.ceco
}

func (i ImputationUnity) GetCia() string {
	return i.cia
}

func (i ImputationUnity) GetPercentage() float64 {
	return i.percentage
}
