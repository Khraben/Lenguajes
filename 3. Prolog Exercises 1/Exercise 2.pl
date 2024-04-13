es_lista([]).
es_lista([_|_]).

aplanar([],[]).
aplanar([H|T],Resultado):- es_lista(H),!,aplanar(H,SubListaH),aplanar(T,SubListaT),append(SubListaH,SubListaT,Resultado).
aplanar([H|T],[H|SubListaT]):- aplanar(T,SubListaT).
