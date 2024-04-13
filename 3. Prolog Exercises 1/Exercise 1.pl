sub_conjunto([],_).
sub_conjunto([H|T],Lista):- member(H,Lista),sub_conjunto(T,Lista).
