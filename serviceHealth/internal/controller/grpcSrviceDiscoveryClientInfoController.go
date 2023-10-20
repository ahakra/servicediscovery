package controller

import (
	"github.com/ahakra/servicediscovery/pkg/serviceDiscoveryProto"
	repository "github.com/ahakra/servicediscovery/serviceHealth/internal/repo"
)

type ClientInfoCtrl struct {
	ClientInfoRepo repository.ClientInfoRepo
}

func NewClientInfoCltr(ci repository.ClientInfoRepo) *ClientInfoCtrl {
	return &ClientInfoCtrl{ClientInfoRepo: ci}
}

func (ci *ClientInfoCtrl) GetAllService() (*serviceDiscoveryProto.Services, error) {
	return ci.ClientInfoRepo.GetAllServices()

}

func (ci *ClientInfoCtrl) Close() {
	ci.ClientInfoRepo.Close()
}
