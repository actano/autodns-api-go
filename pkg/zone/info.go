package zone

import (
	"encoding/xml"
	"errors"
	"github.com/actano/autodns-api-go/pkg/api"
)

type zoneTask struct {
	Code string `xml:"code"`
	Zone zone   `xml:"zone"`
}

type zoneInfoRequest struct {
	XMLName xml.Name `xml:"request"`
	Auth    api.Auth `xml:"auth"`
	Task    zoneTask `xml:"task"`
}

type zoneInfoResponse struct {
	XMLName xml.Name           `xml:"response"`
	Records []ResourceRecord   `xml:"result>data>zone>rr"`
	Status  api.ResponseStatus `xml:"result>status"`
}

type ZoneInfo struct {
	Records []ResourceRecord
}

func newZoneInfoRequest(zoneName string, auth api.Auth) zoneInfoRequest {
	return zoneInfoRequest{
		Auth: auth,
		Task: zoneTask{
			Code: "0205",
			Zone: zone{
				Name: zoneName,
			},
		},
	}
}

func GetZoneInfo(zoneName string, auth api.Auth) (*ZoneInfo, error) {
	request := newZoneInfoRequest(zoneName, auth)
	data, err := xml.Marshal(request)

	if err != nil {
		return nil, err
	}

	response, err := api.MakeRequest(data)

	if err != nil {
		return nil, err
	}

    zoneInfoResponse := &zoneInfoResponse{}
	err = xml.Unmarshal(response, zoneInfoResponse)

	if err != nil {
		return nil, err
	}

	if zoneInfoResponse.Status.Type != "success" {
		return nil, errors.New("GetZoneInfo was not successful")
	}

	zoneInfo := &ZoneInfo{
		Records: zoneInfoResponse.Records,
	}

	return zoneInfo, nil
}
