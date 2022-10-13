substring(E,L) :-
    string_to_list(E, E1), string_to_list(L, L1),
    append(_,X,L1), append(E1,_,X),
    E \= [], !.

verificar(E, L) :-
    substring(E,L).

sub_cadenas(Substring, Cadena, Filtrado) :-
    include(verificar(Substring), Cadena, Filtrado).


%sub_cadenas("la", ["la casa", "el perro", "pintando la cerca"],Strings).
%Strings = ["la casa", "pintando la cerca"].
