package zencoder_test

import (
	zencoder "."
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateJob(t *testing.T) {
	expectedStatus := http.StatusCreated
	returnBody := true

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatus)
		if expectedStatus != http.StatusCreated {
			return
		}

		if !returnBody {
			return
		}

		fmt.Fprintln(w, `{
  "id": "1234",
  "outputs": [
    {
      "id": "4321"
    }
  ]
}`)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	var settings zencoder.EncodingSettings
	resp, err := zc.CreateJob(&settings)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	if resp == nil {
		t.Fatal("Expected a response")
	}

	if resp.Id != "1234" {
		t.Fatal("Expected Id=1234", resp.Id)
	}

	if len(resp.Outputs) != 1 {
		t.Fatal("Expected one output", len(resp.Outputs))
	}

	if resp.Outputs[0].Id != "4321" {
		t.Fatal("Expected Id=4321", resp.Outputs[0].Id)
	}

	expectedStatus = http.StatusInternalServerError
	resp, err = zc.CreateJob(&settings)
	if err == nil {
		t.Fatal("Expected error")
	}

	if resp != nil {
		t.Fatal("Expected no response")
	}

	returnBody = false
	expectedStatus = http.StatusOK

	resp, err = zc.CreateJob(&settings)
	if err == nil {
		t.Fatal("Expected error")
	}

	if resp != nil {
		t.Fatal("Expected no response")
	}

	returnBody = true
	srv.Close()

	resp, err = zc.CreateJob(&settings)
	if err == nil {
		t.Fatal("Expected error")
	}

	if resp != nil {
		t.Fatal("Expected no response")
	}
}

func TestListJobs(t *testing.T) {
	expectedStatus := http.StatusOK
	returnBody := true

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs.json", func(w http.ResponseWriter, r *http.Request) {
		if expectedStatus != http.StatusOK {
			w.WriteHeader(expectedStatus)
			return
		}

		if !returnBody {
			return
		}

		fmt.Fprintln(w, `[{
  "job": {
    "created_at": "2010-01-01T00:00:00Z",
    "finished_at": "2010-01-01T00:00:00Z",
    "updated_at": "2010-01-01T00:00:00Z",
    "submitted_at": "2010-01-01T00:00:00Z",
    "pass_through": null,
    "id": 1,
    "input_media_file": {
      "format": "mpeg4",
      "created_at": "2010-01-01T00:00:00Z",
      "frame_rate": 29,
      "finished_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "duration_in_ms": 24883,
      "audio_sample_rate": 48000,
      "url": "s3://bucket/test.mp4",
      "id": 1,
      "error_message": null,
      "error_class": null,
      "audio_bitrate_in_kbps": 95,
      "audio_codec": "aac",
      "height": 352,
      "file_size_bytes": 1862748,
      "video_codec": "h264",
      "test": false,
      "total_bitrate_in_kbps": 593,
      "channels": "2",
      "width": 624,
      "video_bitrate_in_kbps": 498,
      "state": "finished"
    },
    "test": false,
    "output_media_files": [{
      "format": "mpeg4",
      "created_at": "2010-01-01T00:00:00Z",
      "frame_rate": 29,
      "finished_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "duration_in_ms": 24883,
      "audio_sample_rate": 44100,
      "url": "http://s3.amazonaws.com/bucket/video.mp4",
      "id": 1,
      "error_message": null,
      "error_class": null,
      "audio_bitrate_in_kbps": 92,
      "audio_codec": "aac",
      "height": 352,
      "file_size_bytes": 1386663,
      "video_codec": "h264",
      "test": false,
      "total_bitrate_in_kbps": 443,
      "channels": "2",
      "width": 624,
      "video_bitrate_in_kbps": 351,
      "state": "finished",
      "label": "Web"
    }],
    "thumbnails": [{
      "created_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "url": "http://s3.amazonaws.com/bucket/video/frame_0000.png",
      "id": 1
    }],
    "state": "finished"
  }
}]`)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	jobs, err := zc.ListJobs()
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	if jobs == nil {
		t.Fatal("Expected jobs")
	}

	if len(jobs) != 1 {
		t.Fatal("Expected 1 job", len(jobs))
	}

	expectedStatus = http.StatusInternalServerError
	jobs, err = zc.ListJobs()
	if err == nil {
		t.Fatal("Expected error")
	}

	if jobs != nil {
		t.Fatal("Expected no response")
	}

	expectedStatus = http.StatusOK
	returnBody = false
	jobs, err = zc.ListJobs()
	if err == nil {
		t.Fatal("Expected error")
	}

	if jobs != nil {
		t.Fatal("Expected no response")
	}

	srv.Close()
	returnBody = true
	jobs, err = zc.ListJobs()
	if err == nil {
		t.Fatal("Expected error")
	}

	if jobs != nil {
		t.Fatal("Expected no response")
	}
}

