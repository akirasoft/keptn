// This file is safe to edit. Once it exists it will not be overwritten

package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/keptn/go-utils/pkg/utils"
)

const eventbrokerURL = "http://event-broker.keptn.svc.cluster.local/keptn"

// PostToEventBroker makes a post request to the eventbroker
func PostToEventBroker(e interface{}, shkeptncontext string) error {

	data, err := json.Marshal(e)
	if err != nil {
		utils.Error(shkeptncontext, fmt.Sprintf("Error marshaling data %s", err.Error()))
		return err
	}

	req, err := http.NewRequest("POST", eventbrokerURL, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/cloudevents+json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		utils.Error(shkeptncontext, fmt.Sprintf("Error making POST %s", err.Error()))
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		return nil
	}
	return errors.New(resp.Status)
}
