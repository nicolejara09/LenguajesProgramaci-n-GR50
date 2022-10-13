miembro(N,[N|_]).
miembro(N,[_|Tail]) :- miembro(N,Tail).