func TestGetJobDetails(t *testing.T) {
	expectedStatus := http.StatusOK
	returnBody := true

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs/123.json", func(w http.ResponseWriter, r *http.Request) {
		if expectedStatus != http.StatusOK {
			w.WriteHeader(expectedStatus)
			return
		}

		if !returnBody {
			return
		}

		fmt.Fprintln(w, `{
  "job": {
    "created_at": "2010-01-01T00:00:00Z",
    "finished_at": "2010-01-01T00:00:00Z",
    "updated_at": "2010-01-01T00:00:00Z",
    "submitted_at": "2010-01-01T00:00:00Z",
    "pass_through": null,
    "id": 1,
    "input_media_file": {
      "format": "mpeg4",
      "created_at": "2010-01-01T00:00:00Z",
      "frame_rate": 29,
      "finished_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "duration_in_ms": 24883,
      "audio_sample_rate": 48000,
      "url": "s3://bucket/test.mp4",
      "id": 1,
      "error_message": null,
      "error_class": null,
      "audio_bitrate_in_kbps": 95,
      "audio_codec": "aac",
      "height": 352,
      "file_size_bytes": 1862748,
      "video_codec": "h264",
      "test": false,
      "total_bitrate_in_kbps": 593,
      "channels": "2",
      "width": 624,
      "video_bitrate_in_kbps": 498,
      "state": "finished",
      "md5_checksum":"7f106918e02a69466afa0ee014174143"
    },
    "test": false,
    "output_media_files": [{
      "format": "mpeg4",
      "created_at": "2010-01-01T00:00:00Z",
      "frame_rate": 29,
      "finished_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "duration_in_ms": 24883,
      "audio_sample_rate": 44100,
      "url": "http://s3.amazonaws.com/bucket/video.mp4",
      "id": 1,
      "error_message": null,
      "error_class": null,
      "audio_bitrate_in_kbps": 92,
      "audio_codec": "aac",
      "height": 352,
      "file_size_bytes": 1386663,
      "video_codec": "h264",
      "test": false,
      "total_bitrate_in_kbps": 443,
      "channels": "2",
      "width": 624,
      "video_bitrate_in_kbps": 351,
      "state": "finished",
      "label": "Web",
      "md5_checksum":"7f106918e02a69466afa0ee014172496"
    }],
    "thumbnails": [{
      "created_at": "2010-01-01T00:00:00Z",
      "updated_at": "2010-01-01T00:00:00Z",
      "url": "http://s3.amazonaws.com/bucket/video/frame_0000.png",
      "id": 1
    }],
    "state": "finished"
  }
}`)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	details, err := zc.GetJobDetails(123)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	if details == nil {
		t.Fatal("Expected details")
	}

	if details.Job.State != "finished" {
		t.Fatal("Expected state to be", details.Job.State)
	}

	if len(details.Job.Thumbnails) != 1 {
		t.Fatal("Expected thumbnails to be", len(details.Job.Thumbnails))
	}

	expectedStatus = http.StatusNotFound
	details, err = zc.GetJobDetails(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if details != nil {
		t.Fatal("Expected no response")
	}

	expectedStatus = http.StatusOK
	returnBody = false
	details, err = zc.GetJobDetails(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if details != nil {
		t.Fatal("Expected no response")
	}

	returnBody = false
	srv.Close()
	details, err = zc.GetJobDetails(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if details != nil {
		t.Fatal("Expected no response")
	}
}

func TestResubmitJob(t *testing.T) {
	expectedStatus := http.StatusNoContent

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs/123/resubmit.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatus)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	err := zc.ResubmitJob(123)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	expectedStatus = http.StatusConflict

	err = zc.ResubmitJob(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	expectedStatus = http.StatusOK
	srv.Close()

	err = zc.ResubmitJob(123)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestCancelJob(t *testing.T) {
	expectedStatus := http.StatusNoContent

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs/123/cancel.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatus)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	err := zc.CancelJob(123)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	expectedStatus = http.StatusConflict

	err = zc.CancelJob(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	expectedStatus = http.StatusOK
	srv.Close()

	err = zc.CancelJob(123)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestFinishLiveJob(t *testing.T) {
	expectedStatus := http.StatusNoContent

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs/123/finish", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatus)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	err := zc.FinishLiveJob(123)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	expectedStatus = http.StatusConflict

	err = zc.FinishLiveJob(123)
	if err == nil {
		t.Fatal("Expected error")
	}
}

func TestGetJobProgress(t *testing.T) {
	expectedStatus := http.StatusOK
	returnBody := true

	mux := http.NewServeMux()
	mux.HandleFunc("/jobs/123/progress.json", func(w http.ResponseWriter, r *http.Request) {
		if expectedStatus != http.StatusOK {
			w.WriteHeader(expectedStatus)
			return
		}

		if !returnBody {
			return
		}

		fmt.Fprintln(w, `{
  "state": "processing",
  "progress": 32.34567345,
  "input": {
    "id": 1234,
    "state": "finished"
  },
  "outputs": [
    {
      "id": 4567,
      "state": "processing",
      "current_event": "Transcoding",
      "current_event_progress": 25.0323,
      "progress": 35.23532
    },
    {
      "id": 4568,
      "state": "processing",
      "current_event": "Uploading",
      "current_event_progress": 82.32,
      "progress": 95.3223
    }
  ]
}`)
	})

	srv := httptest.NewServer(mux)

	zc := zencoder.NewZencoder("abc")
	zc.BaseUrl = srv.URL

	progress, err := zc.GetJobProgress(123)
	if err != nil {
		t.Fatal("Expected no error", err)
	}

	if progress == nil {
		t.Fatal("Expected response")
	}

	expectedStatus = http.StatusNotFound
	progress, err = zc.GetJobProgress(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if progress != nil {
		t.Fatal("Expected no response", progress)
	}

	expectedStatus = http.StatusOK
	returnBody = false
	progress, err = zc.GetJobProgress(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if progress != nil {
		t.Fatal("Expected no response", progress)
	}

	returnBody = true
	srv.Close()
	progress, err = zc.GetJobProgress(123)
	if err == nil {
		t.Fatal("Expected error")
	}

	if progress != nil {
		t.Fatal("Expected no response", progress)
	}
}
