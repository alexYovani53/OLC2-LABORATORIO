package expresion

import (
	"OLC2/analizador/entorno"
)

type Identificador struct {
	Identificador string
}

func NewIdentificador(identificador string) Identificador {
	return Identificador{Identificador: identificador}
}

func (ide Identificador) ObtenerValor(ent entorno.Entorno) entorno.ValorType {

	var encontrado bool = ent.ExisteSimbolo(ide.Identificador)

	if encontrado == false {
		return entorno.ValorType{Valor: nil, Tipo: entorno.NULL}
	}

	simbo := ent.ObtenerSimbolo(ide.Identificador)

	return entorno.ValorType{Valor: simbo.Valor, Tipo: simbo.Tipo}

}
