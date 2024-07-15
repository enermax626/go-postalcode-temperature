package dao

import (
	"github.com/enermax626/go-postalcode-temperature/internal/model"
	"github.com/stretchr/testify/mock"
)

type AddressDaoMock struct {
	mock.Mock
}

func (m *AddressDaoMock) FindByPostalCode(postalCode string) (*model.Address, error) {
	args := m.Mock.Called(postalCode)
	if args.Get(0) != nil {
		// args.Get(0) will be the first return value, which is *model.Address
		return args.Get(0).(*model.Address), args.Error(1)
	}
	// If the first argument is nil, return nil for *model.Address and whatever error was set
	return nil, args.Error(1)
}
