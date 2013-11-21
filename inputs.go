package zencoder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get Input Details
func (z *Zencoder) GetInputDetails(id int32) (*InputMediaFile, error) {
	resp, err := z.call("GET", fmt.Sprintf("inputs/%d.json", id), nil)
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

	var details InputMediaFile
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// Input Progress
func (z *Zencoder) GetInputProgress(id int32) (*FileProgress, error) {
	resp, err := z.call("GET", fmt.Sprintf("inputs/%d/progress.json", id), nil)
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
