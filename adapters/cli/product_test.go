package cli_test

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/adapters/cli"
	"github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application"
	mock_application "github.com/roaires/fullcycle-arquitetura-hexagonal-golang/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Prod test CLI"
	productPrice := 1.99
	productStatus := application.ENABLED
	productId := "xpto"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := fmt.Sprintf(cli.RESULT_CREATE, productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(cli.RESULT_ENABLE, productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(cli.RESULT_DISABLE, productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf(cli.RESULT_GET, productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	_, err = cli.Run(service, "xpto", productId, "", 0)
	require.NotNil(t, err)
	require.Equal(t, cli.RESULT_ACAO_INVALIDA, err.Error())

}
