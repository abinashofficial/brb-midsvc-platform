package service

import "brb-midsvc-platform/model"

type ServiceRepository interface {
	Create(service *model.Service) error
	Update(service *model.Service) error
	AssignVendor(service *model.LinkServiceVendor) error
	FindByID(data model.ServiceVendor) (*model.LinkServiceVendor, error)
	Toggle(service *model.LinkServiceVendor) error
	FindByServicesID(id string) (*model.Service, error)
}
