"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Instruccion = void 0;
const NodoAst_1 = require("./NodoAst");
class Instruccion extends NodoAst_1.NodoAst {
    constructor(line, columna) {
        super(line, columna);
    }
}
exports.Instruccion = Instruccion;
//# sourceMappingURL=Instruccion.js.map