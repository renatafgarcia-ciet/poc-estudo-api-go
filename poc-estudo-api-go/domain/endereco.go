package domain

type Endereco struct {
    Cidade  string `json:"cidade,omitempty"`
    Estado string `json:"estado,omitempty"`
}