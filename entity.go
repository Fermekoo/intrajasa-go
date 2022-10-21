package intrajasa

type IntraResult struct {
	MerchantRefCode string  `json:"merchantRefCode"`
	MerchantId      string  `json:"merchantId"`
	VaNumber        string  `json:"vaNumber"`
	Type            string  `json:"type"`
	TotalAmount     float64 `json:"totalAmount"`
	ResponseMsg     string  `json:"responseMsg"`
	ResponseCode    string  `json:"responseCode"`
}

type CustomerData struct {
	CustName     string `json:"custName"`
	CustAddress1 string `json:"custAddress1"`
	// CustPhoneNumber    string `json:"custPhoneNumber"`
	CustEmail          string `json:"custEmail"`
	CustRegisteredDate string `json:"custRegisteredDate"`
	CustCountryCode    string `json:"custCountryCode"`
}

type CreateVa struct {
	MerchantRefCode string        `json:"merchantRefCode"`
	TotalAmount     int           `json:"totalAmount"`
	CustomerData    *CustomerData `json:"customerData"`
	VaType          VaType        `json:"vaType"`
}

type RequestPayload struct {
	MerchantId      string        `json:"merchantId"`
	MerchantRefCode string        `json:"merchantRefCode"`
	CustomerData    *CustomerData `json:"customerData"`
	TotalAmount     string        `json:"totalAmount"`
	VaType          int           `json:"vaType"`
	SecureCode      string        `json:"secureCode"`
}
