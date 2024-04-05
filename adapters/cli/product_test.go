package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/matheus-santos-souza/go-hexagonal-architecture/adapters/cli"
	mock_application "github.com/matheus-santos-souza/go-hexagonal-architecture/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	serviceMock := mock_application.NewMockProductServiceInterface(ctrl)
	serviceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	serviceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	createExpected := fmt.Sprintf(
		"Product ID %s, white the name %s has been created with the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)
	result, err := cli.Run(serviceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, createExpected, result)

	enabledExpected := fmt.Sprintf(
		"Product %s has been enabled",
		productName,
	)
	result, err = cli.Run(serviceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, enabledExpected, result)

	disabledExpected := fmt.Sprintf(
		"Product %s has been disabled",
		productName,
	)
	result, err = cli.Run(serviceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, disabledExpected, result)

	getExpected := fmt.Sprintf(
		"Product ID %s, white the name %s has been get with the price %f and status %s",
		productId, productName, productPrice, productStatus,
	)
	result, err = cli.Run(serviceMock, "get", productId, "", 0.0)
	require.Nil(t, err)
	require.Equal(t, getExpected, result)

}
