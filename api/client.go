package api

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	intrajasa "github.com/Fermekoo/intrajasa-go"
)

type Client struct {
	MerchantId string
	SecretWord string
	Env        intrajasa.EnvType
	BaseUrl    string
}

type Token struct {
	Token        string `json:"Token"`
	ResponseMsg  string `json:"responseMsg"`
	ResponseCode string `json:"responseCode"`
}

func NewClient(merchant_id string, secret_word string, env intrajasa.EnvType) *Client {
	var c Client
	c.MerchantId = merchant_id
	c.SecretWord = secret_word
	c.Env = env
	c.BaseUrl = intrajasa.BaseUrl[env]

	return &c
}

func (c *Client) SecretWordHash() string {
	sha := sha1.New()

	sha.Write([]byte(c.SecretWord))

	encrypted := sha.Sum(nil)

	encrypted_string := fmt.Sprintf("%x", encrypted)

	return encrypted_string
}

func (c *Client) SecureCodeToken(reference_id string) string {
	secret_word_hash := c.SecretWordHash()

	code := c.MerchantId + reference_id + secret_word_hash

	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func (c *Client) SecureCodeVa(reference_id string, amount string, display_name string, token string) string {
	code := c.MerchantId + token + reference_id + display_name + amount
	hash_code := sha256.Sum256([]byte(code))

	encrypted_string := fmt.Sprintf("%x", hash_code[:])

	return encrypted_string
}

func (c *Client) GenerateToken(reference_id string) Token {
	var token Token

	secure_code := c.SecureCodeToken(reference_id)

	payloads := map[string]string{
		"merchantId":      c.MerchantId,
		"merchantRefCode": reference_id,
		"secureCode":      secure_code,
	}

	url := c.BaseUrl + "/vaonline/rest/json/gettoken"

	json_payloads, _ := json.Marshal(payloads)

	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_payloads))

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&token)

	if err != nil {
		log.Fatal(err)
	}

	return token
}

func (c *Client) CreateVa(payloads *intrajasa.CreateVa) (*intrajasa.IntraResult, error) {
	var result intrajasa.IntraResult

	reference_id := payloads.MerchantRefCode
	token := c.GenerateToken(reference_id)
	amount := strconv.FormatFloat(float64(payloads.TotalAmount), 'f', 2, 64)
	secure_code := c.SecureCodeVa(reference_id, amount, payloads.CustomerData.CustName, token.Token)

	body_request := &intrajasa.RequestPayload{
		MerchantId:      c.MerchantId,
		MerchantRefCode: reference_id,
		CustomerData:    payloads.CustomerData,
		TotalAmount:     amount,
		VaType:          int(payloads.VaType),
		SecureCode:      secure_code,
	}

	hash_token := sha256.Sum256([]byte(token.Token))
	encrypted_token := fmt.Sprintf("%x", hash_token[:])

	url := c.BaseUrl + "/vaonline/rest/json/generateva?t=" + encrypted_token

	json_body, err := json.Marshal(body_request)

	if err != nil {
		return &result, err
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(json_body))

	if err != nil {
		return &result, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return &result, err
	}

	fmt.Println(&result)

	return &result, nil
}
