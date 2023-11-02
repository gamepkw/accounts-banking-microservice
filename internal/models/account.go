package model

import (
	"time"
)

type Account struct {
	AccountNo string     `json:"account_no,omitempty"`
	Uuid      string     `json:"uuid,omitempty"`
	Name      string     `json:"name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Tel       string     `json:"tel,omitempty"`
	Balance   float64    `json:"balance"`
	Bank      string     `json:"bank,omitempty"`
	Status    string     `json:"status,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	IsClosed  int        `json:"is_closed,omitempty"`
}

type CountAccount struct {
	Status string `json:"status,"`
	Count  int    `json:"count"`
}
