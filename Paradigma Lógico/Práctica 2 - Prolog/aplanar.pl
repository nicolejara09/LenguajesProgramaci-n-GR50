aplanar([],[]).

aplanar([[H|T]|R],S):- aplanar([H|T],T1),
    aplanar(R,R1), append(T1,R1,S),!.

aplanar([[]|T],S):- aplanar(T,S),!.

aplanar([H|T],[H|S]):- aplanar(T,S).

% aplanar([1,2,[3,[4,5],[6,7]]],X).
% X = [1, 2, 3, 4, 5, 6, 7].
