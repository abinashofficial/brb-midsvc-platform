package service
import(
	"brb-midsvc-platform/model"
)

type ServiceService interface {
	CreateService(service *model.Service) error
	UpdateService(service *model.Service, id string) error
	AssignVendor(data *model.LinkServiceVendor)  error
	ToggleService(id model.ServiceVendor) (*model.LinkServiceVendor, error)
}