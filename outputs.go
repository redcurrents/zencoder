package zencoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get Output Details
func (z *Zencoder) GetOutputDetails(id int32) (*OutputMediaFile, error) {
	resp, err := z.call("GET", fmt.Sprintf("outputs/%d.json", id), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var details OutputMediaFile
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// Output Progress
func (z *Zencoder) GetOutputProgress(id int32) (*FileProgress, error) {
	resp, err := z.call("GET", fmt.Sprintf("outputs/%d/progress.json", id), nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var details FileProgress
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}
