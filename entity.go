package chainalysis

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// https://docs.chainalysis.com/api/address-screening

const (
	urlEntityRegisterAddress = "/api/risk/v2/entities"
	urlEntityRetrieveAddress = "/api/risk/v2/entities/%s"
)

func (c *ClientImpl) EntityAddressRegister(address string) (resp EntityAddressRegisterResp, err error) {
	body, err := sonic.MarshalString(map[string]string{
		"address": address,
	})
	if err != nil {
		return resp, err
	}

	_, err = c.client.R().
		SetBody(body).
		SetResult(&resp).
		SetError(&resp).
		Post(urlEntityRegisterAddress)
	return
}

func (c *ClientImpl) EntityAddressRetrieve(address string) (resp EntityAddressRetrieveResp, err error) {
	_, err = c.client.R().
		SetResult(&resp).
		SetError(&resp).
		Get(fmt.Sprintf(urlEntityRetrieveAddress, address))
	return
}

type EntityAddressRegisterResp struct {
	ErrorResp
	Address string `json:"address"`
}

type EntityAddressRetrieveResp struct {
	Message string `json:"message"`

	Address    string `json:"address"`
	Risk       string `json:"risk"` // Severe, High, Medium, Low
	RiskReason string `json:"riskReason"`
	Cluster    struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	} `json:"cluster"` // VASP?
	AddressType            string                  `json:"addressType"`
	AddressIdentifications []AddressIdentification `json:"addressIdentifications"`
	Exposures              []Exposure              `json:"exposures"`
	Triggers               []Trigger               `json:"triggers"`
}

type AddressIdentification struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Url         string `json:"url"`
	CreatedAt   int64  `json:"createdAt"` // unix*1e3
}

type Exposure struct {
	Category string  `json:"category"`
	Value    float64 `json:"value"`
}

type Trigger struct {
	Category      string  `json:"category"`
	Percentage    float64 `json:"percentage"`
	Message       string  `json:"message"`
	RuleTriggered struct {
		Risk         string  `json:"risk"`
		MinThreshold float64 `json:"minThreshold"`
		MaxThreshold float64 `json:"maxThreshold"`
		ExposureType string  `json:"exposureType"`
		Direction    string  `json:"direction"`
	} `json:"ruleTriggered"`
}
