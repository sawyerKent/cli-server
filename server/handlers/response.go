package handlers

import (
	"fmt"
)

type Response interface {
	Display() string
}

type TextResponse struct {
	Message string `json:"message"`
}

func (t TextResponse) Display() string {
	return t.Message
}

type HappyLangResponse struct {
	FRVRID   string `json:"FRVRID"`
	Language string `json:"language"`
}

func (h HappyLangResponse) Display() string {
	return fmt.Sprintf("FRVRID: %s language: %s", h.FRVRID, h.Language)
}

type User struct {
	Name       string `json:"name"`
	MonthOfBDate string `json:"monthofbdate"`
}

type ReturnUser struct {
	Name         string `json:"name"`
	MonthOfBDate  string `json:"monthofbdate"`
	NumericMonth int    `json:"Numericmonth"`
}

type IncomingData map[string][]User

type ReturnData map[string][]ReturnUser
