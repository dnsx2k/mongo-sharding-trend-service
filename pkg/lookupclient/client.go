package lookupclient

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type ClientCtx struct {
	httpClient *http.Client
	baseURL    string
}

func New(baseURL string) *ClientCtx {
	return &ClientCtx{httpClient: http.DefaultClient, baseURL: baseURL}
}

func (cc *ClientCtx) SendLookupEntries(location string, IDs []string) error {
	payload := map[string]interface{}{"location": location, "keys": IDs}
	bb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, cc.baseURL, bytes.NewReader(bb))
	if err != nil {
		return err
	}
	_, err = cc.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}

func (cc *ClientCtx) DeleteLookupEntries(location string, IDs []string) error {
	payload := map[string]interface{}{"location": location, "keys": IDs}
	bb, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, cc.baseURL, bytes.NewReader(bb))
	if err != nil {
		return err
	}
	_, err = cc.httpClient.Do(req)
	if err != nil {
		return err
	}

	return nil
}
