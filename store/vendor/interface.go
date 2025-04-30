package vendor

import "brb-midsvc-platform/model"

type VendorRepository interface {
	Create(vendor *model.Vendor) error
}
