package zencoder

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type ReportSettings struct {
	From     *time.Time //: Start date in the format YYYY-MM-DD (default: 30 days ago).
	To       *time.Time //: End date in the format YYYY-MM-DD (default: yesterday).
	Grouping *string    //: Minute usage for only one report grouping (default: none).
}

type VodTotalStatistics struct {
	EncodedMinutes  int32 `json:"encoded_minutes,omitempty"`
	BillableMinutes int32 `json:"billable_minutes,omitempty"`
}

type VodStatistic struct {
	Grouping        string `json:"grouping,omitempty"`
	CollectedOn     string `json:"collected_on,omitempty"`
	EncodedMinutes  int32  `json:"encoded_minutes,omitempty"`
	BillableMinutes int32  `json:"billable_minutes,omitempty"`
}

type VodUsage struct {
	Total      *VodTotalStatistics `json:"total,omitempty"`
	Statistics []*VodStatistic     `json:"statistics,omitempty"`
}

type LiveTotalStatistics struct {
	StreamHours          int32 `json:"stream_hours,omitempty"`
	BillableStreamHours  int32 `json:"billable_stream_hours,omitempty"`
	EncodedHours         int32 `json:"encoded_hours,omitempty"`
	BillableEncodedHours int32 `json:"billable_encoded_hours,omitempty"`
	TotalHours           int32 `json:"total_hours,omitempty"`
	TotalBillableHours   int32 `json:"total_billable_hours,omitempty"`
}

type LiveStatistic struct {
	Grouping             string `json:"grouping,omitempty"`
	CollectedOn          string `json:"collected_on,omitempty"`
	StreamHours          int32  `json:"stream_hours,omitempty"`
	BillableStreamHours  int32  `json:"billable_stream_hours,omitempty"`
	EncodedHours         int32  `json:"encoded_hours,omitempty"`
	BillableEncodedHours int32  `json:"billable_encoded_hours,omitempty"`
	TotalHours           int32  `json:"total_hours,omitempty"`
	TotalBillableHours   int32  `json:"total_billable_hours,omitempty"`
}

type LiveUsage struct {
	Total      *LiveTotalStatistics `json:"total,omitempty"`
	Statistics []*LiveStatistic     `json:"statistics,omitempty"`
}

type CombinedUsage struct {
	Total struct {
		Live LiveTotalStatistics `json:"live,omitempty"`
		Vod  VodTotalStatistics  `json:"vod,omitempty"`
	} `json:"total,omitempty"`
	Statistics struct {
		Live []*LiveStatistic `json:"live,omitempty"`
		Vod  []*VodStatistic  `json:"vod,omitempty"`
	} `json:"statistics,omitempty"`
}

func GetReportQuery(path string, settings *ReportSettings) string {
	if settings != nil {
		query := make(url.Values)
		if settings.From != nil {
			query.Set("from", settings.From.Format("2006-01-02"))
		}
		if settings.To != nil {
			query.Set("to", settings.To.Format("2006-01-02"))
		}
		if settings.Grouping != nil {
			query.Set("grouping", *settings.Grouping)
		}

		if len(query) > 0 {
			return path + "?" + query.Encode()
		}
	}

	return path
}

// Get VOD Usage
func (z *Zencoder) GetVodUsage(settings *ReportSettings) (*VodUsage, error) {
	resp, err := z.call("GET", GetReportQuery("reports/vod", settings), nil)
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

	var details VodUsage
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// Get Live Usage
func (z *Zencoder) GetLiveUsage(settings *ReportSettings) (*LiveUsage, error) {
	resp, err := z.call("GET", GetReportQuery("reports/live", settings), nil)
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

	var details LiveUsage
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func (z *Zencoder) GetUsage(settings *ReportSettings) (*CombinedUsage, error) {
	resp, err := z.call("GET", GetReportQuery("reports/all", settings), nil)
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

	var details CombinedUsage
	err = json.Unmarshal(b, &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}
