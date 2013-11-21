package zencoder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type AccountDetails struct {
	AccountState    string `json:"account_state,omitempty"`
	Plan            string `json:"plan,omitempty"`
	MinutesUsed     int32  `json:"minutes_used,omitempty"`
	MinutesIncluded int32  `json:"minutes_included,omitempty"`
	BillingState    string `json:"billing_state,omitempty"`
	IntegrationMode bool   `json:"integration_mode,omitempty"`
}

type CreateAccountRequest struct {
	Email                string  `json:"email,omitempty"`
	TermsOfService       string  `json:"terms_of_service,omitempty"`
	Password             *string `json:"password,omitempty,omitempty"`
	PasswordConfirmation *string `json:"password_confirmation,omitempty,omitempty"`
}

type CreateAccountResponse struct {
	ApiKey   string `json:"api_key,omitempty"`
	Password string `json:"password,omitempty"`
}

// Create an account
func (z *Zencoder) CreateAccount(email, password string) (*CreateAccountResponse, error) {
	request := &CreateAccountRequest{
		Email:          email,
		TermsOfService: "1",
	}

	if len(password) > 0 {
		request.Password = &password
		request.PasswordConfirmation = &password
	}

	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	resp, err := z.call("POST", "account", NewByteReaderCloser(b))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CreateAccountResponse
	err = json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get Account Details
func (z *Zencoder) GetAccount() (*AccountDetails, error) {
	resp, err := z.call("GET", "account", nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var details AccountDetails
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// Set Integration Mode
func (z *Zencoder) SetIntegrationMode() error {
	resp, err := z.call("PUT", "account/integration", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(resp.Status)
	}

	return nil
}

// Set Live Mode
func (z *Zencoder) SetLiveMode() error {
	resp, err := z.call("PUT", "account/live", nil)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(resp.Status)
	}

	return nil
}
