package chainalysis

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

const baseUrl = "https://api.chainalysis.com"
const iso8601 = "2006-01-02T15:04:05.000000"

type Client interface {
	EntityAddressRegister(address string) (resp EntityAddressRegisterResp, err error)
	EntityAddressRetrieve(address string) (resp EntityAddressRetrieveResp, err error)

	KYTRegisterTransfer(userId string, param KYTRegisterTransferParam) (resp KYTRegisterTransferResp, err error)
	KYTGetTransferSummary(externalId string) (resp KYTGetTransferSummaryResp, err error)
	KYTGetTransferAlerts(externalId string) (resp KYTGetTransferAlertsResp, err error)

	KYTRegisterWithdrawalAttempt(userId string, param KYTRegisterWithdrawalAttemptParam) (resp KYTRegisterWithdrawalAttemptResp, err error)
	KYTGetWithdrawalAttemptSummary(externalId string) (resp KYTGetWithdrawalAttemptSummaryResp, err error)
	KYTGetWithdrawalAttemptAlerts(externalId string) (resp KYTGetTransferAlertsResp, err error)

	RetrieveCategories() (resp RetrieveCategoriesResp, err error)
}

type ClientImpl struct {
	host   string
	apiKey string

	client *resty.Client
}

func NewClient(apiKey string, host ...string) *ClientImpl {
	client := resty.New()
	client.SetTimeout(30 * time.Second)
	client.SetHeader("Content-Type", "application/json")
	client.SetHeader("Token", apiKey)
	baseHost := baseUrl
	if len(host) > 0 {
		baseHost = host[0]
	}
	client.SetBaseURL(baseHost)

	return &ClientImpl{
		host:   baseHost,
		apiKey: apiKey,
		client: client,
	}
}

// SetTransport
// Used to set Transport for rate limit
func (c *ClientImpl) SetTransport(transport http.RoundTripper) {
	c.client.SetTransport(transport)
}

func (c *ClientImpl) SetTimeout(timeout time.Duration) {
	c.client.SetTimeout(timeout)
}

func (c *ClientImpl) SetDebug(debug bool) {
	c.client.SetDebug(debug)
}
