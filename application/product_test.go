package application_test

import (
	"testing"

	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{
		Name:   "Produto 1",
		Status: application.DISABLED,
		Price:  10,
	}

	err := product.Enable()
	require.Nil(t, err)

	product.Price = -5
	err = product.Enable()
	require.Equal(t, application.ERROR_PRICE_ENABLE, err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{
		Name:   "Produto 1",
		Status: application.DISABLED,
		Price:  0,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 99
	err = product.Disable()
	require.Equal(t, application.ERROR_PRICE_DISABLE, err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{
		ID:     uuid.NewV4().String(),
		Name:   "Produto 1",
		Status: "",
		Price:  10,
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "xpto"
	_, err = product.IsValid()
	require.Equal(t, application.ERROR_STATUS, err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -20
	_, err = product.IsValid()
	require.Equal(t, application.ERROR_PRICE_NEVATIVE, err.Error())
}
