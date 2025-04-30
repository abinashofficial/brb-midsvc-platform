package mock_repos

import (
	"github.com/stretchr/testify/mock"
	"brb-midsvc-platform/model"
)

type MockVendorRepo struct {
	mock.Mock	
}

func (m *MockVendorRepo) Create(vendor *model.Vendor) error {
	args := m.Called(vendor)
	return  args.Error(0)
}