From Hammer Require Import Tactics.
Section Example.

Variables P Q R : Prop.
Lemma imp_perm : (P -> Q -> R) -> Q -> P -> R.
(* TODO *)
Admitted.


End Example.