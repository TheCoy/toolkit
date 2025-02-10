package thirdpart

import (
	"encoding/base64"
	"encoding/json"
	"github.com/TheCoy/nextcaptcha-go"
	"github.com/pkg/errors"
	"sync"
)

var (
	CipherClientKey string

	initOnce     sync.Once
	reCaptchaAPI *nextcaptcha.NextCaptchaAPI
)

func GenRecaptchaToken(host, websiteKey string) (interface{}, error) {
	initOnce.Do(func() {
		plainBytes, _ := base64.StdEncoding.DecodeString(CipherClientKey)
		clientKey := string(plainBytes)
		reCaptchaAPI = nextcaptcha.NewNextCaptchaAPI(clientKey, "", "", false)
	})
	if reCaptchaAPI == nil {
		return nil, errors.New("reCaptchaAPI not initialized")
	}
	result, err := reCaptchaAPI.RecaptchaV3(host, websiteKey, nextcaptcha.RecaptchaV3Options{
		Type:       nextcaptcha.RECAPTCHAV3_HS_PROXYLESS_TYPE,
		PageAction: "SIGN_IN",
		ApiDomain:  "www.recaptcha.net",
	})
	if err != nil {
		return nil, err
	}
	bytes, _ := json.Marshal(result)
	var response struct {
		ErrorId  int         `json:"errorId"`
		Status   string      `json:"status"`
		TaskId   interface{} `json:"taskId"`
		Solution struct {
			GRecaptchaResponse string `json:"gRecaptchaResponse"`
		}
	}
	_ = json.Unmarshal(bytes, &response)
	return response.Solution.GRecaptchaResponse, nil
}
