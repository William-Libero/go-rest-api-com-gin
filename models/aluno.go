package models

type Aluno struct {
	Nome string `json:"nome"`
	RG   string `json:"rg"`
	CPF  string `json:"cpf"`
}

var Alunos []Aluno
