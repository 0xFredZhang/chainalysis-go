package chainalysis

import (
	"os"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func getClient() Client {
	return NewClient(os.Getenv("ChainalysisAPIKey"))
}

// Entity API

func TestEntityAddressRegister(t *testing.T) {
	data, err := getClient().EntityAddressRegister("ATte1mfS3F8QMGmhVSgzjM9sPtAYumubX6")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestEntityAddressRetrieve(t *testing.T) {
	data, err := getClient().EntityAddressRetrieve("ATte1mfS3F8QMGmhVSgzjM9sPtAYumubXY6")
	assert.Nil(t, err)
	spew.Dump(data)
}

// KYT API

func TestKYTRegisterTransfer(t *testing.T) {
	data, err := getClient().KYTRegisterTransfer("vddGRBDsjX", KYTRegisterTransferParam{
		Network:           "Ethereum",
		Asset:             "ETH",
		TransferReference: "0x7a1dd6b2b4a162279154cd1edbcc9eff3a6f02dd55bb945ed257ab8519b1cedc",
		Direction:         "received",
	})
	assert.Nil(t, err)
	spew.Dump(data)
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

func TestKYTGetTransferSummary(t *testing.T) {
	data, err := getClient().KYTGetTransferSummary("0x7a1dd6b2b4a162279154cd1edbcc9eff3a6f02dd55bb945ed257ab8519b1cedc")
	assert.Nil(t, err)
	spew.Dump(data)
}

func TestKYTGetWithdrawalAttemptSummary(t *testing.T) {
	data, err := getClient().KYTGetWithdrawalAttemptSummary("0x7a1dd6b2b4a162279154cd1edbcc9eff3a6f02dd55bb945ed257ab8519b1cedc")
	assert.Nil(t, err)
	spew.Dump(data)
}
