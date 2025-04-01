package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)







func (c *Client) GetProjects(activeIAM string, teamID string) (map[string]interface{}, error) {
	urlNode := c.Api_endpoint + "/teams/" + teamID + "/projects/"
	req, err := http.NewRequest("GET", urlNode, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] CLIENT | NODE READ")
	params := req.URL.Query()
	params.Add("apikey", c.Api_key)
	params.Add("active_iam",activeIAM)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Authorization", "Bearer "+c.Auth_token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "terraform/e2e")
	log.Println(req)
	response, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	log.Println(response)
	if response.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(response.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", response.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", response.StatusCode, respBody.String())
	}
	defer response.Body.Close()
	resBody, _ := io.ReadAll(response.Body)
	stringresponse := string(resBody)
	log.Printf("%s", stringresponse)
	resBytes := []byte(stringresponse)
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBytes, &jsonRes)
	if err != nil {
		log.Printf("[ERROR] CLIENT GET NDE | error when unmarshalling")
		return nil, err
	}
	return jsonRes, nil
}