package vendor

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/store/vendor"
)

type vendorService struct {
	repo vendor.VendorRepository
}

func NewVendorService(r vendor.VendorRepository) VendorService {
	return &vendorService{repo: r}
}

func (s *vendorService) CreateVendor(vendor *model.Vendor) error {
	return s.repo.Create(vendor)
}
