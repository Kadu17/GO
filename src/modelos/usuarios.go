package modelos

import "time"

type Usuario struct {
	ID       uint64 `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Senha    string `json:"senha,omitempty"`
	Criadoem time.Time `json:"Criadoem,omitempty"`
}