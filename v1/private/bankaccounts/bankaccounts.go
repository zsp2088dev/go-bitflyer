package bankaccounts

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type Request struct{}

type Response []BankAccount

type BankAccount struct {
	ID            int    `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	BankName      string `json:"bank_name"`
	BranchName    string `json:"branch_name"`
	AccountType   string `json:"account_type"`
	AccountNumber string `json:"account_number"`
	AccountName   string `json:"account_name"`
}

func (req *Request) Endpoint() string {
	return "/v1/me/getbankaccounts"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
