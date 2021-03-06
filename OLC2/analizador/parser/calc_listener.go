// Code generated from Calc.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by Calc.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterListaClases is called when entering the listaClases production.
	EnterListaClases(c *ListaClasesContext)

	// EnterClases is called when entering the clases production.
	EnterClases(c *ClasesContext)

	// EnterCuerpoClase is called when entering the cuerpoClase production.
	EnterCuerpoClase(c *CuerpoClaseContext)

	// EnterContenidoClase is called when entering the contenidoClase production.
	EnterContenidoClase(c *ContenidoClaseContext)

	// EnterItemClase is called when entering the itemClase production.
	EnterItemClase(c *ItemClaseContext)

	// EnterFuncionItem is called when entering the funcionItem production.
	EnterFuncionItem(c *FuncionItemContext)

	// EnterModaccess is called when entering the modaccess production.
	EnterModaccess(c *ModaccessContext)

	// EnterParametros is called when entering the parametros production.
	EnterParametros(c *ParametrosContext)

	// EnterParametro is called when entering the parametro production.
	EnterParametro(c *ParametroContext)

	// EnterFuncmain is called when entering the funcmain production.
	EnterFuncmain(c *FuncmainContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterDec_objeto is called when entering the dec_objeto production.
	EnterDec_objeto(c *Dec_objetoContext)

	// EnterDec_arr is called when entering the dec_arr production.
	EnterDec_arr(c *Dec_arrContext)

	// EnterDimensiones is called when entering the dimensiones production.
	EnterDimensiones(c *DimensionesContext)

	// EnterDimension is called when entering the dimension production.
	EnterDimension(c *DimensionContext)

	// EnterListanchos is called when entering the listanchos production.
	EnterListanchos(c *ListanchosContext)

	// EnterAncho is called when entering the ancho production.
	EnterAncho(c *AnchoContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterIf_instr is called when entering the if_instr production.
	EnterIf_instr(c *If_instrContext)

	// EnterListaelseif is called when entering the listaelseif production.
	EnterListaelseif(c *ListaelseifContext)

	// EnterElse_if is called when entering the else_if production.
	EnterElse_if(c *Else_ifContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterConsola is called when entering the consola production.
	EnterConsola(c *ConsolaContext)

	// EnterLlamada is called when entering the llamada production.
	EnterLlamada(c *LlamadaContext)

	// EnterListaExpresiones is called when entering the listaExpresiones production.
	EnterListaExpresiones(c *ListaExpresionesContext)

	// EnterDeclaracionIni is called when entering the declaracionIni production.
	EnterDeclaracionIni(c *DeclaracionIniContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterRetorno is called when entering the retorno production.
	EnterRetorno(c *RetornoContext)

	// EnterListides is called when entering the listides production.
	EnterListides(c *ListidesContext)

	// EnterTiposvars is called when entering the tiposvars production.
	EnterTiposvars(c *TiposvarsContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// EnterArraydata is called when entering the arraydata production.
	EnterArraydata(c *ArraydataContext)

	// EnterInstancia is called when entering the instancia production.
	EnterInstancia(c *InstanciaContext)

	// EnterInstanciaClase is called when entering the instanciaClase production.
	EnterInstanciaClase(c *InstanciaClaseContext)

	// EnterAccesoarr is called when entering the accesoarr production.
	EnterAccesoarr(c *AccesoarrContext)

	// EnterAccesoObjeto is called when entering the accesoObjeto production.
	EnterAccesoObjeto(c *AccesoObjetoContext)

	// EnterListaAccesos is called when entering the listaAccesos production.
	EnterListaAccesos(c *ListaAccesosContext)

	// EnterAcceso is called when entering the acceso production.
	EnterAcceso(c *AccesoContext)

	// EnterExpr_rel is called when entering the expr_rel production.
	EnterExpr_rel(c *Expr_relContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterExpr_valor is called when entering the expr_valor production.
	EnterExpr_valor(c *Expr_valorContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitListaClases is called when exiting the listaClases production.
	ExitListaClases(c *ListaClasesContext)

	// ExitClases is called when exiting the clases production.
	ExitClases(c *ClasesContext)

	// ExitCuerpoClase is called when exiting the cuerpoClase production.
	ExitCuerpoClase(c *CuerpoClaseContext)

	// ExitContenidoClase is called when exiting the contenidoClase production.
	ExitContenidoClase(c *ContenidoClaseContext)

	// ExitItemClase is called when exiting the itemClase production.
	ExitItemClase(c *ItemClaseContext)

	// ExitFuncionItem is called when exiting the funcionItem production.
	ExitFuncionItem(c *FuncionItemContext)

	// ExitModaccess is called when exiting the modaccess production.
	ExitModaccess(c *ModaccessContext)

	// ExitParametros is called when exiting the parametros production.
	ExitParametros(c *ParametrosContext)

	// ExitParametro is called when exiting the parametro production.
	ExitParametro(c *ParametroContext)

	// ExitFuncmain is called when exiting the funcmain production.
	ExitFuncmain(c *FuncmainContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitDec_objeto is called when exiting the dec_objeto production.
	ExitDec_objeto(c *Dec_objetoContext)

	// ExitDec_arr is called when exiting the dec_arr production.
	ExitDec_arr(c *Dec_arrContext)

	// ExitDimensiones is called when exiting the dimensiones production.
	ExitDimensiones(c *DimensionesContext)

	// ExitDimension is called when exiting the dimension production.
	ExitDimension(c *DimensionContext)

	// ExitListanchos is called when exiting the listanchos production.
	ExitListanchos(c *ListanchosContext)

	// ExitAncho is called when exiting the ancho production.
	ExitAncho(c *AnchoContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitIf_instr is called when exiting the if_instr production.
	ExitIf_instr(c *If_instrContext)

	// ExitListaelseif is called when exiting the listaelseif production.
	ExitListaelseif(c *ListaelseifContext)

	// ExitElse_if is called when exiting the else_if production.
	ExitElse_if(c *Else_ifContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitConsola is called when exiting the consola production.
	ExitConsola(c *ConsolaContext)

	// ExitLlamada is called when exiting the llamada production.
	ExitLlamada(c *LlamadaContext)

	// ExitListaExpresiones is called when exiting the listaExpresiones production.
	ExitListaExpresiones(c *ListaExpresionesContext)

	// ExitDeclaracionIni is called when exiting the declaracionIni production.
	ExitDeclaracionIni(c *DeclaracionIniContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitRetorno is called when exiting the retorno production.
	ExitRetorno(c *RetornoContext)

	// ExitListides is called when exiting the listides production.
	ExitListides(c *ListidesContext)

	// ExitTiposvars is called when exiting the tiposvars production.
	ExitTiposvars(c *TiposvarsContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)

	// ExitArraydata is called when exiting the arraydata production.
	ExitArraydata(c *ArraydataContext)

	// ExitInstancia is called when exiting the instancia production.
	ExitInstancia(c *InstanciaContext)

	// ExitInstanciaClase is called when exiting the instanciaClase production.
	ExitInstanciaClase(c *InstanciaClaseContext)

	// ExitAccesoarr is called when exiting the accesoarr production.
	ExitAccesoarr(c *AccesoarrContext)

	// ExitAccesoObjeto is called when exiting the accesoObjeto production.
	ExitAccesoObjeto(c *AccesoObjetoContext)

	// ExitListaAccesos is called when exiting the listaAccesos production.
	ExitListaAccesos(c *ListaAccesosContext)

	// ExitAcceso is called when exiting the acceso production.
	ExitAcceso(c *AccesoContext)

	// ExitExpr_rel is called when exiting the expr_rel production.
	ExitExpr_rel(c *Expr_relContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitExpr_valor is called when exiting the expr_valor production.
	ExitExpr_valor(c *Expr_valorContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
