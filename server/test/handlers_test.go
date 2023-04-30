package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sawyerKent/cli-server/server/handlers"
	"github.com/sawyerKent/cli-server/server/models"
	"github.com/stretchr/testify/assert"
)

func TestGetEndpoint(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"Message": "Hello, test!"}`))
	}))
	defer ts.Close()

	res, err := handlers.GetEndpoint(ts.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Hello, test!", res.Display())
}

func TestPostEndpoint(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		frvrID := r.FormValue("FRVRID")
		language := r.FormValue("language")

		if frvrID == "12345" && language == "English" {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"Message": "Posted successfully!"}`))
		} else {
			http.Error(w, "Invalid data received", http.StatusBadRequest)
		}
	}))
	defer ts.Close()

	data := models.HappyLangResponse{
		FRVRID:   "12345",
		Language: "English",
	}

	res, err := handlers.PostEndpoint(ts.URL, data)

	assert.NoError(t, err)
	assert.Equal(t, "Posted successfully!", res.Display())
}

func TestPostJsonEndpoint(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "FRVRID: 12345 language: English"}`))
	}))
	defer ts.Close()

	data := models.HappyLangResponse{
		FRVRID:   "12345",
		Language: "English",
	}

	res, err := handlers.PostJsonEndpoint(ts.URL, data)

	assert.NoError(t, err)
	assert.Equal(t, "FRVRID: 12345 language: English", res.Display())
}

func TestSendJson(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"OK","message":"JSON received!"}`))
	}))
	defer ts.Close()

	data := models.IncomingData{
		"Group1": []models.User{
			{Name: "Alice", MonthOfBDate: "January"},
			{Name: "Bob", MonthOfBDate: "February"},
		},
		"Group2": []models.User{
			{Name: "Charlie", MonthOfBDate: "March"},
			{Name: "David", MonthOfBDate: "April"},
		},
	}

	res, err := handlers.SendJson(ts.URL, data)

	assert.NoError(t, err)

	expectedResponse := []byte(`{"status":"OK","message":"JSON received!"}`)
	assert.Equal(t, expectedResponse, res)
}
