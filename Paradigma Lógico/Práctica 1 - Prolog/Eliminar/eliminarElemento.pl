remove(_,[],[]).

remove(X,[X|W],Y) :-
    remove(X,W,Y),!.

remove(X,[F|W],[F|W1]) :-
    remove(X,W,W1).

%Y es la lista creada sin el elemento que se elimin�.

%X es el elemento que se buscar� para eliminarlo de la lista.

%W es la cola, esta recorrer� la lista hasta hacer matching con el
%elemento que se quiere eliminar.

%Resultado a mostrar
%?- remove(2,[1,2,3,4,5],L).
%L = [1, 3, 4, 5].
