INSERT INTO plato(nombre, descripcion) VALUES('Bandeja paisa', 'Plato de origen antioquenio');
INSERT INTO plato(nombre, descripcion) VALUES('Lulada', 'Plato de origen valluno');

INSERT INTO ingrediente(nombre, descripcion) VALUES('Banano', 'Fruta muy conocida.');
INSERT INTO ingrediente(nombre, descripcion) VALUES('Frijol', 'Frijol convencional.');
INSERT INTO ingrediente(nombre, descripcion) VALUES('Arroz', 'Arroz convencional.');
INSERT INTO ingrediente(nombre, descripcion) VALUES('Lulo', 'Lulo convencional.');

INSERT INTO unidad_medida(abreviacion, nombre, descripcion) VALUES('lb', 'Libra', 'Medio kilogramo.');

INSERT INTO receta(id_plato) VALUES(1);
INSERT INTO receta(id_plato) VALUES(2);

INSERT INTO ingredientes_en_receta VALUES(1,1,1,1,2);
INSERT INTO ingredientes_en_receta VALUES(1,1,2,1,2);
INSERT INTO ingredientes_en_receta VALUES(1,1,3,1,2);
INSERT INTO ingredientes_en_receta VALUES(2,2,4,1,2);

INSERT INTO pasos_receta(serial_receta, id_plato_receta, descripcion) VALUES(1,1,'Paso 1');
INSERT INTO pasos_receta(serial_receta, id_plato_receta, descripcion) VALUES(1,1,'Paso 2');
INSERT INTO pasos_receta(serial_receta, id_plato_receta, descripcion) VALUES(1,1,'Paso 3');
INSERT INTO pasos_receta(serial_receta, id_plato_receta, descripcion) VALUES(2,2,'Paso 4');
