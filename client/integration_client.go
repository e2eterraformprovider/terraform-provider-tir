package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"terraform-provider-tir/models"
)

func (c *Client) NewIntegration(item *models.Integration, projectID string, teamID string, activeIAM string) (map[string]interface{}, error) {
	integrationJSON, _ := json.Marshal(item)
	buf := bytes.NewBuffer(integrationJSON)

	url := c.Api_endpoint + "/teams/" + teamID + "/projects/" + projectID + "/integrations/"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
	log.Println("jatin111111")
	params := req.URL.Query()
	params.Add("apikey", c.Api_key)
	params.Add("active_iam", activeIAM)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Authorization", "Bearer "+c.Auth_token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "terraform/e2e")
	response, err := c.HttpClient.Do(req)
	if err != nil {
		log.Println("Error while creating")
		return nil, err
	}
	err = CheckResponseStatus(response)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	resBody, _ := io.ReadAll(response.Body)
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBody, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func (c *Client) DeleteIntegration(integrationID string, projectID string, teamID string, activeIAM string) (map[string]interface{}, error) {
	url := c.Api_endpoint + "/teams/" + teamID + "/projects/" + projectID + "/integrations/" + integrationID + "/"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	log.Println("jatin111111DELETE")
	params := req.URL.Query()
	params.Add("apikey", c.Api_key)
	params.Add("active_iam", activeIAM)
	req.URL.RawQuery = params.Encode()
	req.Header.Add("Authorization", "Bearer "+c.Auth_token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("User-Agent", "terraform/e2e")
	response, err := c.HttpClient.Do(req)
	if err != nil {
		log.Println("Error while Reading")
		return nil, err
	}
	defer response.Body.Close()
	return nil, nil
}
