package controller

import (
	"time"

	repository "github.com/ahakra/servicediscovery/serviceHealth/internal/repo"
	"github.com/gofiber/fiber/v2"
)

var GRPCClient *repository.ClientInfoRepo

type Services struct {
	Servicename    string
	Serviceaddress string
	LastUpdate     time.Time
	Messages       []string
	Status         string
}

func GetServiceStatus(seconds int) string {
	if int(time.Now().Unix())-seconds > 20 {
		return "UnHealthy"
	}
	return "Healthy"
}

func GetAllServicesData(c *fiber.Ctx) error {

	ctrl := NewClientInfoCltr(*GRPCClient)
	services, _ := ctrl.GetAllService()

	listOfServices := []Services{}

	for _, service := range services.Services {
		listOfServices = append(listOfServices,
			Services{
				Servicename:    service.Servicename,
				Serviceaddress: service.Serviceaddress,
				Messages:       service.Messages,
				LastUpdate:     time.Unix(service.Lastupdate.Seconds, 0),
				Status:         GetServiceStatus(int(service.Lastupdate.Seconds)),
			},
		)
	}

	return c.Render("index", fiber.Map{
		"Services": listOfServices,
	})
}
