package test

import (
	"testing"

	"github.com/sawyerKent/cli-server/server/models"
	"github.com/stretchr/testify/assert"
)

func TestHappyLangResponse_Display(t *testing.T) {
	response := models.HappyLangResponse{
		FRVRID:   "12345",
		Language: "English",
	}

	assert.Equal(t, "FRVRID: 12345 language: English", response.Display())
}
