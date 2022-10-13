sub_conjunto([],_).
sub_conjunto([H|T],E):- member(H,E), select(H,E,R), sub_conjunto(T,R).

%sub_conjunto([],[1,2,3,4,5]).
%True

%sub_conjunto([1,2,3],[1,2,3,4,5]).
%True

%sub_conjunto([1,2,6,8],[1,2,3,4,5]).
%False

