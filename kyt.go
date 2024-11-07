package chainalysis

import (
	"fmt"

	"github.com/bytedance/sonic"
)

// https://docs.chainalysis.com/api/kyt/#how-it-works
// https://docs.chainalysis.com/api/kyt/guides/#supported-networks-and-assets

const (
	urlKytRegisterTransfer            = "/api/kyt/v2/users/%s/transfers"
	urlKytGetTransferSummary          = "/api/kyt/v2/transfers/%s"
	urlKytGetTransferAlerts           = "/api/kyt/v2/transfers/%s/alerts"
	urlKycRegisterWithdrawalAttempt   = "/api/kyt/v2/users/%s/withdrawal-attempts"
	urlKytGetWithdrawalAttemptSummary = "/api/kyt/v2/withdrawal-attempts/%s"
	urlKytGetWithdrawalAttemptAlerts  = "/api/kyt/v2/withdrawal-attempts/%s/alerts"
)

func (c *ClientImpl) KYTRegisterTransfer(userId string, param KYTRegisterTransferParam) (resp KYTRegisterTransferResp, err error) {
	body, err := sonic.MarshalString(param)
	if err != nil {
		return resp, err
	}

	_, err = c.client.R().
		SetBody(body).
		SetResult(&resp).
		SetError(&resp).
		Post(fmt.Sprintf(urlKytRegisterTransfer, userId))
	return
}

func (c *ClientImpl) KYTGetTransferSummary(externalId string) (resp KYTGetTransferSummaryResp, err error) {
	rsp, err := c.client.R().
		SetResult(&resp).
		Get(fmt.Sprintf(urlKytGetTransferSummary, externalId))
	if err != nil {
		return resp, err
	}
	if rsp.IsError() {
		return resp, fmt.Errorf("%s", string(rsp.Body()))
	}
	return
}

func (c *ClientImpl) KYTGetTransferAlerts(externalId string) (resp KYTGetTransferAlertsResp, err error) {
	rsp, err := c.client.R().
		SetResult(&resp).
		Get(fmt.Sprintf(urlKytGetTransferAlerts, externalId))
	if err != nil {
		return resp, err
	}
	if rsp.IsError() {
		return resp, fmt.Errorf("%s", string(rsp.Body()))
	}
	return
}

func (c *ClientImpl) KYTRegisterWithdrawalAttempt(userId string, param KYTRegisterWithdrawalAttemptParam) (resp KYTRegisterWithdrawalAttemptResp, err error) {
	body, err := sonic.MarshalString(param)
	if err != nil {
		return resp, err
	}

	_, err = c.client.R().
		SetBody(body).
		SetResult(&resp).
		SetError(&resp).
		Post(fmt.Sprintf(urlKycRegisterWithdrawalAttempt, userId))
	return
}

func (c *ClientImpl) KYTGetWithdrawalAttemptSummary(externalId string) (resp KYTGetWithdrawalAttemptSummaryResp, err error) {
	rsp, err := c.client.R().
		SetResult(&resp).
		Get(fmt.Sprintf(urlKytGetWithdrawalAttemptSummary, externalId))
	if err != nil {
		return resp, err
	}
	if rsp.IsError() {
		return resp, fmt.Errorf("%s", string(rsp.Body()))
	}
	return
}

func (c *ClientImpl) KYTGetWithdrawalAttemptAlerts(externalId string) (resp KYTGetTransferAlertsResp, err error) {
	rsp, err := c.client.R().
		SetResult(&resp).
		Get(fmt.Sprintf(urlKytGetWithdrawalAttemptAlerts, externalId))
	if err != nil {
		return resp, err
	}
	if rsp.IsError() {
		return resp, fmt.Errorf("%s", string(rsp.Body()))
	}
	return
}

