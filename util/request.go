package util

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"bitbucket.org/tixid/tix_id_cms_api/config"
)

var httpClient *http.Client

var unexpectedHttpCode = errors.New("unexpected http code")

func DoFormPost(logId string, url string, values *url.Values, expCode int, rs interface{}) error {
	httpClient = &http.Client{
		Timeout: time.Millisecond * time.Duration(config.MustGetInt("server.http_timeout")),
	}

	vs := values.Encode()
	log.DebugfWithId(logId, "url: %s, values: %s", url, vs)

	req, err := http.NewRequest("POST", url, strings.NewReader(vs))
	if err != nil {
		log.ErrorfWithId(logId, "error creating request: %s", err.Error())
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.ErrorfWithId(logId, "error executing request: %s", err.Error())
		return err
	}
	if resp.StatusCode != expCode {
		log.ErrorfWithId(logId, "expecting http status %d, got %d", expCode, resp.StatusCode)
		return unexpectedHttpCode
	}

	defer resp.Body.Close()

	// skip reading the response if rs is nil
	if rs == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.ErrorfWithId(logId, "error reading body: %s", err.Error())
		return err
	}
	log.DebugfWithId(logId, "got body: %s", string(body))

	err = json.Unmarshal(body, rs)
	if err != nil {
		log.Error(logId, "error parsing response json: %s for body %+v", err.Error(), body)
		return err
	}
	return nil

}

//func DoGet(logId string, url string, values *url.Values, expCode int, header map[string]string, rs interface{}) error {
func DoGet(logId string, url string, values *url.Values, expCode int, rs interface{}) error {
	vs := values.Encode()
	log.DebugfWithId(logId, "url: %s, values: %s", url, vs)

	httpClient = &http.Client{
		Timeout: time.Millisecond * time.Duration(config.MustGetInt("server.http_timeout")),
	}

	req, err := http.NewRequest("GET", url+"?"+vs, nil)
	if err != nil {
		log.ErrorfWithId(logId, "error creating request: %s", err.Error())
		return err
	}
	//for key, value := range header {
	//	req.Header.Set(key, value)
	//}
	log.Debug(req.URL)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.ErrorfWithId(logId, "error executing request: %s", err.Error())
		return err
	}
	if resp.StatusCode != expCode {
		log.ErrorfWithId(logId, "expecting http status %d, got %d", expCode, resp.StatusCode)
		return unexpectedHttpCode
	}

	defer resp.Body.Close()

	// skip reading the response if rs is nil
	if rs == nil {
		return nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.ErrorfWithId(logId, "error reading body: %s", err.Error())
		return err
	}

	err = json.Unmarshal(body, rs)
	if err != nil {
		log.Error(logId, "error parsing response json: %s for body %+v", err.Error(), body)
		return err
	}
	return nil
}

func DoPost(logId string, url string, payload string, expCode int, rs interface{}) (string, error) {

	httpClient = &http.Client{
		Timeout: time.Millisecond * time.Duration(config.MustGetInt("server.http_timeout")),
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		log.ErrorfWithId(logId, "error creating request: %s", err.Error())
		return "", err
	}
	//for key, value := range header {
	//	req.Header.Set(key, value)
	//}
	log.Debug(req.URL)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.ErrorfWithId(logId, "error executing request: %s", err.Error())
		return "", err
	}
	if resp.StatusCode != expCode {
		log.ErrorfWithId(logId, "expecting http status %d, got %d", expCode, resp.StatusCode)
		return "", unexpectedHttpCode
	}

	defer resp.Body.Close()

	// skip reading the response if rs is nil
	if rs == nil {
		return "", nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.ErrorfWithId(logId, "error reading body: %s", err.Error())
		return "", err
	}

	err = json.Unmarshal(body, rs)
	if err != nil {
		log.Error(logId, "error parsing response json: %s for body %+v", err.Error(), body)
		return "", err
	}
	return string(body), nil

}
