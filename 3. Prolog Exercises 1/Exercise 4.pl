entrada(guacamole).
entrada(ensalada).
entrada(consomé).
entrada(tostadas_caprese).

carne(filete_de_cerdo).
carne(pollo_al_horno).
carne(carne_en_salsa).

pescado(tilapia).
pescado(salmón).
pescado(trucha).

postre(flan).
postre(nueces_con_miel).
postre(naranja_confitada).
postre(flan_de_coco).

calorias(guacamole, 200).
calorias(ensalada, 150).
calorias(consomé, 300).
calorias(tostadas_caprese, 250).
calorias(filete_de_cerdo, 400).
calorias(pollo_al_horno, 280).
calorias(carne_en_salsa, 320).
calorias(tilapia, 160).
calorias(salmón, 300).
calorias(trucha, 225).
calorias(flan, 200).
calorias(nueces_con_miel, 500).
calorias(naranja_confitada, 450).
calorias(flan_de_coco, 375).

plato_principal(comida):- carne(comida);pescado(comida).

comida(Entrada, PlatoPrincipal, Postre, Calorias) :-
    entrada(Entrada),
    plato_principal(PlatoPrincipal),
    postre(Postre),
    calorias(Entrada, CaloriasEntrada),
    calorias(PlatoPrincipal, CaloriasPlatoPrincipal),
    calorias(Postre, CaloriasPostre),
    Calorias is CaloriasEntrada + CaloriasPlatoPrincipal + CaloriasPostre.

comidas(MaxCalorias, Comidas) :-
    findall((E, PP, P, Cal),
            (comida(E, PP, P, Cal), Cal =< MaxCalorias),
            Lista),
    length(Comidas, 5),
    append(Comidas, _, Lista).

%ejemplo prueba:
%comidas(1000,Comidas)


