package models

import (
	"encoding/json"
	"fmt"
)

// INTERFACES
type JsonData interface {
    MarshalJSON() ([]byte, error)
}

type Response interface {
	Display() string
}

type UserData interface {
    GetName() string
    GetMonthOfBDate() string
}

type DataContainer interface {
    GetAllGroups() []string
    GetUsersInGroup(group string) ([]UserData, bool)
    AddUserToGroup(group string, user UserData)
}

// STRUCTS and MAPS

// TextResponse
type TextResponse struct {
	Message string `json:"message"`
}

func (t TextResponse) Display() string {
	return t.Message
}


// HappyLangResponse
type HappyLangResponse struct {
	FRVRID   string `json:"FRVRID"`
	Language string `json:"language"`
}

func (h HappyLangResponse) Display() string {
	return fmt.Sprintf("FRVRID: %s language: %s", h.FRVRID, h.Language)
}


// User
type User struct {
	Name       string `json:"name"`
	MonthOfBDate string `json:"monthofbdate"`
}

func (u User) GetName() string {
    return u.Name
}

func (u User) GetMonthOfBDate() string {
    return u.MonthOfBDate
}


// ReturnUser
type ReturnUser struct {
	Name         string `json:"name"`
	MonthOfBDate  string `json:"monthofbdate"`
	NumericMonth int    `json:"Numericmonth"`
}

func (ru ReturnUser) GetName() string {
    return ru.Name
}

func (ru ReturnUser) GetMonthOfBDate() string {
    return ru.MonthOfBDate
}


// IncomingData
type IncomingData map[string][]User

func (id IncomingData) GetAllGroups() []string {
    groups := make([]string, 0, len(id))
    for group := range id {
        groups = append(groups, group)
    }
    return groups
}

func (id IncomingData) GetUsersInGroup(group string) ([]UserData, bool) {
    users, ok := id[group]
    if !ok {
        return nil, false
    }

    userData := make([]UserData, len(users))
    for i, user := range users {
        userData[i] = user
    }
    return userData, true
}

func (id IncomingData) AddUserToGroup(group string, user UserData) {
    id[group] = append(id[group], user.(User))
}

func (id IncomingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string][]User(id))
}


// ReturnData
type ReturnData map[string][]ReturnUser

func (rd ReturnData) GetAllGroups() []string {
    groups := make([]string, 0, len(rd))
    for group := range rd {
        groups = append(groups, group)
    }
    return groups
}

func (rd ReturnData) GetUsersInGroup(group string) ([]UserData, bool) {
    users, ok := rd[group]
    if !ok {
        return nil, false
    }

    userData := make([]UserData, len(users))
    for i, user := range users {
        userData[i] = user
    }
    return userData, true
}

func (rd ReturnData) AddUserToGroup(group string, user UserData) {
    rd[group] = append(rd[group], user.(ReturnUser))
}

