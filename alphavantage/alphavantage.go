package alphavantage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go_stock_analysis/message"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const baseURL = "https://www.alphavantage.co/query"

type APIClient struct {
	key        string
	httpClient *http.Client
}

type Ticker struct {
	Open   float32   `json:"open"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
	Close  float32   `json:"close"`
	Volume uint      `json:"volume"`
	Time   time.Time `json:"create_at"`
}

func New(key string) *APIClient {
	apiClient := &APIClient{key, &http.Client{}}
	return apiClient
}

func (api *APIClient) GetKey() string {
	return api.key
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string, data []byte) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
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

func checkInterval(interval string) bool {

	check := []string{"1min", "5min", "15min", "30min", "60min"}

	for i := range check {
		if check[i] == interval {
			return true
		}
	}

	return false
}

func (api *APIClient) GetQuery(series string, symbol string, interval string) (map[string]string, error) {

	query := make(map[string]string)

	if interval != "" {
		if ok := checkInterval(interval); !ok {
			return nil, fmt.Errorf(message.ErrorInterval)
		}
		query["interval"] = interval
	}

	query["function"] = series
	query["symbol"] = symbol
	query["apikey"] = api.GetKey()

	return query, nil
}

func (api *APIClient) GetTicker(series string, symbol string, interval string) ([]Ticker, error) {

	query, err := api.GetQuery(series, symbol, interval)
	if err != nil {
		return nil, err
	}

	bytesBody, err := api.doRequest("GET", "", query, nil)
	if err != nil {
		return nil, err
	}

	var data map[string]map[string]interface{}
	if err = json.Unmarshal(bytesBody, &data); err != nil {
		return nil, err
	}

	tickerSlice := make([]Ticker, 0, 100)
	for k, v := range data {
		if k == "Meta Data" {
			continue
		}
		for at, price := range v {
			name, ok := price.(map[string]interface{})
			if !ok {
				continue
			}

			open := name["1. open"].(string)
			o, err := convertStringToFloat32(open)
			if err != nil {
				continue
			}

			high := name["2. high"].(string)
			h, err := convertStringToFloat32(high)
			if err != nil {
				continue
			}

			low := name["3. low"].(string)
			l, err := convertStringToFloat32(low)
			if err != nil {
				continue
			}

			close := name["4. close"].(string)
			c, err := convertStringToFloat32(close)
			if err != nil {
				continue
			}

			volume := name["5. volume"].(string)
			v, err := convertStringToUint(volume)
			if err != nil {
				continue
			}

			var ticker Ticker
			ticker.Open = *o
			ticker.High = *h
			ticker.Low = *l
			ticker.Close = *c
			ticker.Volume = *v

			t, err := time.Parse("2006-01-02 15:04:05", at)
			if err != nil {
				continue
			}
			ticker.Time = t

			tickerSlice = append(tickerSlice, ticker)
		}
	}

	return tickerSlice, nil
}

func convertStringToFloat32(s string) (*float32, error) {
	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return nil, err
	}
	f32 := float32(f64)
	return &f32, nil
}

func convertStringToUint(s string) (*uint, error) {
	ui64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return nil, err
	}
	ui := uint(ui64)
	return &ui, nil
}
