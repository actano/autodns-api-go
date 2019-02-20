package zone

import (
	"github.com/actano/autodns-api-go/pkg/api"
)

type zone struct {
	Name string `xml:"name"`
}

type ResourceRecord struct {
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
	TTL   uint   `xml:"ttl"`
}

type ZoneService struct {
	client api.Client
}

func NewZoneService(client api.Client) *ZoneService {
	return &ZoneService{
		client: client,
	}
}
