package security

import (
	"encoding/json"
	"testing"

	"github.com/emitter-io/emitter/network/http"
	"github.com/stretchr/testify/assert"
)

func TestNewSingleContractProvider(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewSingleContractProvider(license)

	assert.EqualValues(t, p.owner.MasterID, 1)
	assert.EqualValues(t, p.owner.Signature, license.Signature)
	assert.EqualValues(t, p.owner.ID, license.Contract)
}

func TestSingleContractProvider_Create(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewSingleContractProvider(license)
	contract, err := p.Create()

	assert.Nil(t, contract)
	assert.Error(t, err)
}

func TestSingleContractProvider_Get(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewSingleContractProvider(license)
	contractByID := p.Get(license.Contract)
	contractByWrongID := p.Get(0)

	assert.NotNil(t, contractByID)
	assert.Nil(t, contractByWrongID)
}

func TestSingleContractProvider_Validate(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewSingleContractProvider(license)
	contract := p.Get(license.Contract)

	key := Key(make([]byte, 24))
	key.SetMaster(1)
	key.SetContract(license.Contract)
	key.SetSignature(license.Signature)

	assert.True(t, contract.Validate(key))
}

func TestNewHTTPContractProvider(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewHTTPContractProvider(license)

	assert.EqualValues(t, p.owner.MasterID, 1)
	assert.EqualValues(t, p.owner.Signature, license.Signature)
	assert.EqualValues(t, p.owner.ID, license.Contract)
}

func TestHTTPContractProvider_Create(t *testing.T) {
	license, err := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewHTTPContractProvider(license)
	contract, err := p.Create()

	assert.Nil(t, contract)
	assert.Error(t, err)
}

func TestHTTPContractProvider_Get(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")

	p := NewHTTPContractProvider(license)
	oldGet := http.Get
	defer func() {
		http.Get = oldGet
	}()

	http.Get = mockGet
	contractByID := p.Get(1)
	contractByWrongID := p.Get(0)
	http.Get = oldGet
	assert.NotNil(t, contractByID)
	assert.Nil(t, contractByWrongID)
}

func TestHTTPContractProvider_Validate(t *testing.T) {
	license, _ := ParseLicense("zT83oDV0DWY5_JysbSTPTDr8KB0AAAAAAAAAAAAAAAI")
	p := NewSingleContractProvider(license)
	contract := p.Get(license.Contract)

	key := Key(make([]byte, 24))
	key.SetMaster(1)
	key.SetContract(license.Contract)
	key.SetSignature(license.Signature)

	assert.True(t, contract.Validate(key))
}

func mockGet(url string, output interface{}, headers ...http.HeaderValue) error {
	if url == "http://meta.emitter.io/v1/contract/1" {
		return json.Unmarshal([]byte(`{"id": 1}`), output)
	}
	return nil
}
