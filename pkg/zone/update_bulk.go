package zone

import (
	"encoding/xml"
	"errors"
	"github.com/actano/autodns-api-go/pkg/api"
)

type task struct {
	Code    string           `xml:"code"`
	Zone    zone             `xml:"zone"`
	Adds    []ResourceRecord `xml:"default>rr_add"`
	Removes []ResourceRecord `xml:"default>rr_rem"`
}

type updateBulkRequest struct {
	XMLName xml.Name `xml:"request"`
	Auth    api.Auth `xml:"auth"`
	Task    task     `xml:"task"`
}

type updateBulkResponse struct {
	XMLName xml.Name           `xml:"response"`
	Status  api.ResponseStatus `xml:"result>status"`
}

func newUpdateBulkRequest(zoneName string, adds []ResourceRecord, removes []ResourceRecord, auth api.Auth) updateBulkRequest {
	return updateBulkRequest{
		Auth: auth,
		Task: task{
			Code: "0202001",
			Zone: zone{
				Name: zoneName,
			},
			Adds:    adds,
			Removes: removes,
		},
	}
}

func UpdateBulk(zoneName string, adds []ResourceRecord, removes []ResourceRecord, auth api.Auth) error {
	request := newUpdateBulkRequest(zoneName, adds, removes, auth)
	data, err := xml.Marshal(request)

	if err != nil {
		return err
	}

	response, err := api.MakeRequest(data)

	if err != nil {
		return err
	}

	updateBulkResponse := &updateBulkResponse{}
	err = xml.Unmarshal(response, updateBulkResponse)

	if err != nil {
		return err
	}

	if updateBulkResponse.Status.Type != "success" {
		return errors.New("UpdateBulk was not successful")
	}

	return nil
}
