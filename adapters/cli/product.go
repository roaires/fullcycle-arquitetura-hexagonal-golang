package cli

import (
	"errors"
	"fmt"

	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application"
)

const (
	RESULT_CREATE        = "Prouto %s - %s criado com sucesso. \nPreço: %f\n Status: %s"
	RESULT_ENABLE        = "Produto %s com status enabled."
	RESULT_DISABLE       = "Produto %s com status disabled."
	RESULT_GET           = "Informações do produto:\nID: %s\nName: %s\nPrice: %f\nStatus: %s"
	RESULT_ACAO_INVALIDA = "Ação inválida"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {

	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(RESULT_CREATE,
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(RESULT_ENABLE, res.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		res, err := service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(RESULT_DISABLE, res.GetName())

	case "get":
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf(RESULT_GET,
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())

	default:
		return result, errors.New(RESULT_ACAO_INVALIDA)

	}

	return result, nil
}
