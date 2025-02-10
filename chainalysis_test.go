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
	data, err := getClient().EntityAddressRegister("0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestEntityAddressRetrieve(t *testing.T) {
	// RISK : 0xdcbEfFBECcE100cCE9E4b153C4e15cB885643193
	data, err := getClient().EntityAddressRetrieve("0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263")
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
		Network:           "Arbitrum",
		Asset:             "ETH",
		Address:           "0x4675C7e5BaAFBFFbca748158bEcBA61ef3b0a263",
		AttemptIdentifier: "rXfXciGRqh",
		AssetAmount:       100,
		AttemptTimestamp:  time.Now().UTC().Add(-10 * time.Second).Format(iso8601),
	})
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetWithdrawalAttemptSummary(t *testing.T) {
	data, err := getClient().KYTGetWithdrawalAttemptSummary("1cd7bfad-4d97-39b9-a3ae-82b1c344035b")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetWithdrawalAttemptAlerts(t *testing.T) {
	data, err := getClient().KYTGetWithdrawalAttemptAlerts("1cd7bfad-4d97-39b9-a3ae-82b1c344035b")
	assert.Nil(t, err)
	spew.Dump(data)
}

// Category API

func TestRetrieveCategories(t *testing.T) {
	data, err := getClient().RetrieveCategories()
	assert.Nil(t, err)
	t.Log(sonic.MarshalString(data))
}
