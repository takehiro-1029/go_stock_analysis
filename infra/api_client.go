package infra

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type apiClient struct {
	key        string
	httpClient *http.Client
}

func newAPIClient() (*apiClient, error) {

	if os.Getenv("API_KEY") == "" {
		return nil, errors.New("keyを設定してください")
	}

	return &apiClient{os.Getenv("API_KEY"), &http.Client{}}, nil
}

func (api *apiClient) getKey() string {
	return api.key
}

func (api *apiClient) doRequest(baseRawURL, method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseRawURL)
	if err != nil {
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
