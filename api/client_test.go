package api

import (
	"testing"

	intrajasa "github.com/Fermekoo/intrajasa-go"

	"github.com/google/uuid"
	assert "github.com/stretchr/testify/require"
)

func TestGenerateVa(t *testing.T) {

	intra_client := NewClient("put your merchant id", "put your secret word", intrajasa.Sandbox)
	merchantRefcode := uuid.NewString()
	payloads := &intrajasa.CreateVa{
		MerchantRefCode: merchantRefcode,
		TotalAmount:     10000,
		VaType:          intrajasa.OneTime,
		CustomerData: &intrajasa.CustomerData{
			CustName:           "DANDI TEST",
			CustAddress1:       "Jakarta",
			CustEmail:          "dandifermeko@gmail.com",
			CustRegisteredDate: "2022-10-11",
			CustCountryCode:    "021",
		},
	}
	va, err := intra_client.CreateVa(payloads)
	assert.NoError(t, err)
	assert.Equal(t, "200", va.ResponseCode)
}

func TestGenerateVaInvalidToken(t *testing.T) {

	intra_client := NewClient("", "", intrajasa.Sandbox)
	merchantRefcode := uuid.NewString()
	payloads := &intrajasa.CreateVa{
		MerchantRefCode: merchantRefcode,
		TotalAmount:     10000,
		VaType:          intrajasa.OneTime,
		CustomerData: &intrajasa.CustomerData{
			CustName:           "DANDI TEST",
			CustAddress1:       "Jakarta",
			CustEmail:          "dandifermeko@gmail.com",
			CustRegisteredDate: "2022-10-11",
			CustCountryCode:    "021",
		},
	}
	va, err := intra_client.CreateVa(payloads)
	assert.NoError(t, err)
	assert.Equal(t, va.ResponseCode, "221")
}

func TestGenerateToken(t *testing.T) {
	intra_client := NewClient("", "", intrajasa.Sandbox)
	ref_code := uuid.NewString()

	token := intra_client.GenerateToken(ref_code)
	assert.NotEmpty(t, token)
	assert.Equal(t, "200", token.ResponseCode)
}
