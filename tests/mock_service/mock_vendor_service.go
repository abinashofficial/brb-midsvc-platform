package mock_service

import (
	"github.com/stretchr/testify/mock"
	"brb-midsvc-platform/model"

)

type MockVendorService struct {		
	mock.Mock
}

func (m *MockVendorService) CreateVendor(vendor *model.Vendor) error {
	args := m.Called(vendor)
	return  args.Error(0)
}