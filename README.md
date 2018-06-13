# golang_backend
REST API para recetas de cocina

# Modelo SQL
En recetas_cocina/datos.sql y recetas_cocina/recetas_cocina.sql se encuentran algunos ejemplos de datos y su modelo respectivo.

Los otros archivos .sql corresponden a un modelo inicial no implementado mas extenso

# Endpoints

## Endpoint para listar recetas  (GET /recetas)
Tiene dos posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusOk, listado de recetas) -> Consulta exitosa

## Endpoint para ver una receta  (GET /recetas/{id})
Tiene cuatro posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> JSON que representa la receta

## Endpoint para crear una receta (POST /recetas)
Tiene dos posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusOk, receta) -> Impresion de la receta creada (no JSON)

## Endpoint para modificar recetas (POST /recetas/{id})
Tiene cuatro posibles salidas:
(StatusInternalServerError, listado_vacio) -> Error interno
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> Impresion de la receta modificada (no JSON)

## Endpoint para eliminar recetas (DELETE /recetas/{id})
Tiene tres posibles salidas
(StatusBadRequest, listado_vacio) -> Mala peticion
(StatusNotFound, listado_vacio) -> No se encontro la receta con esa id
(StatusOk, receta) -> Impresion de la receta eliminada (no JSON)



   
