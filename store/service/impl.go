package service

import (
	"brb-midsvc-platform/model"
	"gorm.io/gorm"
)

type serviceRepository struct {
	db *gorm.DB
}

func NewServiceRepository(db *gorm.DB) ServiceRepository {
	return &serviceRepository{db}
}

func (r *serviceRepository) Create(service *model.Service) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) Update(service *model.Service) error {
	return r.db.Save(service).Error
}

func (r *serviceRepository) AssignVendor(service *model.LinkServiceVendor) error {
	return r.db.Create(service).Error
}

func (r *serviceRepository) FindByID(data model.ServiceVendor) (*model.LinkServiceVendor, error) {
	var service model.LinkServiceVendor
	err := r.db.Where("vendor_id = ? AND service_id = ?", data.VendorID, data.ServiceID).First(&service).Error
	if err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *serviceRepository) FindByServicesID(id string) (*model.Service, error) {
	var service model.Service
	if err := r.db.First(&service, id).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (r *serviceRepository) Toggle(service *model.LinkServiceVendor) error {
	service.Active = !service.Active
	return r.db.Save(service).Error
}
