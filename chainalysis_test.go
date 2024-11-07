package chainalysis

import (
	"os"
	"testing"
	"time"

	"github.com/bytedance/sonic"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func getClient() Client {
	cli := NewClient(os.Getenv("ChainalysisAPIKey"))
	cli.SetDebug(true)
	return cli
}

// Entity API

func TestEntityAddressRegister(t *testing.T) {
	data, err := getClient().EntityAddressRegister("0xdcbEfFBECcE100cCE9E4b153C4e15cB885643193")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestEntityAddressRetrieve(t *testing.T) {
	// RISK : 0xdcbEfFBECcE100cCE9E4b153C4e15cB885643193
	data, err := getClient().EntityAddressRetrieve("0xdcbEfFBECcE100cCE9E4b153C4e15cB885643193")
	assert.Nil(t, err)
	t.Log(sonic.MarshalString(data))
}

// KYT API

func TestKYTRegisterTransfer(t *testing.T) {
	data, err := getClient().KYTRegisterTransfer("vddGRBDsjX", KYTRegisterTransferParam{
		Network:           "Ethereum",
		Asset:             "ETH",
		TransferReference: "0x8e73f0c0d8bdeab7e9f85e4fc84e0e3d3e956e3e15584fc28222e9126d9ce716",
		Direction:         "received",
	})
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetTransferSummary(t *testing.T) {
	// 70e7170a-e7fe-3ea2-a033-2d5b51da4075 normal
	// fadea2aa-8751-36e0-af84-b07289a2e443 risk
	data, err := getClient().KYTGetTransferSummary("fadea2aa-8751-36e0-af84-b07289a2e443")
	assert.Nil(t, err)
	t.Log(sonic.MarshalString(data))
}

func TestKYTGetTransferAlerts(t *testing.T) {
	data, err := getClient().KYTGetTransferAlerts("fadea2aa-8751-36e0-af84-b07289a2e443")
	assert.Nil(t, err)
	t.Log(sonic.MarshalString(data))
}

func TestKYTRegisterWithdrawalAttempt(t *testing.T) {
	data, err := getClient().KYTRegisterWithdrawalAttempt("vddGRBDsjX", KYTRegisterWithdrawalAttemptParam{
		Network:           "Ethereum",
		Asset:             "ETH",
		Address:           "0x8589427373D6D84E98730D7795D8f6f8731FDA16",
		AttemptIdentifier: "",
		AssetAmount:       100,
		AttemptTimestamp:  time.Now().UTC().Add(-10 * time.Second).Format(iso8601),
	})
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetWithdrawalAttemptSummary(t *testing.T) {
	data, err := getClient().KYTGetWithdrawalAttemptSummary("")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetWithdrawalAttemptAlerts(t *testing.T) {
	data, err := getClient().KYTGetWithdrawalAttemptAlerts("")
	assert.Nil(t, err)
	spew.Dump(data)
}

// Category API

func TestRetrieveCategories(t *testing.T) {
	data, err := getClient().RetrieveCategories()
	assert.Nil(t, err)
	t.Log(sonic.MarshalString(data))
}
