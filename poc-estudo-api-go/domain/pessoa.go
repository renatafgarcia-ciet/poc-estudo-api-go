package domain

type Pessoa struct {
    ID        string   `json:"id,omitempty"`
    Nome string   `json:"nome,omitempty"`
    Sobrenome  string   `json:"sobrenome,omitempty"`
    Endereco   *Endereco `json:"endereco,omitempty"`
}