distanciaH(C1,C2,D):- string_chars(C1,L1),string_chars(C2,L2),distancia(L1,L2,D).

distancia([],[],0).
distancia([X|List1],[X|List2],D):- !,distancia(List1,List2,D).
distancia([_|List1],[_|List2],D):- distancia(List1,List2,D1),D is D1+1.
distancia([_|L1],L2,D):- distancia(L1,L2,D).
distancia(L1,[_|L2],D):- distancia(L1,L2,D).
distancia([],_,0).
distancia(_,[],0).
