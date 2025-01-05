package api

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Params struct {
	Service   string
	Version   string
	Action    string
	Region    string
	Payload   string
	SecretId  string
	SecretKey string
}

func Request(params *Params) (string, error) {

	timestamp := time.Now().Unix()
	host := params.Service + ".tencentcloudapi.com"

	authorization := AuthCode(
		params.Service,
		params.Payload,
		params.SecretId,
		params.SecretKey,
		timestamp,
	)

	// send https request

	headers := map[string]string{
		"Authorization":  authorization,
		"Content-Type":   "application/json; charset=utf-8",
		"Host":           host,
		"X-TC-Action":    params.Action,
		"X-TC-Timestamp": strconv.FormatInt(timestamp, 10),
		"X-TC-Version":   params.Version,
	}

	if params.Region != "" {
		headers["X-TC-Region"] = params.Region
	}

	return httpPost("https://"+host, params.Payload, headers)

}

func httpPost(url, msg string, headers map[string]string) (string, error) {

	var err error
	var req *http.Request
	var resp *http.Response
	var body []byte

	rd := strings.NewReader(msg)
	if req, err = http.NewRequest("POST", url, rd); err != nil {
		return "", err
	}

	for key, header := range headers {
		req.Header.Set(key, header)
	}

	client := &http.Client{}
	if resp, err = client.Do(req); err != nil {
		return "", err
	}

	defer resp.Body.Close()

	if body, err = io.ReadAll(resp.Body); err != nil {
		return "", err
	}

	return string(body), err

}
