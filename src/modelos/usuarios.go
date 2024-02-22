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

//vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar() error{
	if erro := usuario.validar(): erro != nil{
		return erro
	}
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == ""{
		return erros.New("O nome é obrigatorio")
	}

	if usuario.Nick == ""{
		return erros.New("O nick é obrigatorio")
	}

	if usuario.Email == ""{
		return erros.New("O email é obrigatorio")
	}

	if usuario.Senha == ""{
		return erros.New("A senha é obrigatoria")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = string.TrimSpace(usuario.Nome)
	usuario.Nick = string.TrimSpace(usuario.Nick)
	usuario.Email = string.TrimSpace(usuario.Email)
}