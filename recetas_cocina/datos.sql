INSERT INTO receta(nombre, descripcion, ingredientes, pasos) VALUES('Receta 1', 'Descripcion 1', 'Ingredientes 1', 'Pasos 1');
SELECT * FROM receta;
UPDATE receta SET nombre = 'Nombre 2', descripcion = 'Descripcion 2', ingredientes = 'Ingredientes 2', pasos = 'Paso 2' WHERE id_receta = 1;
DELETE FROM receta WHERE id_receta = 1;