type KYTRegisterTransferParam struct {
	Network           string `json:"network"`
	Asset             string `json:"asset"`
	TransferReference string `json:"transferReference"`
	Direction         string `json:"direction"`
	// optional
	AssetId           string   `json:"assetId,omitempty"`
	TransferTimestamp string   `json:"transferTimestamp,omitempty"` // ISO8601 2022-03-10T20:37:32+00:00
	AssetAmount       string   `json:"assetAmount,omitempty"`
	OutputAddress     string   `json:"outputAddress,omitempty"`
	InputAddresses    []string `json:"inputAddress,omitempty"`
	AssetPrice        float64  `json:"assetPrice,omitempty"`
	AssetDenomination string   `json:"assetDenomination,omitempty"`
}

type KYTRegisterTransferResp struct {
	ErrorResp
	UpdatedAt         string  `json:"updatedAt"`
	Asset             string  `json:"asset"`
	AssetId           string  `json:"assetId"`
	Network           string  `json:"network"`
	TransferReference string  `json:"transferReference"`
	Tx                string  `json:"tx"`
	Idx               int     `json:"idx"`
	UsdAmount         float64 `json:"usdAmount"`
	AssetAmount       float64 `json:"assetAmount"`
	Timestamp         string  `json:"timestamp"`
	OutputAddress     string  `json:"outputAddress"`
	ExternalId        string  `json:"externalId"`
}

type KYTRegisterWithdrawalAttemptParam struct {
	Network           string  `json:"network"`
	Asset             string  `json:"asset"`
	Address           string  `json:"address"`
	AttemptIdentifier string  `json:"attemptIdentifier"`
	AssetAmount       float64 `json:"assetAmount"`
	AttemptTimestamp  string  `json:"attemptTimestamp"` // UTC ISO 8601: 2020-12-09T17:25:40.008307
	// Optional
	AssetId           string  `json:"assetId,omitempty"`
	AssetPrice        float64 `json:"assetPrice,omitempty"`
	AssetDenomination string  `json:"assetDenomination,omitempty"`
}

type KYTRegisterWithdrawalAttemptResp struct {
	ErrorResp
	Asset             string  `json:"asset"`
	AssetId           string  `json:"assetId"`
	Network           string  `json:"network"`
	Address           string  `json:"address"`
	AttemptIdentifier string  `json:"attemptIdentifier"`
	AssetAmount       float64 `json:"assetAmount"`
	UsdAmount         float64 `json:"usdAmount"`
	UpdatedAt         string  `json:"updatedAt"`
	ExternalId        string  `json:"externalId"`
}

type KYTGetTransferSummaryResp struct {
	ErrorResp
	UpdatedAt         string  `json:"updatedAt"`
	Asset             string  `json:"asset"`
	AssetId           string  `json:"assetId"`
	Network           string  `json:"network"`
	TransferReference string  `json:"transferReference"`
	Tx                string  `json:"tx"`
	Idx               int     `json:"idx"`
	UsdAmount         float64 `json:"usdAmount"`
	AssetAmount       float64 `json:"assetAmount"`
	Timestamp         string  `json:"timestamp"`
	OutputAddress     string  `json:"outputAddress"`
	ExternalId        string  `json:"externalId"`
}

type KYTGetWithdrawalAttemptSummaryResp struct {
	ErrorResp
	Asset             string  `json:"asset"`
	AssetId           string  `json:"assetId"`
	Network           string  `json:"network"`
	Address           string  `json:"address"`
	AttemptIdentifier string  `json:"attemptIdentifier"`
	AssetAmount       float64 `json:"assetAmount"`
	UsdAmount         float64 `json:"usdAmount"`
	UpdatedAt         string  `json:"updatedAt"`
	ExternalId        string  `json:"externalId"`
}

type KYTGetTransferAlertsResp struct {
	ErrorResp
	Alerts []KYTTransferAlert `json:"alerts"`
}

type KYTTransferAlert struct {
	AlertLevel   string  `json:"alertLevel"` // SEVERE, HIGH, MEDIUM, LOW
	Service      string  `json:"service"`
	ExternalId   string  `json:"externalId"`
	AlertAmount  float64 `json:"alertAmount"`
	ExposureType string  `json:"exposureType"`
	CategoryId   int64   `json:"categoryId"`
}

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
