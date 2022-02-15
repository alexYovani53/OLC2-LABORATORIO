package definicion

import (
	"OLC2/analizador/ast/expresion"
	"OLC2/analizador/ast/interfaces"
	"OLC2/analizador/entorno"
	"encoding/json"
	"fmt"
	"github.com/colegno/arraylist"
)

var tipoDef = [5][5]entorno.TipoDato{
	{entorno.INTEGER, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.FLOAT, entorno.FLOAT, entorno.NULL, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.STRING, entorno.NULL, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.BOOLEAN, entorno.NULL},
	{entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL, entorno.NULL},
}

type Declaracion struct {
	ValorInicializacion interfaces.Expresion
	TipoVariables       entorno.TipoDato
	ListaVars           *arraylist.List
}

func NewDeclaracion(listaVars *arraylist.List, tipoVariables entorno.TipoDato) *Declaracion {
	return &Declaracion{
		TipoVariables: tipoVariables,
		ListaVars:     listaVars,
	}
}
func NewDeclaracionInicializacion(listaVars *arraylist.List, tipoVariables entorno.TipoDato, valInicial interfaces.Expresion) *Declaracion {
	return &Declaracion{
		TipoVariables:       tipoVariables,
		ListaVars:           listaVars,
		ValorInicializacion: valInicial,
	}
}

func (dec *Declaracion) Ejecutar(ent entorno.Entorno) interface{} {

	if dec.esInicializado() {
		if dec.ListaVars.Len() > 1 {
			return nil
		}

		retornoExpresion := dec.ValorInicializacion.ObtenerValor(ent)

		tipoExpresion := retornoExpresion.Tipo
		tipoDeclaracion := dec.TipoVariables

		tipoResultante := tipoDef[tipoDeclaracion][tipoExpresion]

		if tipoResultante == entorno.NULL {
			return nil
		}

		for i := 0; i < dec.ListaVars.Len(); i++ {

			varDeclarar := dec.ListaVars.GetValue(i).(expresion.Identificador)

			if ent.ExisteSimbolo(varDeclarar.Identificador) {
				fmt.Printf("Errror, variable %s ya declarada", varDeclarar.Identificador)
			} else {
				simboloTabala := entorno.NewSimboloIdentificadorValor(
					0,
					0,
					varDeclarar.Identificador,
					retornoExpresion.Valor,
					tipoResultante)

				ent.AgregarSimbolo(varDeclarar.Identificador, simboloTabala)
			}

		}

	}

	data, err := json.MarshalIndent(ent, "", "  ")
	if err != nil {
		panic(err)
	}

	stringEsQuery := string(data)
	fmt.Println(stringEsQuery)
	fmt.Printf("%v", dec.ListaVars)

	return nil
}

func (dec *Declaracion) esInicializado() bool {
	return dec.ValorInicializacion != nil
}
