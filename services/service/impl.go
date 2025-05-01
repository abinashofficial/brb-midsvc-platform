package service

import (
	"brb-midsvc-platform/model"
	"brb-midsvc-platform/store/service"
	"errors"
)

type serviceService struct {
	repo service.ServiceRepository
}

func NewServiceService(r service.ServiceRepository) ServiceService {
	return &serviceService{repo: r}
}

func (s *serviceService) CreateService(service *model.Service) error {
	return s.repo.Create(service)
}

func (s *serviceService) UpdateService(service *model.Service, id string) error {
	data, err := s.repo.FindByServicesID(id)
	 if err != nil {
		return errors.New("service not found")	}
		service.ID = data.ID
		data.Name = service.Name
		data.Description = service.Description
		data.Price = service.Price

	return s.repo.Update(data)
}

func (s *serviceService) AssignVendor(data *model.LinkServiceVendor)error {

	_, err := s.repo.FindByID(	 model.ServiceVendor{
		VendorID: data.VendorID,
		ServiceID: data.ServiceID,
	 })
	if err == nil {
		return errors.New("service vendor already exists")
	}
	
	if err := s.repo.AssignVendor(data); err != nil {
		return  errors.New("vendor not found")
	}
	return nil
}

func (s *serviceService) ToggleService(data model.ServiceVendor) (*model.LinkServiceVendor, error) {
	serv, err := s.repo.FindByID(data)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Toggle(serv); err != nil {
		return nil, err
	}
	return serv, nil
}
