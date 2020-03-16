package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	api "github.com/zsp2088dev/go-bitflyer"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type Auth struct {
	APIKey    string
	APISecret string
}

func CreateHeaders(auth *Auth, req api.Request, u *url.URL) *http.Header {
	timestamp := fmt.Sprintf("%d", time.Now().Unix())

	var path string
	if req.Query() != "" {
		path = u.Path + "?" + u.RawQuery
	} else {
		path = u.Path
	}

	h := hmac.New(sha256.New, []byte(auth.APISecret))
	h.Write([]byte(timestamp))
	h.Write([]byte(req.Method()))
	h.Write([]byte(path))
	if req.Payload() != nil {
		h.Write(req.Payload())
	}

	sign := hex.EncodeToString(h.Sum(nil))

	header := http.Header{}
	header.Set("Content-Type", "application/json")
	header.Set("ACCESS-KEY", auth.APIKey)
	header.Set("ACCESS-TIMESTAMP", timestamp)
	header.Set("ACCESS-SIGN", sign)

	return &header
}

type Params struct {
	APIKey    string `json:"api_key"`
	Timestamp int    `json:"timestamp"`
	Nonce     string `json:"nonce"`
	Signature string `json:"signature"`
}

func CreateAuthParams(auth *Auth) Params {
	now := time.Now().Unix()
	rand.Seed(now)

	timestamp := int(now)
	nonce := fmt.Sprintf("%d", rand.Int())

	h := hmac.New(sha256.New, []byte(auth.APISecret))
	h.Write([]byte(fmt.Sprintf("%d%s", timestamp, nonce)))

	sign := hex.EncodeToString(h.Sum(nil))

	return Params{
		APIKey:    auth.APIKey,
		Timestamp: timestamp,
		Nonce:     nonce,
		Signature: sign,
	}

}
