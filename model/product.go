package model

// CRIANDO UMA STRUCT DE ACORDO COM A ESTRUTURA DE PRODUTOS NO BANDO DE DADOS

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
