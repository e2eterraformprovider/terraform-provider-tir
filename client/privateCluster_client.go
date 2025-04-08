package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/e2eterraformprovider/terraform-provider-tir/models"
)

func (c *Client) NewPrivateCluster(item *models.PrivateCluster, projectID string, teamID string, activeIAM string) (map[string]interface{}, error) {
	integrationJSON, _ := json.Marshal(item)
	buf := bytes.NewBuffer(integrationJSON)

	url := c.Api_endpoint + "/teams/" + teamID + "/projects/" + projectID + "/private-cluster/"
	req, err := http.NewRequest("POST", url, buf)
	if err != nil {
		return nil, err
	}
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
	var jsonRes map[string]interface{}
	err = json.Unmarshal(resBody, &jsonRes)
	if err != nil {
		return nil, err
	}
	return jsonRes, nil
}

func (c *Client) DeletePrivateCluster(privateClusterID string, projectID string, teamID string, activeIAM string) (map[string]interface{}, error) {
	url := c.Api_endpoint + "/teams/" + teamID + "/projects/" + projectID + "/private-cluster/" + privateClusterID + "/"
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
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



func (c *Client) GetPlansPrivateCluster( activeIAM string) (map[string]interface{}, error) {
	urlNode := c.Api_endpoint + "/gpu_service/" + "sku/"
	req, err := http.NewRequest("GET", urlNode, nil)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] CLIENT | NODE READ")
	params := req.URL.Query()
	params.Add("apikey", c.Api_key)
	params.Add("active_iam", activeIAM)
	params.Add("service", "private_cloud")
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