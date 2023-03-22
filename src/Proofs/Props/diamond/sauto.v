From Hammer Require Import Tactics.
Section Example.

Variables P Q R S : Prop.
Lemma diamond : (P -> Q) -> (P -> R) -> (Q -> R -> S) -> P -> S.
(* TODO *)
Admitted.


End Example.