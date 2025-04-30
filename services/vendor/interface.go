package vendor

import (
	"brb-midsvc-platform/model"
)

type VendorService interface {
	CreateVendor(vendor *model.Vendor) error
}