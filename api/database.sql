CREATE DATABASE recipes;

CREATE TABLE IF NOT EXISTS recipes.recipes (
    id SERIAL PRIMARY KEY,
    name STRING(255),
    calories DECIMAL,
    img STRING(255),
    ingredients STRING,
    instructions STRING
);

INSERT INTO recipes.recipes (name, calories, img, ingredients, instructions) VALUES ('Torta 3 Leches', 225, 'https://img-global.cpcdn.com/002_recipes/fdf55158b4cad507/640x640sq70/photo.jpg',
'---- Para las 3 Leches
1 lata leche evaporada
1 lata leche condensada
200 ml Leche

---- Para la Cubierta
100 gr Queso Crema
400 gr Crema de Leche o Crema para batir
1/4 Taza de Azúcar glass o impalpable (O al gusto)
100 gr Dulce de Leche
Fresas para Decorar',
'Mezclamos la leche evaporada, la leche condensada y la leche entera. Y le echamos encima al bizcocho poco a poco para que la vaya absorbiendo, (el bizcocho debe estar en el molde donde se preparó previamente cubierto con papel transparente)

Si se desean bien húmeda hay que echarle buena cantidad de la mezcla. Desps se tapa con papel transparente y se mete a la nevera mínimo 3 hrs para que lo absorba bien (preferiblemente dejarla toda la noche)

En un bol, mezclar la crema de leche, con el queso crema y el azúcar glass, hasta que quede a punto nieve, el azúcar se va agregando poco a poco (ir probando la mezcla para dejar el dulzor deseado, si lo desea más dulce, agregar más azúcar.)

Dividir la mezcla en dos, en una agregar el dulce de leche (También puede agregar chocolate, mermelada, o alguna fruta. lo que más les guste) y mezclar y la otra dejarla así,

Dividir el bizcocho en 2 y agregarle la mezcla con dulce de leche, dejando más o menos 1cm antes del borde

Poner la otra capa del bizcocho y tapar con la crema blanca, y decorar. (La crema debe quedar bien firme o se deslizara del pastel, si se siente muy suave meter al congelador 15 min antes de untarla.');


INSERT INTO recipes.recipes (name, calories, img, ingredients, instructions) VALUES ('Postre Napoleón', 151, 'http://recetasdeisabel.com/wp-content/uploads/2016/09/POSTRE-NAPOLEO%CC%81N-TODO-UN-CAMPEO%CC%81N-PORTADA-860-X-573.jpg',
'900 mililitros leche
90 gr maizena
350 gr leche condensada
100 gr crema de leche
2 chocolatinas Jet
8 recipientes plásticos para postre
1 taco galletas ducales',
'Se licúan los 90 g de maizena con los 900 ml de leche.

La mezcla anterior se vierte en un recipiente métalico como una olla, y se va revolviendo a fuego medio hasta que espese.

Nota: es importante no dejar de revolver, el objetivo es que no se pegue la mezcla a la olla, ni queden grumitos.

Cuando la mezcla ya esté espesa, agregue la leche condensada, revuelva un minuto, y luego agregue la crema de leche.

Apague el fogón, espere un minuto, y empiece a verter la mezcla en los recipientes de forma que quede una capa de mezcla y una de galletas ducales y así sucesivamente como lo desee preparar o al gusto de cada quien.

Espere 15 minutos a que el postre adquiera más textura, Luego de esto rallé un cuadrito de chocolatina a cada postre o más si lo considera necesario.

Opcional: puede agregarle otro tipo de aditivos como maní, M&Ms, almendras, coco rallado, etc.');