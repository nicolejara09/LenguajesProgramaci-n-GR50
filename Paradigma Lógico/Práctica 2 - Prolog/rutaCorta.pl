edge(i,a,3).
edge(i,b,4).
edge(a,i,6).
edge(a,c,9).
edge(a,d,2).
edge(b,i,6).
edge(b,c,7).
edge(b,d,3).
edge(c,a,5).
edge(c,x,1).
edge(c,b,9).
edge(d,a,6).
edge(d,f,9).
edge(d,b,5).
edge(x,c,3).
edge(f,d,2).

path(X, Y, N, Path) :- path(X, Y, N, [], Path).

path(X, Y, N, Seen, [X]) :- \+memberchk(X, Seen),edge(X, Y, N).
path(X, Z, N, Seen, [X|T]) :- \+memberchk(X, Seen),edge(X, Y, N0),
    path(Y, Z, N1, [X|Seen], T), \+memberchk(X, T),N is N0 + N1.

memberchk(X, L) :- once(member1(X, L)).
member1(X, [X|_]).
member1(X, [_|Xs]) :- member(X, Xs).

uniq_shortest_path(X, Y, MinCost, Path) :- path(X, Y, MinCost, Path),
    \+ (path(X, Y, LowerCost, OtherPath),OtherPath \= Path,
    LowerCost =< MinCost).

% uniq_shortest_path(x, f, MinCost, Path).
% MinCost = 19,
% Path = [x, c, a, d]
